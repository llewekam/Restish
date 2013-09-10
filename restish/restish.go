// Restish package.
package restish

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
	StatusTeapot         = StatusCode{418, "I'm a teapot"}
)

var (
	// All the Endpoint dispatchers
	dispatchers     = map[string]Dispatcher{}
	defaultDispatch Dispatcher
)

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
	route := resource.Self()

	if nil != dispatchers[route.Href] {
		dispatch = dispatchers[route.Href]
	} else {
		if nil != defaultDispatch {
			dispatch = defaultDispatch
		} else {
			error = new(DispatchError)
		}
	}

	return
}
