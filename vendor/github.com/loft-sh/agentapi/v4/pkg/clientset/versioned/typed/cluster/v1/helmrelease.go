// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/agentapi/v4/pkg/apis/loft/cluster/v1"
	scheme "github.com/loft-sh/agentapi/v4/pkg/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HelmReleasesGetter has a method to return a HelmReleaseInterface.
// A group's client should implement this interface.
type HelmReleasesGetter interface {
	HelmReleases(namespace string) HelmReleaseInterface
}

// HelmReleaseInterface has methods to work with HelmRelease resources.
type HelmReleaseInterface interface {
	Create(ctx context.Context, helmRelease *v1.HelmRelease, opts metav1.CreateOptions) (*v1.HelmRelease, error)
	Update(ctx context.Context, helmRelease *v1.HelmRelease, opts metav1.UpdateOptions) (*v1.HelmRelease, error)
	UpdateStatus(ctx context.Context, helmRelease *v1.HelmRelease, opts metav1.UpdateOptions) (*v1.HelmRelease, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.HelmRelease, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.HelmReleaseList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.HelmRelease, err error)
	HelmReleaseExpansion
}

// helmReleases implements HelmReleaseInterface
type helmReleases struct {
	client rest.Interface
	ns     string
}

// newHelmReleases returns a HelmReleases
func newHelmReleases(c *ClusterV1Client, namespace string) *helmReleases {
	return &helmReleases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the helmRelease, and returns the corresponding helmRelease object, and an error if there is any.
func (c *helmReleases) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.HelmRelease, err error) {
	result = &v1.HelmRelease{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("helmreleases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HelmReleases that match those selectors.
func (c *helmReleases) List(ctx context.Context, opts metav1.ListOptions) (result *v1.HelmReleaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.HelmReleaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("helmreleases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested helmReleases.
func (c *helmReleases) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("helmreleases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a helmRelease and creates it.  Returns the server's representation of the helmRelease, and an error, if there is any.
func (c *helmReleases) Create(ctx context.Context, helmRelease *v1.HelmRelease, opts metav1.CreateOptions) (result *v1.HelmRelease, err error) {
	result = &v1.HelmRelease{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("helmreleases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(helmRelease).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a helmRelease and updates it. Returns the server's representation of the helmRelease, and an error, if there is any.
func (c *helmReleases) Update(ctx context.Context, helmRelease *v1.HelmRelease, opts metav1.UpdateOptions) (result *v1.HelmRelease, err error) {
	result = &v1.HelmRelease{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("helmreleases").
		Name(helmRelease.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(helmRelease).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *helmReleases) UpdateStatus(ctx context.Context, helmRelease *v1.HelmRelease, opts metav1.UpdateOptions) (result *v1.HelmRelease, err error) {
	result = &v1.HelmRelease{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("helmreleases").
		Name(helmRelease.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(helmRelease).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the helmRelease and deletes it. Returns an error if one occurs.
func (c *helmReleases) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("helmreleases").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *helmReleases) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("helmreleases").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched helmRelease.
func (c *helmReleases) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.HelmRelease, err error) {
	result = &v1.HelmRelease{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("helmreleases").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
