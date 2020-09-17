// Code generated by injection-gen. DO NOT EDIT.

package httpsource

import (
	"context"

	v1alpha1 "github.com/kyma-project/kyma/components/event-sources/client/generated/informer/externalversions/sources/v1alpha1"
	factory "github.com/kyma-project/kyma/components/event-sources/client/generated/injection/informers/factory"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterInformer(withInformer)
}

// Key is used for associating the Informer inside the context.Context.
type Key struct{}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := factory.Get(ctx)
	inf := f.Sources().V1alpha1().HTTPSources()
	return context.WithValue(ctx, Key{}, inf), inf.Informer()
}

// Get extracts the typed informer from the context.
func Get(ctx context.Context) v1alpha1.HTTPSourceInformer {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch github.com/kyma-project/kyma/components/event-sources/client/generated/informer/externalversions/sources/v1alpha1.HTTPSourceInformer from context.")
	}
	return untyped.(v1alpha1.HTTPSourceInformer)
}
