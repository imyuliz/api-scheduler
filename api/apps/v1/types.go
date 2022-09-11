package v1

import (
	metav1 "github.com/imyuliz/api-scheduler/pkg/apis/meta/v1"
)

type Flow struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object metadata.
	// +optional
	metav1.ObjectMeta `json:",inline"`
	Spec              FlowSpec   `json:"spec"`
	Status            FlowStatus `json:"status"`
}

type FlowSpec struct {
	Storage  DataStorage      `json:"dataStorage,omitempty"`
	Template StepTemplateSpec `json:"template,omitempty"`

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

	// The number of old ReplicaSets to retain to allow rollback.
	// This is a pointer to distinguish between explicit zero and not specified.
	// Defaults to 10.
	// +optional
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty"`
	// Indicates that the deployment is paused.
	// +optional
	Paused bool `json:"paused,omitempty" protobuf:"varint,7,opt,name=paused"`
}

type FlowStatus struct {
	DataStoragePath string `json:"dataStoragePath,omitempty"`
}

type StepTemplateSpec struct {
	metav1.ObjectMeta `json:",inline"`
	Spec              StepSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type StepSpec struct {
	Handlers []Requests `json:"handlers,omitempty"`
}

type Requests []Request

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
	// 如果报错是否跳过
	NoError bool
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
	URL string `json:"path,omitempty"`
	// +optional
	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty"`
	// Custom query to setin the request
	HTTPQueries []HTTPQuery `json:"httpQueries,omitempty"`
	HTTPBody    HTTPBody    `json:"httpBody,omitempty"`
}

type HTTPDeleteAction struct {
	Address string `json:"address,omitempty"`
	// +optional
	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty"`

	// Custom query to setin the request
	HTTPQueries []HTTPQuery `json:"HttpQueries,omitempty"`
}

type HTTPPutAction struct {
	Address string `json:"address,omitempty"`
	// +optional
	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty"`

	// Custom query to setin the request
	HTTPQueries []HTTPQuery `json:"HttpQueries,omitempty"`
}

type HTTPPatchAction struct {
	Address string `json:"address,omitempty"`
	// +optional
	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty"`

	// Custom query to setin the request
	HTTPQueries []HTTPQuery `json:"HttpQueries,omitempty"`
}

// HTTPGetAction describes an action based on HTTP Get requests.
type HTTPGetAction struct {
	Address string `json:"address,omitempty"`
	// +optional
	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty"`

	// Custom query to setin the request
	HTTPQueries []HTTPQuery `json:"HttpQueries,omitempty"`
}

// HTTPHeader describes a custom header to be used in HTTP probes
type HTTPHeader struct {
	// The header field name
	Name string `json:"name,omitempty"`
	// The header field value
	Value    string `json:"value,omitempty"`
	Required bool   `json:"required,omitempty"`
}

type HTTPQuery struct {
	// The header field name
	Name string `json:"name,omitempty"`
	// The header field value
	Value    string `json:"value,omitempty"`
	Required bool   `json:"required,omitempty"`
}

type HTTPBody struct {
	Name     string      `json:"name,omitempty"`
	Value    interface{} `json:"value,omitempty"`
	Required bool        `json:"required,omitempty"`
}
