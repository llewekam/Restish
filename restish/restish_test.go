package restish

import (
	"github.com/llewekam/assertion"
	"testing"
)

// Test NewResource(). Must return a Resource with all of the appropriate default values.
func TestNewResource(test *testing.T) {
	assert := assertion.Assertion{test}

	resource := NewResource("/url")

	assert.AssertEqual("/url", resource.Self, "Resource Self does not match expected URL")
	assert.AssertEqual(StatusOk, resource.Status, "Resource does not have a default status of OK")
}

//func TestGetDispatch(test *testing.T) {
//	assert := assertion.Assertion{test}
//
//	resource := NewResource("/url")
//}
