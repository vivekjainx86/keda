/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/Azure/Kore/pkg/apis/kesc/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ScaledObjectLister helps list ScaledObjects.
type ScaledObjectLister interface {
	// List lists all ScaledObjects in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ScaledObject, err error)
	// ScaledObjects returns an object that can list and get ScaledObjects.
	ScaledObjects(namespace string) ScaledObjectNamespaceLister
	ScaledObjectListerExpansion
}

// scaledObjectLister implements the ScaledObjectLister interface.
type scaledObjectLister struct {
	indexer cache.Indexer
}

// NewScaledObjectLister returns a new ScaledObjectLister.
func NewScaledObjectLister(indexer cache.Indexer) ScaledObjectLister {
	return &scaledObjectLister{indexer: indexer}
}

// List lists all ScaledObjects in the indexer.
func (s *scaledObjectLister) List(selector labels.Selector) (ret []*v1alpha1.ScaledObject, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScaledObject))
	})
	return ret, err
}

// ScaledObjects returns an object that can list and get ScaledObjects.
func (s *scaledObjectLister) ScaledObjects(namespace string) ScaledObjectNamespaceLister {
	return scaledObjectNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ScaledObjectNamespaceLister helps list and get ScaledObjects.
type ScaledObjectNamespaceLister interface {
	// List lists all ScaledObjects in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ScaledObject, err error)
	// Get retrieves the ScaledObject from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ScaledObject, error)
	ScaledObjectNamespaceListerExpansion
}

// scaledObjectNamespaceLister implements the ScaledObjectNamespaceLister
// interface.
type scaledObjectNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ScaledObjects in the indexer for a given namespace.
func (s scaledObjectNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ScaledObject, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScaledObject))
	})
	return ret, err
}

// Get retrieves the ScaledObject from the indexer for a given namespace and name.
func (s scaledObjectNamespaceLister) Get(name string) (*v1alpha1.ScaledObject, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("scaledobject"), name)
	}
	return obj.(*v1alpha1.ScaledObject), nil
}
