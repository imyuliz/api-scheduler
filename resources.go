package main

import "k8s.io/apimachinery/pkg/util/intstr"

type Scheduler struct {
	TypeMeta `json:",inline"`
	// Standard object metadata.
	// +optional
	ObjectMeta `json:",inline"`
	Spec       SchedulerSpec   `json:"spec"`
	Status     SchedulerStatus `json:"status"`
}

type SchedulerSpec struct {
	// The path is http url
	Path string `json:"path,omitempty"`

	// HTTP Method
	Method string `json:"method,omitempty"`

	DataStorage DataStorage `json:"dataStorage,omitempty"`

	Template StepTemplateSpec `json:"template,omitempty"`

	// The deployment strategy to use to replace existing pods with new ones.
	// +optional
	// +patchStrategy=retainKeys
	Strategy SchedulerStrategy `json:"strategy,omitempty"`

	// Number of desired pods. This is a pointer to distinguish between explicit
	// zero and not specified. Defaults to 1.
	// +optional
	RetryLimit *int32 `json:"retryLimit,omitempty"`

	// Restart policy for all containers within the pod.
	// One of Always, OnFailure, Never.
	// Default to Always.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	// +optional
	RetryPolicy RetryPolicy `json:"retryPolicy,omitempty"`

	// Specifies the duration in seconds relative to the startTime that the job may be active
	// before the system tries to terminate it; value must be positive integer
	// +optional
	ActiveDeadlineSeconds *int64 `json:"activeDeadlineSeconds,omitempty" protobuf:"varint,3,opt,name=activeDeadlineSeconds"`

	// The number of old ReplicaSets to retain to allow rollback.
	// This is a pointer to distinguish between explicit zero and not specified.
	// Defaults to 10.
	// +optional
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty"`
	// Indicates that the deployment is paused.
	// +optional
	Paused bool `json:"paused,omitempty" protobuf:"varint,7,opt,name=paused"`
}

type SchedulerStatus struct {
	DataStoragePath string `json:"dataStoragePath,omitempty"`
}

type StepTemplateSpec struct {
	ObjectMeta `json:",inline"`
	Spec       StepSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type StepSpec struct {
	InitHandlers     []Request `json:"initHandlers,omitempty"`
	Handlers         []Request `json:"handlers,omitempty"`
	CallbackHandlers []Request `json:"callbackHandlers,omitempty"`
}

// 如何解决数据存储差异化的问题:
// 1. 请求来源:(仅入口参数接口)
// 	1. query  -> req_q_key
// 	2. path   -> req_p_key
// 	3. header -> req_h_key
// 	4. body   -> req_body_key

// 2. 响应来源:
// 	1. header -> resp_h_key
// 	2. body   -> resp_body_key

type DataStorage struct {
	Enable         *bool     `json:"enable,omitempty"`
	RequestPrefix  string    `json:"requestPrefix,omitempty"`
	ResponsePrefix string    `json:"requestPrefix,omitempty"`
	RemoteWrite    []Address `json:"requestPrefix,omitempty"`
}

type Address struct {
	URL string `json:"url,omitempty"`
}

type SchedulerStrategy struct {
	// Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
	// +optional
	Type          SchedulerStrategyType   `json:"type,omitempty"`
	RollingUpdate *RollingUpdateScheduler `json:"rollingUpdate,omitempty"`
}

type RollingUpdateScheduler struct {
}

type SchedulerStrategyType string

const (
	// Kill all existing pods before creating new ones.
	RecreateSchedulerStrategyType SchedulerStrategyType = "Recreate"

	// Replace the old ReplicaSets by new one using rolling update i.e gradually scale down the old ReplicaSets and scale up the new one.
	RollingUpdateSchedulerStrategyType SchedulerStrategyType = "RollingUpdate"
	RetrySchedulerStrategyType         SchedulerStrategyType = "Retry"
)

// RestartPolicy describes how the container should be restarted.
// Only one of the following restart policies may be specified.
// If none of the following policies is specified, the default one
// is RestartPolicyAlways.
type RetryPolicy string

const (
	RetryPolicyAlways    RetryPolicy = "Always"
	RetryPolicyOnFailure RetryPolicy = "OnFailure"
	RetryPolicyNever     RetryPolicy = "Never"
)

type Request struct {
	Name string `json:"name,omitempty"`

	Annotations map[string]string `json:"annotations,omitempty"`

	Conditions []Condition `json:"conditions,omitempty"`
	Handler
	// The priority value. Various system components use this field to find the
	// priority of the pod. When Priority Admission Controller is enabled, it
	// prevents users from setting this field. The admission controller populates
	// this field from PriorityClassName.
	// The higher the value, the higher the priority.
	// +optional
	Priority *int32 `json:"priority,omitempty"`
	// Number of seconds after which the probe times out.
	// Defaults to 1 second. Minimum value is 1.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	// Defaults to 3. Minimum value is 1.
	// +optional
	FailureThreshold int32 `json:"failureThreshold,omitempty"`
	// 失败时终止
	FailureTerminate *bool `json:"failureTerminate,omitempty"`
}

// A label selector requirement is a selector that contains values, a key, and an operator that
// relates the key and values.
type Condition struct {
	// key is the label key that the selector applies to.
	// +patchMergeKey=key
	// +patchStrategy=merge
	Key string `json:"key" patchStrategy:"merge" patchMergeKey:"key" protobuf:"bytes,1,opt,name=key"`
	// operator represents a key's relationship to a set of values.
	// Valid operators are In, NotIn, Exists and DoesNotExist.
	Operator LabelSelectorOperator `json:"operator"`
	// values is an array of string values. If the operator is In or NotIn,
	// the values array must be non-empty. If the operator is Exists or DoesNotExist,
	// the values array must be empty. This array is replaced during a strategic
	// merge patch.
	// +optional
	Value string `json:"value,omitempty"`
}

// A label selector operator is the set of operators that can be used in a selector requirement.
type LabelSelectorOperator string

const (
	LabelSelectorOpEqual        LabelSelectorOperator = "Equal"
	LabelSelectorOpNotEqual     LabelSelectorOperator = "NotEqual"
	LabelSelectorOpDoesNotExist LabelSelectorOperator = "DoesNotExist"
	LabelSelectorOpExists       LabelSelectorOperator = "Exists"
)

type Handler struct {
	// HTTPGet specifies the http request to perform.
	// +optional
	HTTPGet    *HTTPGetAction    `json:"httpGet,omitempty"`
	HTTPPost   *HTTPPostAction   `json:"httpPost,omitempty"`
	HTTPDelete *HTTPDeleteAction `json:"httpDelete,omitempty"`
	HTTPPut    *HTTPPutAction    `json:"httpPut,omitempty"`
	HTTPPatch  *HTTPPatchAction  `json:"httpPatch,omitempty"`
}

type HTTPPostAction struct {
}

type HTTPDeleteAction struct {
}

type HTTPPutAction struct {
}

type HTTPPatchAction struct {
}

// HTTPGetAction describes an action based on HTTP Get requests.
type HTTPGetAction struct {
	// Path to access on the HTTP server.
	// +optional
	Path string `json:"path,omitempty"`
	// Name or number of the port to access on the container.
	// Number must be in the range 1 to 65535.
	// Name must be an IANA_SVC_NAME.
	Port intstr.IntOrString `json:"port"`
	// Host name to connect to, defaults to the pod IP. You probably want to set
	// "Host" in httpHeaders instead.
	// +optional
	Host string `json:"host,omitempty"`
	// Scheme to use for connecting to the host.
	// Defaults to HTTP.
	// +optional
	Scheme URIScheme `json:"scheme,omitempty"`
	// Custom headers to set in the request. HTTP allows repeated headers.
	// +optional
	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty"`

	// Custom query to setin the request
	HTTPQueries []HTTPQuery `json:"HttpQueries,omitempty"`
}

// URIScheme identifies the scheme used for connection to a host for Get actions
type URIScheme string

const (
	// URISchemeHTTP means that the scheme used will be http://
	URISchemeHTTP URIScheme = "HTTP"
	// URISchemeHTTPS means that the scheme used will be https://
	URISchemeHTTPS URIScheme = "HTTPS"
)

// HTTPHeader describes a custom header to be used in HTTP probes
type HTTPHeader struct {
	// The header field name
	Name string `json:"name,omitempty"`
	// The header field value
	Value string `json:"value,omitempty"`
}

type HTTPQuery struct {
	// The header field name
	Name string `json:"name,omitempty"`
	// The header field value
	Value string `json:"value,omitempty"`
}

type HTTPBody struct {
	Name  string      `json:"name,omitempty"`
	Value interface{} `json:"value,omitempty"`
}
