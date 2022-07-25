package transfer

import (
	"time"

	"github/isabellasouzas/banking-go/app/domain/types"
)

type Transfer struct {
	OriginID      int
	DestinationID int
	Amount        int
}

type Info struct {
	ID            types.TransferID
	OriginID      int
	DestinationID int
	Amount        int
	CreateAt      time.Time
}
