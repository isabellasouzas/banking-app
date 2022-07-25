package types

type DomainError struct {
	Operation string
	Err       error
}

func (d DomainError) Error(op string, err error) DomainError {
	return DomainError{
		Operation: op,
		Err:       err,
	}
}
