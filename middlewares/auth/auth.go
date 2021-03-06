/*
The authentication is currently very basic and based on idea proposed
on http://www.thebuzzmedia.com/designing-a-secure-rest-api-without-oauth-authentication/.
*/
package auth

import (
	"code.google.com/p/go-uuid/uuid"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/yageek/datex/middlewares/mongo"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	ApiKeysCollection = "api_keys"
	PublicKeyKey      = "public_key"
	TimestampKey      = "timestamp"
	SignatureKey      = "signature"
	maxSecondDelta    = 10
)

type ApiUser struct {
	Name       string `bson:"name"`
	PublicKey  string `bson:"public_key"`
	PrivateKey string `bson:"private_key"`
}

var chk *accessTokenChecker

func init() {
	chk = newaccessTokenChecker()
}

// Wrap a http.HandlerFunc to secure it
func SecureHandleFunc(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chk.ServeHTTP(w, r, handler)
	}
}

func newApiUser(name string) *ApiUser {

	return &ApiUser{Name: name, PublicKey: uuid.New(), PrivateKey: uuid.New()}
}

func newaccessTokenChecker() *accessTokenChecker {
	return &accessTokenChecker{}
}

type accessTokenChecker struct {
}

func (a *accessTokenChecker) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.Handler) {

	r := &http.Request{}
	*r = *req

	v := r.URL.Query()

	clientPublicKey := v.Get(PublicKeyKey)
	timestamp := v.Get(TimestampKey)
	sentSignature := v.Get(SignatureKey)
	user := userFromPublicKey(req, clientPublicKey)

	now := time.Now().UTC()
	reqTime, err := time.Parse("20060102150405", timestamp)

	if err != nil {
		http.Error(w, "Invalid timestamp", http.StatusBadRequest)
		return
	}

	if now.Sub(reqTime).Seconds() > maxSecondDelta {
		http.Error(w, "Too old request", http.StatusUnauthorized)
		return
	}
	if clientPublicKey == "" || timestamp == "" || sentSignature == "" || user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	v.Del(SignatureKey)
	r.URL.RawQuery = v.Encode()

	expected_signature := signature(r, user)
	if ok := hmac.Equal(expected_signature, []byte(sentSignature)); ok {
		//Check signature
		next.ServeHTTP(w, req)
	} else {
		http.Error(w, "Unauthorized - Wrong credentials", http.StatusUnauthorized)
	}

}

func signature(r *http.Request, user *ApiUser) []byte {

	args := []string{r.Method, r.Host, r.URL.Path}

	for key, arg := range r.URL.Query() {
		args = append(args, fmt.Sprintf("%s=%s", key, arg[0]))
	}

	sort.Strings(args)

	data := []byte(strings.Join(args, "_"))
	mac := hmac.New(sha256.New, []byte(user.PrivateKey))
	mac.Write(data)
	return mac.Sum(nil)
}

func userFromPublicKey(r *http.Request, publicKey string) *ApiUser {

	if publicKey == "" {
		return nil
	}

	c := mongo.Collection(r, ApiKeysCollection)
	u := ApiUser{}

	err := c.Find(bson.M{"public_key": publicKey}).One(&u)

	if err != nil {
		log.Println("Error during DB query:", err)
		return nil
	}

	return &u
}

func (user *ApiUser) Sign(req *http.Request) {
	v := req.URL.Query()

	v.Add(TimestampKey, time.Now().Local().Format("20060102150405"))
	v.Add(PublicKeyKey, user.PublicKey)
	req.URL.RawQuery = v.Encode()

	v.Add(SignatureKey, string(signature(req, user)))
	req.URL.RawQuery = v.Encode()
}
