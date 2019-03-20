/*
Copyright The Flagger Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	appmeshv1alpha1 "github.com/stefanprodan/flagger/pkg/client/clientset/versioned/typed/appmesh/v1alpha1"
	flaggerv1alpha3 "github.com/stefanprodan/flagger/pkg/client/clientset/versioned/typed/flagger/v1alpha3"
	networkingv1alpha3 "github.com/stefanprodan/flagger/pkg/client/clientset/versioned/typed/istio/v1alpha3"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AppmeshV1alpha1() appmeshv1alpha1.AppmeshV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Appmesh() appmeshv1alpha1.AppmeshV1alpha1Interface
	FlaggerV1alpha3() flaggerv1alpha3.FlaggerV1alpha3Interface
	// Deprecated: please explicitly pick a version if possible.
	Flagger() flaggerv1alpha3.FlaggerV1alpha3Interface
	NetworkingV1alpha3() networkingv1alpha3.NetworkingV1alpha3Interface
	// Deprecated: please explicitly pick a version if possible.
	Networking() networkingv1alpha3.NetworkingV1alpha3Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	appmeshV1alpha1    *appmeshv1alpha1.AppmeshV1alpha1Client
	flaggerV1alpha3    *flaggerv1alpha3.FlaggerV1alpha3Client
	networkingV1alpha3 *networkingv1alpha3.NetworkingV1alpha3Client
}

// AppmeshV1alpha1 retrieves the AppmeshV1alpha1Client
func (c *Clientset) AppmeshV1alpha1() appmeshv1alpha1.AppmeshV1alpha1Interface {
	return c.appmeshV1alpha1
}

// Deprecated: Appmesh retrieves the default version of AppmeshClient.
// Please explicitly pick a version.
func (c *Clientset) Appmesh() appmeshv1alpha1.AppmeshV1alpha1Interface {
	return c.appmeshV1alpha1
}

// FlaggerV1alpha3 retrieves the FlaggerV1alpha3Client
func (c *Clientset) FlaggerV1alpha3() flaggerv1alpha3.FlaggerV1alpha3Interface {
	return c.flaggerV1alpha3
}

// Deprecated: Flagger retrieves the default version of FlaggerClient.
// Please explicitly pick a version.
func (c *Clientset) Flagger() flaggerv1alpha3.FlaggerV1alpha3Interface {
	return c.flaggerV1alpha3
}

// NetworkingV1alpha3 retrieves the NetworkingV1alpha3Client
func (c *Clientset) NetworkingV1alpha3() networkingv1alpha3.NetworkingV1alpha3Interface {
	return c.networkingV1alpha3
}

// Deprecated: Networking retrieves the default version of NetworkingClient.
// Please explicitly pick a version.
func (c *Clientset) Networking() networkingv1alpha3.NetworkingV1alpha3Interface {
	return c.networkingV1alpha3
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.appmeshV1alpha1, err = appmeshv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.flaggerV1alpha3, err = flaggerv1alpha3.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.networkingV1alpha3, err = networkingv1alpha3.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.appmeshV1alpha1 = appmeshv1alpha1.NewForConfigOrDie(c)
	cs.flaggerV1alpha3 = flaggerv1alpha3.NewForConfigOrDie(c)
	cs.networkingV1alpha3 = networkingv1alpha3.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.appmeshV1alpha1 = appmeshv1alpha1.New(c)
	cs.flaggerV1alpha3 = flaggerv1alpha3.New(c)
	cs.networkingV1alpha3 = networkingv1alpha3.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
