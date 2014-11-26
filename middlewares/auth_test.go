package middlewares

import (
	"net/http"
	"testing"
	"time"
)

func TestSignature(t *testing.T) {

	user := &ApiUser{PrivateKey: "9d7692a7-245b-404e-9872-373699960f8a", PublicKey: "c2368ab1-4679-4083-9069-8cda10c86e08"}

	req, _ := http.NewRequest("POST", "http://localhost:3000/ellipse/create", nil)
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
