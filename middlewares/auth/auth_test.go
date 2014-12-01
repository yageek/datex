package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestSignature(t *testing.T) {

	user := &ApiUser{PrivateKey: "qwertz", PublicKey: "asdf"}

	body := `{"name": "test", "a" : 1 , "b" : 1}`
	req, _ := http.NewRequest("POST", "http://localhost:3000/ellipsoid/create", strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	user.Sign(req)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Errorf("Error:%s\n", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Create failed :(\n")
	}
}
