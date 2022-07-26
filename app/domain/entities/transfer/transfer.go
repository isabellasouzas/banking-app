package transfer

import (
	"time"

	"banking-app/app/domain/types"
)

type Transfer struct {
	OriginID      string
	DestinationID string
	Amount        int
}

type Info struct {
	ID            types.TransferID
	OriginID      string
	DestinationID string
	Amount        int
	CreateAt      time.Time
}
