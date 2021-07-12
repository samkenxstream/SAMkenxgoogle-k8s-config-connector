// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/identityplatform/v1beta1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IdentityPlatformTenantOAuthIDPConfigsGetter has a method to return a IdentityPlatformTenantOAuthIDPConfigInterface.
// A group's client should implement this interface.
type IdentityPlatformTenantOAuthIDPConfigsGetter interface {
	IdentityPlatformTenantOAuthIDPConfigs(namespace string) IdentityPlatformTenantOAuthIDPConfigInterface
}

// IdentityPlatformTenantOAuthIDPConfigInterface has methods to work with IdentityPlatformTenantOAuthIDPConfig resources.
type IdentityPlatformTenantOAuthIDPConfigInterface interface {
	Create(ctx context.Context, identityPlatformTenantOAuthIDPConfig *v1beta1.IdentityPlatformTenantOAuthIDPConfig, opts v1.CreateOptions) (*v1beta1.IdentityPlatformTenantOAuthIDPConfig, error)
	Update(ctx context.Context, identityPlatformTenantOAuthIDPConfig *v1beta1.IdentityPlatformTenantOAuthIDPConfig, opts v1.UpdateOptions) (*v1beta1.IdentityPlatformTenantOAuthIDPConfig, error)
	UpdateStatus(ctx context.Context, identityPlatformTenantOAuthIDPConfig *v1beta1.IdentityPlatformTenantOAuthIDPConfig, opts v1.UpdateOptions) (*v1beta1.IdentityPlatformTenantOAuthIDPConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.IdentityPlatformTenantOAuthIDPConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.IdentityPlatformTenantOAuthIDPConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.IdentityPlatformTenantOAuthIDPConfig, err error)
	IdentityPlatformTenantOAuthIDPConfigExpansion
}

// identityPlatformTenantOAuthIDPConfigs implements IdentityPlatformTenantOAuthIDPConfigInterface
type identityPlatformTenantOAuthIDPConfigs struct {
	client rest.Interface
	ns     string
}

// newIdentityPlatformTenantOAuthIDPConfigs returns a IdentityPlatformTenantOAuthIDPConfigs
func newIdentityPlatformTenantOAuthIDPConfigs(c *IdentityplatformV1beta1Client, namespace string) *identityPlatformTenantOAuthIDPConfigs {
	return &identityPlatformTenantOAuthIDPConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the identityPlatformTenantOAuthIDPConfig, and returns the corresponding identityPlatformTenantOAuthIDPConfig object, and an error if there is any.
func (c *identityPlatformTenantOAuthIDPConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.IdentityPlatformTenantOAuthIDPConfig, err error) {
	result = &v1beta1.IdentityPlatformTenantOAuthIDPConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("identityplatformtenantoauthidpconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IdentityPlatformTenantOAuthIDPConfigs that match those selectors.
func (c *identityPlatformTenantOAuthIDPConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.IdentityPlatformTenantOAuthIDPConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.IdentityPlatformTenantOAuthIDPConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("identityplatformtenantoauthidpconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested identityPlatformTenantOAuthIDPConfigs.
func (c *identityPlatformTenantOAuthIDPConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("identityplatformtenantoauthidpconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a identityPlatformTenantOAuthIDPConfig and creates it.  Returns the server's representation of the identityPlatformTenantOAuthIDPConfig, and an error, if there is any.
func (c *identityPlatformTenantOAuthIDPConfigs) Create(ctx context.Context, identityPlatformTenantOAuthIDPConfig *v1beta1.IdentityPlatformTenantOAuthIDPConfig, opts v1.CreateOptions) (result *v1beta1.IdentityPlatformTenantOAuthIDPConfig, err error) {
	result = &v1beta1.IdentityPlatformTenantOAuthIDPConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("identityplatformtenantoauthidpconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(identityPlatformTenantOAuthIDPConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a identityPlatformTenantOAuthIDPConfig and updates it. Returns the server's representation of the identityPlatformTenantOAuthIDPConfig, and an error, if there is any.
func (c *identityPlatformTenantOAuthIDPConfigs) Update(ctx context.Context, identityPlatformTenantOAuthIDPConfig *v1beta1.IdentityPlatformTenantOAuthIDPConfig, opts v1.UpdateOptions) (result *v1beta1.IdentityPlatformTenantOAuthIDPConfig, err error) {
	result = &v1beta1.IdentityPlatformTenantOAuthIDPConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("identityplatformtenantoauthidpconfigs").
		Name(identityPlatformTenantOAuthIDPConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(identityPlatformTenantOAuthIDPConfig).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *identityPlatformTenantOAuthIDPConfigs) UpdateStatus(ctx context.Context, identityPlatformTenantOAuthIDPConfig *v1beta1.IdentityPlatformTenantOAuthIDPConfig, opts v1.UpdateOptions) (result *v1beta1.IdentityPlatformTenantOAuthIDPConfig, err error) {
	result = &v1beta1.IdentityPlatformTenantOAuthIDPConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("identityplatformtenantoauthidpconfigs").
		Name(identityPlatformTenantOAuthIDPConfig.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(identityPlatformTenantOAuthIDPConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the identityPlatformTenantOAuthIDPConfig and deletes it. Returns an error if one occurs.
func (c *identityPlatformTenantOAuthIDPConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("identityplatformtenantoauthidpconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *identityPlatformTenantOAuthIDPConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("identityplatformtenantoauthidpconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched identityPlatformTenantOAuthIDPConfig.
func (c *identityPlatformTenantOAuthIDPConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.IdentityPlatformTenantOAuthIDPConfig, err error) {
	result = &v1beta1.IdentityPlatformTenantOAuthIDPConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("identityplatformtenantoauthidpconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}