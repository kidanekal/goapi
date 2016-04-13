package api

import (
	"net/http"
)

func PanicHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("panicking"))
	w.WriteHeader(500)

	go func() {

		panic("intentional panic called")
	}()
}
