package http

//
//import (
//	"encoding/json"
//	"net/http"
//	"strconv"
//
//	"github.com/bellasouzas/code-challenge-stone-solution/controllers"
//	"github.com/bellasouzas/code-challenge-stone-solution/model"
//	"github.com/gorilla/mux"
//)
//
//var Accounts []model.AccountID
//
//// '/Accounts'
//func HandleAccounts(res http.ResponseWriter, req *http.Request) {
//	res.Header().Set("content-Type", "application/json")
//	json.NewEncoder(res).Encode(controllers.Accounts())
//}
//
//func HandleGetAccountByID(res http.ResponseWriter, req *http.Request) {
//	res.Header().Set("content-Type", "application/json")
//	params := mux.Vars(req)
//	result := controllers.GetAccountByID(params)
//	json.NewEncoder(res).Encode(result)
//}
//
//func HandleCreateAccount(res http.ResponseWriter, req *http.Request) {
//	res.Header().Set("content-Type", "application/json")
//	var ReqBody model.AccountID
//	_ = json.NewDecoder(req.Body).Decode(&ReqBody)
//	result := controllers.CreateAccount(ReqBody)
//	json.NewEncoder(res).Encode(result)
//}
//
//// '/transfers'
//func HandleGetTransfers(res http.ResponseWriter, req *http.Request) {
//	res.Header().Set("content-Type", "application/json")
//	var account string
//	vars := mux.Vars(req)
//	account = vars["account"]
//	payload := controllers.GetTransfers(account)
//	json.NewEncoder(res).Encode(payload)
//}
//
//func HandleSendTransfer(res http.ResponseWriter, req *http.Request) {
//	res.Header().Set("content-Type", "application/json")
//
//	vars := mux.Vars(req)
//	account := vars["account"]
//	destination := vars["destination"]
//	amount, _ := strconv.ParseFloat(vars["amount"], 64)
//
//	payload, err := controllers.SendTransfer(account, destination, amount)
//	if err != nil {
//		return
//	}
//
//	json.NewEncoder(res).Encode(payload)
//}
