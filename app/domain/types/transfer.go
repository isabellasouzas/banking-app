package types

type TransferID string

func (id TransferID) String() string {
	return string(id)
}
