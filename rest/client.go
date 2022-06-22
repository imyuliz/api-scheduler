package rest

import (
	"net/http"
	"net/url"

	"github.com/imyuliz/api-scheduler/pkg/types"
)

type Interface interface {
	// GetRateLimiter() flowcontrol.RateLimiter
	Verb(verb string) *Request
	Post() *Request
	Put() *Request
	Patch(pt types.PatchType) *Request
	Get() *Request
	Delete() *Request
	// APIVersion() schema.GroupVersion
}

type RESTClient struct {
	// base is the root URL for all invocations of the client
	base *url.URL
	// versionedAPIPath is a path segment connecting the base URL to the resource root
	versionedAPIPath string

	// content describes how a RESTClient encodes and decodes responses.
	// content ClientContentConfig

	// creates BackoffManager that is passed to requests.
	// createBackoffMgr func() BackoffManager

	// rateLimiter is shared among all requests created by this client unless specifically
	// overridden.
	// rateLimiter flowcontrol.RateLimiter

	// Set specific behavior of the client.  If not set http.DefaultClient will be used.
	Client *http.Client
}
