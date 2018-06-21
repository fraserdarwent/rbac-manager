// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/reactiveops/rbac-manager/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// RBACDefinitions returns a RBACDefinitionInformer.
	RBACDefinitions() RBACDefinitionInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// RBACDefinitions returns a RBACDefinitionInformer.
func (v *version) RBACDefinitions() RBACDefinitionInformer {
	return &rBACDefinitionInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
