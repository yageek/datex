package middlewares

import (
	"compress/gzip"
	"github.com/gorilla/context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const RequestContentDataKey = "Content-Data"

func CheckPostRequest(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if r.Method != "POST" {
		next(rw, r)
	} else {

		content_type := r.Header.Get("Content-Type")

		if content_type != "application/json" {
			http.Error(rw, "Invalid content type", http.StatusBadRequest)
			return
		}

		encoding := r.Header.Get("Content-Encoding")

		var finalreader io.Reader
		body := r.Body

		if encoding == "gzip" {
			finalreader, _ = gzip.NewReader(body)
		} else {
			finalreader = body
		}

		data, err := ioutil.ReadAll(finalreader)

		if err != nil {
			log.Printf("Decode error:%s\n", err)
			http.Error(rw, "Could not read request body", http.StatusInternalServerError)
			return
		}
		context.Set(r, RequestContentDataKey, data)
		next(rw, r)
	}

}

func GetData(r *http.Request) []byte {
	i := context.Get(r, RequestContentDataKey)

	data, ok := i.([]byte)

	if !ok {
		log.Printf("Type assertion failed \n")
		return nil
	}

	return data
}
