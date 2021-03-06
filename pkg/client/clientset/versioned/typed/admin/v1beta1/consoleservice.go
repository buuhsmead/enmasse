/*
 * Copyright 2018-2019, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/enmasseproject/enmasse/pkg/apis/admin/v1beta1"
	scheme "github.com/enmasseproject/enmasse/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ConsoleServicesGetter has a method to return a ConsoleServiceInterface.
// A group's client should implement this interface.
type ConsoleServicesGetter interface {
	ConsoleServices(namespace string) ConsoleServiceInterface
}

// ConsoleServiceInterface has methods to work with ConsoleService resources.
type ConsoleServiceInterface interface {
	Create(*v1beta1.ConsoleService) (*v1beta1.ConsoleService, error)
	Update(*v1beta1.ConsoleService) (*v1beta1.ConsoleService, error)
	UpdateStatus(*v1beta1.ConsoleService) (*v1beta1.ConsoleService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.ConsoleService, error)
	List(opts v1.ListOptions) (*v1beta1.ConsoleServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ConsoleService, err error)
	ConsoleServiceExpansion
}

// consoleServices implements ConsoleServiceInterface
type consoleServices struct {
	client rest.Interface
	ns     string
}

// newConsoleServices returns a ConsoleServices
func newConsoleServices(c *AdminV1beta1Client, namespace string) *consoleServices {
	return &consoleServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the consoleService, and returns the corresponding consoleService object, and an error if there is any.
func (c *consoleServices) Get(name string, options v1.GetOptions) (result *v1beta1.ConsoleService, err error) {
	result = &v1beta1.ConsoleService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("consoleservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ConsoleServices that match those selectors.
func (c *consoleServices) List(opts v1.ListOptions) (result *v1beta1.ConsoleServiceList, err error) {
	result = &v1beta1.ConsoleServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("consoleservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested consoleServices.
func (c *consoleServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("consoleservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a consoleService and creates it.  Returns the server's representation of the consoleService, and an error, if there is any.
func (c *consoleServices) Create(consoleService *v1beta1.ConsoleService) (result *v1beta1.ConsoleService, err error) {
	result = &v1beta1.ConsoleService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("consoleservices").
		Body(consoleService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a consoleService and updates it. Returns the server's representation of the consoleService, and an error, if there is any.
func (c *consoleServices) Update(consoleService *v1beta1.ConsoleService) (result *v1beta1.ConsoleService, err error) {
	result = &v1beta1.ConsoleService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("consoleservices").
		Name(consoleService.Name).
		Body(consoleService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *consoleServices) UpdateStatus(consoleService *v1beta1.ConsoleService) (result *v1beta1.ConsoleService, err error) {
	result = &v1beta1.ConsoleService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("consoleservices").
		Name(consoleService.Name).
		SubResource("status").
		Body(consoleService).
		Do().
		Into(result)
	return
}

// Delete takes name of the consoleService and deletes it. Returns an error if one occurs.
func (c *consoleServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("consoleservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *consoleServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("consoleservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched consoleService.
func (c *consoleServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ConsoleService, err error) {
	result = &v1beta1.ConsoleService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("consoleservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
