package transfer

import (
	"time"

	"banking-app/app/domain/usecases/transfer"
)

type Handler struct {
	Usecase transfer.Usecase
}

func NewHandler(transferUsecase transfer.Usecase) Handler {
	return Handler{
		Usecase: transferUsecase,
	}
}

type transferReqBody struct {
	OriginID      int  `json:"origin_id"`
	DestinationID int  `json:"destination_id"`
	Amount        int  `json:"amount"`
	Info          info `json:"info"`
}

type info struct {
	ID            string    `json:"id"`
	OriginID      int       `json:"origin_id"`
	DestinationID int       `json:"destination_id"`
	Amount        int       `json:"amount"`
	CreateAt      time.Time `json:"create_at"`
}
