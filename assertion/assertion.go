// Test assertion package. Just to make test code a bit more concise.
package assertion

import (
	"reflect"
	"testing"
)

type Assertion struct {
	*testing.T
}

func (assert *Assertion) AssertEqual(expected interface{}, actual interface{}, failMessage string) {
	if !reflect.DeepEqual(expected, actual) {
		assert.Error("Expected does not match Actual: ", failMessage)
	}
}

func (assert *Assertion) AssertEqualNow(actual interface{}, expected interface{}, failMessage string) {
	if !reflect.DeepEqual(expected, actual) {
		assert.Log("Expected does not match Actual: ", failMessage)
		assert.FailNow()
	}
}

func (assert *Assertion) AssertStubCall(actual Stuber, name string, callCount int) {
	call := actual.CallCount(name)

	if call.callCount != callCount {
		assert.Error("Function:", name, "Expected call count (", callCount, ") does not match actual (", call.callCount, ")")
	}
}
