package middleware

import (
	"net/http"
)

func S400(w http.ResponseWriter) {
	http.Error(w, http.StatusText(400), 400)
}

func S401(w http.ResponseWriter) {
	http.Error(w, http.StatusText(401), 401)
}

func S404(w http.ResponseWriter) {
	http.Error(w, http.StatusText(404), 404)
}

func S408(w http.ResponseWriter) {
	http.Error(w, http.StatusText(408), 408)
}

func S500(w http.ResponseWriter) {
	http.Error(w, http.StatusText(500), 500)
}
