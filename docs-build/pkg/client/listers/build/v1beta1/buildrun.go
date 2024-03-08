// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/shipwright-io/build/pkg/apis/build/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BuildRunLister helps list BuildRuns.
// All objects returned here must be treated as read-only.
type BuildRunLister interface {
	// List lists all BuildRuns in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.BuildRun, err error)
	// BuildRuns returns an object that can list and get BuildRuns.
	BuildRuns(namespace string) BuildRunNamespaceLister
	BuildRunListerExpansion
}

// buildRunLister implements the BuildRunLister interface.
type buildRunLister struct {
	indexer cache.Indexer
}

// NewBuildRunLister returns a new BuildRunLister.
func NewBuildRunLister(indexer cache.Indexer) BuildRunLister {
	return &buildRunLister{indexer: indexer}
}

// List lists all BuildRuns in the indexer.
func (s *buildRunLister) List(selector labels.Selector) (ret []*v1beta1.BuildRun, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.BuildRun))
	})
	return ret, err
}

// BuildRuns returns an object that can list and get BuildRuns.
func (s *buildRunLister) BuildRuns(namespace string) BuildRunNamespaceLister {
	return buildRunNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BuildRunNamespaceLister helps list and get BuildRuns.
// All objects returned here must be treated as read-only.
type BuildRunNamespaceLister interface {
	// List lists all BuildRuns in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.BuildRun, err error)
	// Get retrieves the BuildRun from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.BuildRun, error)
	BuildRunNamespaceListerExpansion
}

// buildRunNamespaceLister implements the BuildRunNamespaceLister
// interface.
type buildRunNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BuildRuns in the indexer for a given namespace.
func (s buildRunNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.BuildRun, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.BuildRun))
	})
	return ret, err
}

// Get retrieves the BuildRun from the indexer for a given namespace and name.
func (s buildRunNamespaceLister) Get(name string) (*v1beta1.BuildRun, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("buildrun"), name)
	}
	return obj.(*v1beta1.BuildRun), nil
}
