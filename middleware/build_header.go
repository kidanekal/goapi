package middleware

import (
	"net/http"

	"github.com/kidanekal/goapi/constants"
	"golang.org/x/net/context"
)

func VersionHeader(hdl Handle) Handle {

	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {

		w.Header().Set("goapi-Version", constants.Version)

		hdl(ctx, w, r)
	}
}
