package runtime

import (
	"net/url"

	"github.com/imyuliz/api-scheduler/pkg/runtime/schema"
)

// Object interface must be supported by all API types registered with Scheme. Since objects in a scheme are
// expected to be serialized to the wire, the interface an Object must provide to the Scheme allows
// serializers to set the kind, version, and group the object is represented as. An Object may choose
// to return a no-op ObjectKindAccessor in cases where it is not expected to be serialized.
type Object interface {
	GetObjectKind() schema.ObjectKind
	DeepCopyObject() Object
}

// Decoder attempts to load an object from data.
type Decoder interface {
	// Decode attempts to deserialize the provided data using either the innate typing of the scheme or the
	// default kind, group, and version provided. It returns a decoded object as well as the kind, group, and
	// version from the serialized data, or an error. If into is non-nil, it will be used as the target type
	// and implementations may choose to use it rather than reallocating an object. However, the object is not
	// guaranteed to be populated. The returned object is not guaranteed to match into. If defaults are
	// provided, they are applied to the data by default. If no defaults or partial defaults are provided, the
	// type of the into may be used to guide conversion decisions.
	Decode(data []byte, defaults *schema.GroupVersionKind, into Object) (Object, *schema.GroupVersionKind, error)
}

// ParameterCodec defines methods for serializing and deserializing API objects to url.Values and
// performing any necessary conversion. Unlike the normal Codec, query parameters are not self describing
// and the desired version must be specified.
type ParameterCodec interface {
	// DecodeParameters takes the given url.Values in the specified group version and decodes them
	// into the provided object, or returns an error.
	DecodeParameters(parameters url.Values, from schema.GroupVersion, into Object) error
	// EncodeParameters encodes the provided object as query parameters or returns an error.
	EncodeParameters(obj Object, to schema.GroupVersion) (url.Values, error)
}
