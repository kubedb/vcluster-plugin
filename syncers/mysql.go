package syncers

import (
	"context"
	"fmt"
	synccontext "github.com/loft-sh/vcluster/pkg/controllers/syncer/context"
	"github.com/loft-sh/vcluster/pkg/controllers/syncer/translator"
	"github.com/loft-sh/vcluster/pkg/scheme"
	synctypes "github.com/loft-sh/vcluster/pkg/types"
	"github.com/loft-sh/vcluster/pkg/util/translate"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	api "kubedb.dev/apimachinery/apis/kubedb/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	// Make sure our scheme is registered
	_ = api.AddToScheme(scheme.Scheme)
}

func NewMySQLSyncer(ctx *synccontext.RegisterContext) synctypes.Base {
	return &mysqlSyncer{
		NamespacedTranslator: translator.NewNamespacedTranslator(ctx, api.ResourceSingularMySQL, &api.MySQL{}),
	}
}

type mysqlSyncer struct {
	translator.NamespacedTranslator
	rtx *synccontext.RegisterContext
}

var _ synctypes.Initializer = &mysqlSyncer{}

func (s *mysqlSyncer) Init(ctx *synccontext.RegisterContext) error {
	s.rtx = ctx
	_, _, err := translate.EnsureCRDFromPhysicalCluster(ctx.Context, ctx.PhysicalManager.GetConfig(), ctx.VirtualManager.GetConfig(), api.SchemeGroupVersion.WithKind(api.ResourceKindMySQL))
	return err
}

var _ synctypes.Syncer = &mysqlSyncer{}

func (s *mysqlSyncer) SyncToHost(ctx *synccontext.SyncContext, vObj client.Object) (ctrl.Result, error) {
	// Call removeConflictingOwnerReference before syncing to the host
	err := s.removeConflictingOwnerReference(ctx.Context, vObj.GetNamespace(), vObj.GetName())
	if err != nil {
		return ctrl.Result{}, err
	}
	return s.SyncToHostCreate(ctx, vObj, s.translate(ctx.Context, vObj.(*api.MySQL)))
}

func (s *mysqlSyncer) Sync(ctx *synccontext.SyncContext, pObj client.Object, vObj client.Object) (ctrl.Result, error) {
	return s.SyncToHostUpdate(ctx, vObj, s.translateUpdate(ctx.Context, pObj.(*api.MySQL), vObj.(*api.MySQL)))
}

func (s *mysqlSyncer) translate(ctx context.Context, vObj *api.MySQL) *api.MySQL {
	newObj := s.TranslateMetadata(ctx, vObj).(*api.MySQL)
	return newObj
}

func (s *mysqlSyncer) translateSpec(vObj *api.MySQL) *api.MySQL {
	var pObj api.MySQL
	return &pObj
}

func (s *mysqlSyncer) translateUpdate(ctx context.Context, pObj, vObj *api.MySQL) *api.MySQL {
	var updated *api.MySQL

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

//func printDiff(original, updated client.Object) error {
//	if updated == nil {
//		return nil
//	}
//	originalBytes, err := json.Marshal(original)
//	if err != nil {
//		return err
//	}
//
//	updatedBytes, err := json.Marshal(updated)
//	if err != nil {
//		return err
//	}
//
//	differ := diff.New()
//	d, err := differ.Compare(originalBytes, updatedBytes)
//	if err != nil {
//		return err
//	}
//
//	if d.Modified() {
//		config := formatter.AsciiFormatterConfig{
//			ShowArrayIndex: true,
//			Coloring:       true,
//		}
//
//		f := formatter.NewAsciiFormatter(original, config)
//		result, err := f.Format(d)
//		if err != nil {
//			return err
//		}
//		fmt.Println(result)
//		return nil
//	}
//
//	return nil
//}

func (s *mysqlSyncer) removeConflictingOwnerReference(ctx context.Context, namespace, serviceName string) error {
	clientset := kubernetes.NewForConfigOrDie(s.rtx.VirtualManager.GetConfig())
	service, err := clientset.CoreV1().Services(namespace).Get(ctx, serviceName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get service: %v", err)
	}

	var newOwnerReferences []metav1.OwnerReference
	for _, ownerRef := range service.OwnerReferences {
		if ownerRef.Controller != nil && *ownerRef.Controller {
			ownerRef.Controller = nil
		}
		ownerRef.UID = ""
		newOwnerReferences = append(newOwnerReferences, ownerRef)
	}
	service.OwnerReferences = newOwnerReferences

	_, err = clientset.CoreV1().Services(namespace).Update(ctx, service, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to update service: %v", err)
	}

	return nil
}
