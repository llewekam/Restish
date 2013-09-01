// Restish package.
package restish

// A REST Resource. A Resource can be dispatched to Handlers and Controllers for processing. A Resource is always
// returned from a successful Dispatch
type Resource struct {
	Self       string
	Properties map[string]string
	Status     StatusCode
}

// Contains the Resource Status Codes. Normally used to define the HTTP status code associated with the resource once
// processing is complete.
type StatusCode struct {
	Code   int
	Status string
}

const (
	ActionCreate = "create"
	ActionRead   = "read"
	ActionUpdate = "update"
	ActionDelete = "delete"
)

// HTTP Response Codes.
var (
	StatusOk             = StatusCode{200, "OK"}
	StatusCreated        = StatusCode{201, "Created"}
	StatusBadRequest     = StatusCode{400, "Bad Request"}
	StatusNotFound       = StatusCode{404, "Not Found"}
	StatusServerError    = StatusCode{500, "Internal Server Error"}
	StatusNotImplemented = StatusCode{501, "Not Implemented"}
	StatusTeapot = StatusCode{418, "I'm a teapot"}
)

var (
	// All the Endpoint dispatchers
	dispatchers = map[string]Dispatcher{}
	defaultDispatch Dispatcher
)

// Create a new Resource
func NewResource(self string) *Resource {
	var resource *Resource
	resource = new(Resource)
	resource.Self = self
	resource.Status = StatusOk

	return resource
}

// Add a route to the dispatch set.
func AddRoute(dispatcher Dispatcher, endpoint string) {
	dispatchers[endpoint] = dispatcher
}

// Set the error dispatcher. Used when dispatcher cannot be found
func SetDefaultDispatch(dispatch Dispatcher) {
	defaultDispatch = dispatch
}

// Find the dispatcher responsible for the provided resource
func GetDispatch(resource *Resource) (dispatch Dispatcher, error error) {
	route := resource.Self

	if nil != dispatchers[route] {
		dispatch = dispatchers[route]
	} else {
		if nil != defaultDispatch {
			dispatch = defaultDispatch
		} else {
			error = new(DispatchError)
		}
	}

	return
}


