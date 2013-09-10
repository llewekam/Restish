// restish
//
// Resource Tests

package restish

import (
	"github.com/llewekam/assertion"
	"testing"
)

// Test NewResource(). Must return a Resource with all of the appropriate default values.
func TestNewResource(test *testing.T) {
	assert := assertion.Assertion{test}

	resource := NewResource("/url")

	assert.AssertEqual("/url", resource.Self().Href, "Resource Self does not match expected URL")
	assert.AssertEqual(StatusOk, resource.Status, "Resource does not have a default status of OK")
}

// Test resource.Self() returns the Link to self when a link of rel "self" is added
func TestResourceSelf(test *testing.T) {
	assert := assertion.Assertion{test}

	expected := Link{"self", "/url", "application/vnd"}

	resource := new(Resource)
	resource.AddLink("self", "/url", "application/vnd")

	assert.AssertEqual(resource.Self().Href, expected.Href, "Href does not match")
	assert.AssertEqual(resource.Self().Rel, expected.Rel, "Rel does not match")
	assert.AssertEqual(resource.Self().Type, expected.Type, "Type does not match")

}
