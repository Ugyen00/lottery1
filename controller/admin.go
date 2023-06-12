package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	model "main.go/Model"
	"main.go/utils/httpResp"
)


func RegisterHandler(w http.ResponseWriter,r *http.Request){
	fmt.Println("here")
	var admin model.Admin 
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil{
		fmt.Println("error in decoding the request")
		httpResp.RespondWithError(w,http.StatusBadRequest,"error in decoding the request")
		return
	}
	getErr := admin.CreateAdmin()
	if getErr != nil{
		fmt.Println("error in inserting the data")
		httpResp.RespondWithError(w,http.StatusBadRequest,"error in inserting ")
		return 
	}
	httpResp.RespondWithJSON(w,http.StatusOK,map[string]string{"message":"successful"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request){
	var admin model.Admin 
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil{
		fmt.Println("error in decoding the request")
		return 
	}
	var admin1 model.Admin
	getErr := admin1.Check(admin.Email)
	if getErr != nil{
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w,http.StatusUnauthorized,"invalid login")
			fmt.Println("email doesn't match")
		default:
			httpResp.RespondWithError(w,http.StatusBadRequest,"error in getting the data")
			fmt.Println("error in getting the data")

		}
		return
	}
	if admin.Password != admin1.Password{
		httpResp.RespondWithError(w,http.StatusUnauthorized,"invalid login")
		fmt.Println("invalid login from password")
		return
	}
	//create a cookie
	cookie := http.Cookie{
		Name: "u-cookie",
		// Value: email +admin.Password,
		Value:   "Utech",
		Expires: time.Now().Add(30 * time.Minute),
		Secure:  true,
	}
	//set cookie and send back to client
	http.SetCookie(w, &cookie)
	httpResp.RespondWithJSON(w,http.StatusOK,map[string]string{"message":"successful"})
	fmt.Println("successful")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request){
		http.SetCookie(w, &http.Cookie{
		Name:    "u-cookie",
		Expires: time.Now(),
	})
	fmt.Println("logout successful")
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "logout successful"})
}
func VerifyCookie(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("u-cookie")
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			httpResp.RespondWithError(w, http.StatusSeeOther, "cookie not set")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, "internal server error")
		}
		return false
	}

	if cookie.Value != "Utech" {
		httpResp.RespondWithError(w, http.StatusSeeOther, "invalid cookie")
		return false
	}
	return true
}
