package middlewares

import (
	"errors"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/victorsteven/fullstack/api/auth"
	"github.com/victorsteven/fullstack/api/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc, nrapp *newrelic.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		txn := nrapp.StartTransaction(r.URL.Path)
		defer txn.End()	
		// req is a *http.Request, this marks the transaction as a web transaction
		txn.SetWebRequestHTTP(r)
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
