/**
 * Copyright 2020 Rafael Fernández López <ereslibre@ereslibre.es>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 **/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/oneinfra/oneinfra/apis/cluster/v1alpha1"
	scheme "github.com/oneinfra/oneinfra/pkg/clientset/manager/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ComponentsGetter has a method to return a ComponentInterface.
// A group's client should implement this interface.
type ComponentsGetter interface {
	Components(namespace string) ComponentInterface
}

// ComponentInterface has methods to work with Component resources.
type ComponentInterface interface {
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Component, error)
	List(opts v1.ListOptions) (*v1alpha1.ComponentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	ComponentExpansion
}

// components implements ComponentInterface
type components struct {
	client rest.Interface
	ns     string
}

// newComponents returns a Components
func newComponents(c *ClusterV1alpha1Client, namespace string) *components {
	return &components{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the component, and returns the corresponding component object, and an error if there is any.
func (c *components) Get(name string, options v1.GetOptions) (result *v1alpha1.Component, err error) {
	result = &v1alpha1.Component{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("components").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Components that match those selectors.
func (c *components) List(opts v1.ListOptions) (result *v1alpha1.ComponentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComponentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("components").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested components.
func (c *components) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("components").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Delete takes name of the component and deletes it. Returns an error if one occurs.
func (c *components) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("components").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *components) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("components").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}
