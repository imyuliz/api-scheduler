package v1

type CoreV1Interface interface {
	// FlowsGetter
}

// type FlowsGetter interface {
// 	Pods(namespace string) FlowInterface
// }

// type FlowInterface interface {
// 	Create(ctx context.Context, pod *v1.Pod, opts metav1.CreateOptions) (*v1.Pod, error)
// 	Update(ctx context.Context, pod *v1.Pod, opts metav1.UpdateOptions) (*v1.Pod, error)
// 	UpdateStatus(ctx context.Context, pod *v1.Pod, opts metav1.UpdateOptions) (*v1.Pod, error)
// 	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
// 	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
// 	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Pod, error)
// 	List(ctx context.Context, opts metav1.ListOptions) (*v1.PodList, error)
// 	// Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
// 	// Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Pod, err error)
// 	// GetEphemeralContainers(ctx context.Context, podName string, options metav1.GetOptions) (*v1.EphemeralContainers, error)
// 	// UpdateEphemeralContainers(ctx context.Context, podName string, ephemeralContainers *v1.EphemeralContainers, opts metav1.UpdateOptions) (*v1.EphemeralContainers, error)
// }
