// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/jenkins-x/jx-preview/pkg/apis/preview/v1alpha1"
	scheme "github.com/jenkins-x/jx-preview/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PreviewsGetter has a method to return a PreviewInterface.
// A group's client should implement this interface.
type PreviewsGetter interface {
	Previews(namespace string) PreviewInterface
}

// PreviewInterface has methods to work with Preview resources.
type PreviewInterface interface {
	Create(*v1alpha1.Preview) (*v1alpha1.Preview, error)
	Update(*v1alpha1.Preview) (*v1alpha1.Preview, error)
	UpdateStatus(*v1alpha1.Preview) (*v1alpha1.Preview, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Preview, error)
	List(opts v1.ListOptions) (*v1alpha1.PreviewList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Preview, err error)
	PreviewExpansion
}

// previews implements PreviewInterface
type previews struct {
	client rest.Interface
	ns     string
}

// newPreviews returns a Previews
func newPreviews(c *PreviewV1alpha1Client, namespace string) *previews {
	return &previews{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the preview, and returns the corresponding preview object, and an error if there is any.
func (c *previews) Get(name string, options v1.GetOptions) (result *v1alpha1.Preview, err error) {
	result = &v1alpha1.Preview{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("previews").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Previews that match those selectors.
func (c *previews) List(opts v1.ListOptions) (result *v1alpha1.PreviewList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PreviewList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("previews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested previews.
func (c *previews) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("previews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a preview and creates it.  Returns the server's representation of the preview, and an error, if there is any.
func (c *previews) Create(preview *v1alpha1.Preview) (result *v1alpha1.Preview, err error) {
	result = &v1alpha1.Preview{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("previews").
		Body(preview).
		Do().
		Into(result)
	return
}

// Update takes the representation of a preview and updates it. Returns the server's representation of the preview, and an error, if there is any.
func (c *previews) Update(preview *v1alpha1.Preview) (result *v1alpha1.Preview, err error) {
	result = &v1alpha1.Preview{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("previews").
		Name(preview.Name).
		Body(preview).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *previews) UpdateStatus(preview *v1alpha1.Preview) (result *v1alpha1.Preview, err error) {
	result = &v1alpha1.Preview{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("previews").
		Name(preview.Name).
		SubResource("status").
		Body(preview).
		Do().
		Into(result)
	return
}

// Delete takes name of the preview and deletes it. Returns an error if one occurs.
func (c *previews) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("previews").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *previews) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("previews").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched preview.
func (c *previews) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Preview, err error) {
	result = &v1alpha1.Preview{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("previews").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}