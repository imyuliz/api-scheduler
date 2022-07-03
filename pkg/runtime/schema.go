package runtime

import (
	"reflect"

	"github.com/imyuliz/api-scheduler/pkg/runtime/schema"
)

type Schema struct {
	// versionMap allows one to figure out the go type of an object with
	// the given version and name.
	gvkToType map[schema.GroupVersionKind]reflect.Type

	// typeToGroupVersion allows one to find metadata for a given go object.
	// The reflect.Type we index by should *not* be a pointer.
	typeToGVK map[reflect.Type][]schema.GroupVersionKind

	// unversionedTypes are transformed without conversion in ConvertToVersion.
	unversionedTypes map[reflect.Type]schema.GroupVersionKind
}

// scheme is the registry for the common types that adhere to the meta v1 API spec.
var scheme = NewScheme()

func NewScheme() *Schema {
	return &Schema{
		gvkToType:        map[schema.GroupVersionKind]reflect.Type{},
		typeToGVK:        map[reflect.Type][]schema.GroupVersionKind{},
		unversionedTypes: map[reflect.Type]schema.GroupVersionKind{},
	}
}

// var ParameterCodec = NewParameterCodec(NewScheme())
