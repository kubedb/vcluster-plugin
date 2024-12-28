package syncers

import (
	"context"
	"encoding/json"
	"fmt"
	synccontext "github.com/loft-sh/vcluster/pkg/controllers/syncer/context"
	"github.com/loft-sh/vcluster/pkg/controllers/syncer/translator"
	"github.com/loft-sh/vcluster/pkg/scheme"
	synctypes "github.com/loft-sh/vcluster/pkg/types"
	"github.com/loft-sh/vcluster/pkg/util/translate"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	"k8s.io/apimachinery/pkg/api/equality"
	api "kubedb.dev/apimachinery/apis/kubedb/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	// Make sure our scheme is registered
	_ = api.AddToScheme(scheme.Scheme)
}

func NewPostgresSyncer(ctx *synccontext.RegisterContext) synctypes.Base {
	return &postgresSyncer{
		NamespacedTranslator: translator.NewNamespacedTranslator(ctx, api.ResourceSingularPostgres, &api.Postgres{}),
	}
}

type postgresSyncer struct {
	translator.NamespacedTranslator
}

var _ synctypes.Initializer = &postgresSyncer{}

func (s *postgresSyncer) Init(ctx *synccontext.RegisterContext) error {
	/*
		out, err := os.ReadFile("manifests/crds.yaml")
		if err != nil {
			return err
		}

		gvk := api.SchemeGroupVersion.WithKind(api.ResourceKindPostgres)
		err = util.EnsureCRD(ctx.Context, ctx.PhysicalManager.GetConfig(), out, gvk)
		if err != nil {
			return err
		}
	*/

	_, _, err := translate.EnsureCRDFromPhysicalCluster(ctx.Context, ctx.PhysicalManager.GetConfig(), ctx.VirtualManager.GetConfig(), api.SchemeGroupVersion.WithKind(api.ResourceKindPostgres))
	return err
}

var _ synctypes.Syncer = &postgresSyncer{}

func (s *postgresSyncer) SyncToHost(ctx *synccontext.SyncContext, vObj client.Object) (ctrl.Result, error) {
	return s.SyncToHostCreate(ctx, vObj, s.translate(ctx.Context, vObj.(*api.Postgres)))
}

func (s *postgresSyncer) Sync(ctx *synccontext.SyncContext, pObj client.Object, vObj client.Object) (ctrl.Result, error) {
	return s.SyncToHostUpdate(ctx, vObj, s.translateUpdate(ctx.Context, pObj.(*api.Postgres), vObj.(*api.Postgres)))
}

func (s *postgresSyncer) translate(ctx context.Context, vObj *api.Postgres) *api.Postgres {
	newObj := s.TranslateMetadata(ctx, vObj).(*api.Postgres)

	// newObj.Spec = s.translateSpec(vObj).Spec
	return newObj
}

func (s *postgresSyncer) translateSpec(vObj *api.Postgres) *api.Postgres {
	var pObj api.Postgres

	/*
		pObj.Spec.Refs = make([]api.TypedObjectReference, len(vObj.Spec.Refs))
		for i, ref := range vObj.Spec.Refs {
			switch ref.Kind {
			case "ClusterIssuer":
				ref.Name = translate.Default.PhysicalNameClusterScoped(ref.Name)
			case "Issuer", "Secret":
				vNamespace := ref.Namespace
				if vNamespace == "" {
					vNamespace = vObj.GetNamespace()
				}
				ref.Name = translate.Default.PhysicalName(ref.Name, vNamespace)
				ref.Namespace = translate.Default.PhysicalNamespace(vNamespace)
			}
			pObj.Spec.Refs[i] = ref
		}
	*/

	return &pObj
}

func (s *postgresSyncer) translateUpdate(ctx context.Context, pObj, vObj *api.Postgres) *api.Postgres {
	var updated *api.Postgres

	// check annotations & labels
	changed, updatedAnnotations, updatedLabels := s.TranslateMetadataUpdate(ctx, vObj, pObj)
	if changed {
		updated = translator.NewIfNil(updated, pObj)
		updated.Labels = updatedLabels
		updated.Annotations = updatedAnnotations
	}

	// check spec
	updatedSpecObj := s.translateSpec(vObj)
	if !equality.Semantic.DeepEqual(updatedSpecObj.Spec, pObj.Spec) {
		updated = translator.NewIfNil(updated, pObj)
		updated.Spec = updatedSpecObj.Spec
	}

	return updated
}

func printDiff(original, updated client.Object) error {
	if updated == nil {
		return nil
	}
	originalBytes, err := json.Marshal(original)
	if err != nil {
		return err
	}

	updatedBytes, err := json.Marshal(updated)
	if err != nil {
		return err
	}

	differ := diff.New()
	d, err := differ.Compare(originalBytes, updatedBytes)
	if err != nil {
		return err
	}

	if d.Modified() {
		config := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       true,
		}

		f := formatter.NewAsciiFormatter(original, config)
		result, err := f.Format(d)
		if err != nil {
			return err
		}
		fmt.Println(result)
		return nil
	}

	return nil
}
