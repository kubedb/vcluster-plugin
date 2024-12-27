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

// ChartInfosGetter has a method to return a ChartInfoInterface.
// A group's client should implement this interface.
type ChartInfosGetter interface {
	ChartInfos() ChartInfoInterface
}

// ChartInfoInterface has methods to work with ChartInfo resources.
type ChartInfoInterface interface {
	Create(ctx context.Context, chartInfo *v1.ChartInfo, opts metav1.CreateOptions) (*v1.ChartInfo, error)
	Update(ctx context.Context, chartInfo *v1.ChartInfo, opts metav1.UpdateOptions) (*v1.ChartInfo, error)
	UpdateStatus(ctx context.Context, chartInfo *v1.ChartInfo, opts metav1.UpdateOptions) (*v1.ChartInfo, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ChartInfo, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ChartInfoList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ChartInfo, err error)
	ChartInfoExpansion
}

// chartInfos implements ChartInfoInterface
type chartInfos struct {
	client rest.Interface
}

// newChartInfos returns a ChartInfos
func newChartInfos(c *ClusterV1Client) *chartInfos {
	return &chartInfos{
		client: c.RESTClient(),
	}
}

// Get takes name of the chartInfo, and returns the corresponding chartInfo object, and an error if there is any.
func (c *chartInfos) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ChartInfo, err error) {
	result = &v1.ChartInfo{}
	err = c.client.Get().
		Resource("chartinfos").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ChartInfos that match those selectors.
func (c *chartInfos) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ChartInfoList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ChartInfoList{}
	err = c.client.Get().
		Resource("chartinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested chartInfos.
func (c *chartInfos) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("chartinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a chartInfo and creates it.  Returns the server's representation of the chartInfo, and an error, if there is any.
func (c *chartInfos) Create(ctx context.Context, chartInfo *v1.ChartInfo, opts metav1.CreateOptions) (result *v1.ChartInfo, err error) {
	result = &v1.ChartInfo{}
	err = c.client.Post().
		Resource("chartinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(chartInfo).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a chartInfo and updates it. Returns the server's representation of the chartInfo, and an error, if there is any.
func (c *chartInfos) Update(ctx context.Context, chartInfo *v1.ChartInfo, opts metav1.UpdateOptions) (result *v1.ChartInfo, err error) {
	result = &v1.ChartInfo{}
	err = c.client.Put().
		Resource("chartinfos").
		Name(chartInfo.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(chartInfo).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *chartInfos) UpdateStatus(ctx context.Context, chartInfo *v1.ChartInfo, opts metav1.UpdateOptions) (result *v1.ChartInfo, err error) {
	result = &v1.ChartInfo{}
	err = c.client.Put().
		Resource("chartinfos").
		Name(chartInfo.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(chartInfo).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the chartInfo and deletes it. Returns an error if one occurs.
func (c *chartInfos) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("chartinfos").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *chartInfos) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("chartinfos").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched chartInfo.
func (c *chartInfos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ChartInfo, err error) {
	result = &v1.ChartInfo{}
	err = c.client.Patch(pt).
		Resource("chartinfos").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
