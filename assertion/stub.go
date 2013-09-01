package assertion

// Interface for Stubs
type Stuber interface {
	CallCount(string) CallCount
	Register(string)
}

// Struct for containin the function call count data
type CallCount struct {
	funcName string
	callCount int
}

func NewStub() *Stub{
	callCount := make([]CallCount,0)
	stub := Stub{callCount}

	return &stub
}

// Stub Object. Defines the CallCount and Register functions.
type Stub struct {
	callCount []CallCount
}

// Fetch all of the function calls and counts
func (stub *Stub) CallCount(name string) CallCount {
	for index := range stub.callCount {
		if stub.callCount[index].funcName == name {
			return stub.callCount[index]
		}
	}

	return CallCount{name, 0}
}

// Register a function call with the stuber
func (stub *Stub) Register(name string) {
	for index := range stub.callCount {
		if stub.callCount[index].funcName == name {
			stub.callCount[index].callCount++

			return
		}
	}

	callCount := CallCount{name,1}
	stub.callCount = append(stub.callCount, callCount)
}
