package transfer

import (
	"net/http"
	"time"

	"banking-app/app/domain/usecases/transfer"
	"banking-app/app/gateway/http/rest"
)

type Handler struct {
	Usecase transfer.Usecase
}

func NewHandler(transferUsecase transfer.Usecase) Handler {
	return Handler{
		Usecase: transferUsecase,
	}
}

type TransferHandlers interface {
	List(r *http.Request) rest.Response
	SendMoney(r *http.Request) rest.Response
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
