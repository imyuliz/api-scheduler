package v1

import (
	"context"

	v1 "github.com/imyuliz/api-scheduler/api/apps/v1"
	metav1 "github.com/imyuliz/api-scheduler/pkg/apis/meta/v1"
	"github.com/imyuliz/api-scheduler/pkg/runtime"

	"github.com/imyuliz/api-scheduler/rest"
)

type FlowsGetter interface {
	Flows(namespace string) FlowInterface
}

type FlowInterface interface {
	Create(ctx context.Context, flow *v1.Flow, opts metav1.CreateOptions) (*v1.Flow, error)
	Update(ctx context.Context, flow *v1.Flow, opts metav1.UpdateOptions) (*v1.Flow, error)
	UpdateStatus(ctx context.Context, flow *v1.Flow, opts metav1.UpdateOptions) (*v1.Flow, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Flow, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.Flow, error)
	// Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	// Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Deployment, err error)
	// GetScale(ctx context.Context, deploymentName string, options metav1.GetOptions) (*autoscalingv1.Scale, error)
	// UpdateScale(ctx context.Context, deploymentName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (*autoscalingv1.Scale, error)

}

type flow struct {
	client rest.Interface
	ns     string
}

func (f *flow) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Flow, error) {
	// runtime.ParameterCodec(scheme.ParameterCodec)
	result := &v1.Flow{}
	err := f.client.Get().
		Namespace(f.ns).
		Resource("flows").
		Name(name).
		VersionedParams(&opts, runtime.NewParameterCodec(runtime.NewScheme())).
		Do(ctx).
		Into(result)
	return result, err
}
