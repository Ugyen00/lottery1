package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdmLogin(t *testing.T) {

	url := "http://localhost:3003/login"

	//data of type byte sclice
	var jsonStr = []byte(`{"email" :"tshering@gmail.com", "password" : "zangmo"}`)

	//create http request
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	//set request header
	req.Header.Set("content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	//handle error if any
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	expResp := `{"message":"Login success"}`
	// validate if response body is same as expected response body
	assert.JSONEq(t, expResp, string(body))

}

// test for user which does not exist
func TestAdmUserNotExist(t *testing.T) {
	url := "http://localhost:3003/login"
	var data = []byte(`{"email": "tsz@msitprogram.net", "password":pass12"}`)
	// create req object
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	// create client
	client := &http.Client{}
	// send POST request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	assert.JSONEq(t, `{"error":"sql: no rows in result set"}`, string(body))
}