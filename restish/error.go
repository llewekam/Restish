package restish

type DispatchError struct {
}

func (_ *DispatchError) Error() string {
	return "Dispatcher Error"
}

