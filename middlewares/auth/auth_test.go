package auth

import (
	"net/http"
	"testing"
	"time"
)

func TestSignature(t *testing.T) {

	user := &ApiUser{PrivateKey: "qwertz", PublicKey: "asdf"}

	req, _ := http.NewRequest("POST", "http://localhost:3000/ellipsoid/create", nil)
	req.Header.Add("Content-Type", "application/json")

	v := req.URL.Query()

	v.Add(TimestampKey, time.Now().Local().Format("20060102150405"))
	v.Add(PublicKeyKey, user.PublicKey)
	req.URL.RawQuery = v.Encode()

	v.Add(SignatureKey, string(signature(req, user)))
	req.URL.RawQuery = v.Encode()

	client := &http.Client{}
	client.Do(req)
}
