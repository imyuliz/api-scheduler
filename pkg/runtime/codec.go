package runtime

import (
	"net/url"

	"github.com/imyuliz/api-scheduler/pkg/runtime/schema"
)

// NewParameterCodec creates a ParameterCodec capable of transforming url values into versioned objects and back.
func NewParameterCodec(scheme *Schema) ParameterCodec {
	// todo: covert
	return &parameterCodec{}
}

// parameterCodec implements conversion to and from query parameters and objects.
type parameterCodec struct {
}

// DecodeParameters converts the provided url.Values into an object of type From with the kind of into, and then
// converts that object to into (if necessary). Returns an error if the operation cannot be completed.
func (c *parameterCodec) DecodeParameters(parameters url.Values, from schema.GroupVersion, into Object) error {
	return nil
}

// EncodeParameters converts the provided object into the to version, then converts that object to url.Values.
// Returns an error if conversion is not possible.
func (c *parameterCodec) EncodeParameters(obj Object, to schema.GroupVersion) (url.Values, error) {
	return nil, nil
}
