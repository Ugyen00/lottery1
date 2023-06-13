package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTicket(t *testing.T) {
	url := "http://localhost:3005/ticket"

	var jsonStr = []byte(`{"tikid":106, "fname": "Karma", "lname": "Zangmo", "phone":"77341164"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	expResp := `{"status":"Ticket added"}`
	assert.JSONEq(t, expResp, string(body))
}

func TestGetTicket(t *testing.T) {
	c := http.Client{}
	r, _ := c.Get("http://localhost:3005/ticket/103")
	body, _ := io.ReadAll(r.Body)
	assert.Equal(t, http.StatusOK, r.StatusCode)
	expResp := `{"tikid":345, "fname":"Tshering", "lname":"Yangchen", "phone":1764567}`
	assert.JSONEq(t, expResp, string(body))
}

func TestDeleteTicket(t *testing.T) {
	url := "http://localhost:3005/ticket/107"

	req, _ := http.NewRequest("DELETE", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	expResp := `{"status":"deleted"}`
	assert.JSONEq(t, expResp, string(body))

}
func TestTicketNotFound(t *testing.T) {
	assert := assert.New(t)
	c := http.Client{}
	r, _ := c.Get("http://localhost:3005/ticket/231")
	body, _ := io.ReadAll(r.Body)
	assert.Equal(http.StatusNotFound, r.StatusCode)
	expResp := `{"error":"Ticket not found"}`
	assert.JSONEq(expResp,string(body))
}