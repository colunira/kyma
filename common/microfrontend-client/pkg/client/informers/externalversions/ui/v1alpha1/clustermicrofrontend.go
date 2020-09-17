// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	uiv1alpha1 "github.com/kyma-project/kyma/common/microfrontend-client/pkg/apis/ui/v1alpha1"
	versioned "github.com/kyma-project/kyma/common/microfrontend-client/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kyma-project/kyma/common/microfrontend-client/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kyma-project/kyma/common/microfrontend-client/pkg/client/listers/ui/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterMicroFrontendInformer provides access to a shared informer and lister for
// ClusterMicroFrontends.
type ClusterMicroFrontendInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterMicroFrontendLister
}

type clusterMicroFrontendInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterMicroFrontendInformer constructs a new informer for ClusterMicroFrontend type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterMicroFrontendInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterMicroFrontendInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterMicroFrontendInformer constructs a new informer for ClusterMicroFrontend type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterMicroFrontendInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.UiV1alpha1().ClusterMicroFrontends().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.UiV1alpha1().ClusterMicroFrontends().Watch(context.TODO(), options)
			},
		},
		&uiv1alpha1.ClusterMicroFrontend{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterMicroFrontendInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterMicroFrontendInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterMicroFrontendInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&uiv1alpha1.ClusterMicroFrontend{}, f.defaultInformer)
}

func (f *clusterMicroFrontendInformer) Lister() v1alpha1.ClusterMicroFrontendLister {
	return v1alpha1.NewClusterMicroFrontendLister(f.Informer().GetIndexer())
}
