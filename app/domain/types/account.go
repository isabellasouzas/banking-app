package types

type AccountID string

func (a AccountID) String() string {
	return string(a)
}
