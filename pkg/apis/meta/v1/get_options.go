package v1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (o *GetOptions) GetObjectKind() schema.ObjectKind {
	return nil

}
func (o *GetOptions) DeepCopyObject() v1.Object {
	return nil
}
