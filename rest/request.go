package rest

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
)

// Request allows for building up a request to a server in a chained fashion.
// Any errors are stored until the end of your call, so you only have to
// check once.
type Request struct {
	c *RESTClient

	// rateLimiter flowcontrol.RateLimiter
	// backoff BackoffManager
	timeout time.Duration

	// generic components accessible via method setters
	verb       string
	pathPrefix string
	subpath    string
	params     url.Values
	headers    http.Header

	// structural elements of the request that are part of the Kubernetes API conventions
	namespace    string
	namespaceSet bool
	resource     string
	resourceName string
	subresource  string

	// output
	err  error
	body io.Reader
}

// Result contains the result of calling Request.Do().
type Result struct {
	body        []byte
	contentType string
	err         error
	statusCode  int

	decoder runtime.Decoder
}

// Name sets the name of a resource to access (<resource>/[ns/<namespace>/]<name>)
func (r *Request) Name(resourceName string) *Request {
	/*
		if r.err != nil {
			return r
		}
		if len(resourceName) == 0 {
			r.err = fmt.Errorf("resource name may not be empty")
			return r
		}
		if len(r.resourceName) != 0 {
			r.err = fmt.Errorf("resource name already set to %q, cannot change to %q", r.resourceName, resourceName)
			return r
		}
		if msgs := IsValidPathSegmentName(resourceName); len(msgs) != 0 {
			r.err = fmt.Errorf("invalid resource name %q: %v", resourceName, msgs)
			return r
		}
		r.resourceName = resourceName
		return r
	*/
	return nil
}

// Namespace applies the namespace scope to a request (<resource>/[ns/<namespace>/]<name>)
func (r *Request) Namespace(namespace string) *Request {
	/*
		if r.err != nil {
			return r
		}
		if r.namespaceSet {
			r.err = fmt.Errorf("namespace already set to %q, cannot change to %q", r.namespace, namespace)
			return r
		}
		if msgs := IsValidPathSegmentName(namespace); len(msgs) != 0 {
			r.err = fmt.Errorf("invalid namespace %q: %v", namespace, msgs)
			return r
		}
		r.namespaceSet = true
		r.namespace = namespace
		return r
	*/
	return nil
}

// Resource sets the resource to access (<resource>/[ns/<namespace>/]<name>)
func (r *Request) Resource(resource string) *Request {
	/*
		if r.err != nil {
			return r
		}
		if len(r.resource) != 0 {
			r.err = fmt.Errorf("resource already set to %q, cannot change to %q", r.resource, resource)
			return r
		}
		if msgs := IsValidPathSegmentName(resource); len(msgs) != 0 {
			r.err = fmt.Errorf("invalid resource %q: %v", resource, msgs)
			return r
		}
		r.resource = resource
		return r
	*/
	return nil
}

func (r *Request) Do(ctx context.Context) Result {
	/*
		var result Result
		err := r.request(ctx, func(req *http.Request, resp *http.Response) {
			result = r.transformResponse(resp, req)
		})
		if err != nil {
			return Result{err: err}
		}
		return result
	*/
	return Result{}
}

func (r Result) Into(obj runtime.Object) error {
	/*
		if r.err != nil {
			// Check whether the result has a Status object in the body and prefer that.
			return r.Error()
		}
		if r.decoder == nil {
			return fmt.Errorf("serializer for %s doesn't exist", r.contentType)
		}
		if len(r.body) == 0 {
			return fmt.Errorf("0-length response with status code: %d and content type: %s",
				r.statusCode, r.contentType)
		}

		out, _, err := r.decoder.Decode(r.body, nil, obj)
		if err != nil || out == obj {
			return err
		}
		// if a different object is returned, see if it is Status and avoid double decoding
		// the object.
		switch t := out.(type) {
		case *metav1.Status:
			// any status besides StatusSuccess is considered an error.
			if t.Status != metav1.StatusSuccess {
				return errors.FromObject(t)
			}
		}
		return nil
	*/
	return nil
}

// VersionedParams will take the provided object, serialize it to a map[string][]string using the
// implicit RESTClient API version and the default parameter codec, and then add those as parameters
// to the request. Use this to provide versioned query parameters from client libraries.
// VersionedParams will not write query parameters that have omitempty set and are empty. If a
// parameter has already been set it is appended to (Params and VersionedParams are additive).
func (r *Request) VersionedParams(obj runtime.Object, codec runtime.ParameterCodec) *Request {
	// return r.SpecificallyVersionedParams(obj, codec, r.c.content.GroupVersion)
	return nil
}

// func (r *Request) SpecificallyVersionedParams(obj runtime.Object, codec runtime.ParameterCodec, version schema.GroupVersion) *Request {
// 	if r.err != nil {
// 		return r
// 	}
// 	params, err := codec.EncodeParameters(obj, version)
// 	if err != nil {
// 		r.err = err
// 		return r
// 	}
// 	for k, v := range params {
// 		if r.params == nil {
// 			r.params = make(url.Values)
// 		}
// 		r.params[k] = append(r.params[k], v...)
// 	}
// 	return r
// }
