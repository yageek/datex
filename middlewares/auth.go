package middlewares

import (
	"code.google.com/p/go-uuid/uuid"
	//"github.com/smartystreets/go-aws-auth"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"net/url"
)

const (
	ApiKeysCollection = "api_keys"
	PublicKeyKey      = "public_key"
)

type ApiUser struct {
	Name       string `bson:"name"`
	PublicKey  string `bson:"public_key"`
	PrivateKey string `bson:"private_key"`
}

var chk *AccessTokenChecker

func init() {
	chk = NewAccessTokenChecker()
}

func SecureHandleFunc(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chk.ServeHTTP(w, r, handler)
	}
}

func NewApiUser(name string) *ApiUser {

	return &ApiUser{Name: name, PublicKey: uuid.New(), PrivateKey: uuid.New()}
}

func NewAccessTokenChecker() *AccessTokenChecker {
	return &AccessTokenChecker{}
}

type AccessTokenChecker struct {
}

func (a *AccessTokenChecker) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.Handler) {
	var v url.Values

	if r.Method == "GET" {
		v = r.URL.Query()
	} else {
		r.ParseForm()
		v = r.PostForm
	}

	clientPublicKey := v.Get(PublicKeyKey)

	if clientPublicKey == "" {
		http.Error(w, "Unauthorize - No public key provided", http.StatusUnauthorized)
		log.Println("No public key provided")
		return
	}

	user := userFromPublicKey(r, clientPublicKey)
	log.Println("Public key:", user.PublicKey)

	//Check signature
	next.ServeHTTP(w, r)
}
func collection(r *http.Request) *mgo.Collection {
	return GetDb(r).C(ApiKeysCollection)
}
func userFromPublicKey(r *http.Request, publicKey string) *ApiUser {
	c := collection(r)
	u := ApiUser{}

	err := c.Find(bson.M{"public_key": publicKey}).One(&u)

	if err != nil {
		log.Println("Error during DB query:", err)
		return nil
	}

	return &u
}
