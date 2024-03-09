package handlers

import (
	"net/http"

	"github.com/alvaroglvn/ravensfield-collection/internal"
	"github.com/alvaroglvn/ravensfield-collection/utils"
)

func CreateMasterKeyWare(c internal.ApiConfig) func(http.Handler) http.Handler {
	function := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Api-Key")
			if header == c.MasterKey {
				next.ServeHTTP(w, r)
				return
			}

			utils.RespondWithError(w, 401, "bad api key")

		})
	}
	return function
}
