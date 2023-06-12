package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"main.go/utils/httpResp"

	"github.com/gorilla/mux"
	model "main.go/Model"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// if !VerifyCookie(w, r) {
	// 	return
	// }
	p := mux.Vars(r)
	ticket := p["ticket"]
	_, err := w.Write([]byte("hellow world" + ticket))
	if err != nil {
		fmt.Println("error:", err)
	}
}

func BuyTicket(w http.ResponseWriter, r *http.Request) {
	// if !VerifyCookie(w, r) {
	// 	return
	// }
	var tik model.Ticket
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&tik); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")

		return
	}
	defer r.Body.Close()
	saveErr := tik.Create()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "Ticket Booked"})
}

func GetTik(w http.ResponseWriter, r *http.Request) {
	// if !VerifyCookie(w, r) {
	// 	return
	// }
	tid := mux.Vars(r)["tid"]
	tikid, idErr := getUserId(tid)

	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	s := model.Ticket{TikId: tikid}
	getErr := s.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "Ticket Not Found")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, s)
}

func getUserId(userIdParam string) (int64, error) {

	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, userErr
	}
	return userId, nil
}

func DeleteTik(w http.ResponseWriter, r *http.Request) {
	// if !VerifyCookie(w, r) {
	// 	return
	// }
	tid := mux.Vars(r)["tid"]
	tikId, idErr := getUserId(tid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	s := model.Ticket{TikId: tikId}
	if err := s.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func GetAllTiks(w http.ResponseWriter, r *http.Request) {
	// if !VerifyCookie(w, r) {
	// 	return
	// }
	fmt.Println("Here1")
	tickets, getErr := model.GetAllTickets()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		fmt.Println("error",getErr)
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, tickets)
	fmt.Println("success",tickets)
}

func UpdateTik(w http.ResponseWriter, r *http.Request) {
	// if !VerifyCookie(w, r) {
	// 	return
	// }
	fmt.Println("hi")
	old_tik := mux.Vars(r)["tid"]
	fmt.Println(old_tik)
	old_tikid, idErr := getUserId(old_tik)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	var tik model.Ticket
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tik); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	err := tik.Update(old_tikid)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		fmt.Println("error",err)
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, tik)
}
