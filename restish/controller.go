// Restish Controller. A Restish Controller is a Dispatcher, used to dispatch Resources to the appropriate action handler
// The four standard actions are the usual CRUD methods create, read, update and delete. Each action will be dispatched
// to the corresponding controller function.
package restish

type Controller interface {
	Create(resource *Resource) (*Resource, StatusCode)
	Read(resource *Resource) (*Resource, StatusCode)
	Update(resource *Resource) (*Resource, StatusCode)
	Delete(resource *Resource) (*Resource, StatusCode)
}

// Abstract controller. Implements default handling for all CRUD methods.
// Not sure this is the right way to do this.
type ControllerAbstract struct {
}

func (_ ControllerAbstract) Create(resource *Resource) (*Resource, StatusCode) {
	return resource, StatusNotImplemented
}

func (_ ControllerAbstract) Read(resource *Resource) (*Resource, StatusCode) {
	return resource, StatusNotImplemented
}

func (_ ControllerAbstract) Update(resource *Resource) (*Resource, StatusCode) {
	return resource, StatusNotImplemented
}

func (_ ControllerAbstract) Delete(resource *Resource) (*Resource, StatusCode) {
	return resource, StatusNotImplemented
}
