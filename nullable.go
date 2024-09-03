package compac

// Nl fast and easy to use sql.Null-like abstraction, which allocates memory on stack.
// It's alternative to pointer fields in structs or function parameters.
type Nl[T any] struct {
	Data  T
	Valid bool
}

// NlFromPtr converts a pointer to a Nl[T] type.
func NlFromPtr[T any](ptr *T) Nl[T] {
	if ptr == nil {
		return Nl[T]{}
	}
	return Nl[T]{Data: *ptr, Valid: true}
}

func NlFromValue[T any](value T) Nl[T] {
	return Nl[T]{
		Data:  value,
		Valid: true,
	}
}
