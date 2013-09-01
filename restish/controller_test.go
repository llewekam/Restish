package restish

import (
	"github.com/llewekam/assertion"
	"testing"
)

// Controller Stub object
type ControllerStub struct{
	assertion.Stuber
}

func (stub *ControllerStub) Create(resource *Resource) (*Resource, StatusCode) {
	stub.Stuber.Register("Create")

	resource.Properties = map[string]string{
		"action": "Create",
	}
	return resource, StatusTeapot
}
func (stub *ControllerStub) Read(resource *Resource) (*Resource, StatusCode) {
	stub.Stuber.Register("Read")

	resource.Properties = map[string]string{
		"action": "Read",
	}
	return resource, StatusTeapot
}
func (stub *ControllerStub) Update(resource *Resource) (*Resource, StatusCode) {
	stub.Stuber.Register("Update")

	resource.Properties = map[string]string{
		"action": "Update",
	}
	return resource, StatusTeapot
}
func (stub *ControllerStub) Delete(resource *Resource) (*Resource, StatusCode) {
	stub.Stuber.Register("Delete")

	resource.Properties = map[string]string{
		"action": "Delete",
	}
	return resource, StatusTeapot
}

// Test the default controller object. Must ensure we always get a Not Implemented Status as a response. For all action calls
func TestController (test *testing.T) {
	var status StatusCode

	assert := assertion.Assertion{test}
	controller := ControllerAbstract{}
	resource := NewResource("/url")

	resource, status = controller.Create(resource)
	assert.AssertEqual(StatusNotImplemented, status, "Expected Status Not Implemented")

	resource, status = controller.Read(resource)
	assert.AssertEqual(StatusNotImplemented, status, "Expected Status Not Implemented")

	resource, status = controller.Update(resource)
	assert.AssertEqual(StatusNotImplemented, status, "Expected Status Not Implemented")

	resource, status = controller.Delete(resource)
	assert.AssertEqual(StatusNotImplemented, status, "Expected Status Not Implemented")
}
