package middleware

import (
	"net/http"
	"time"

	"golang.org/x/net/context"
)

func Profile(hdl Handle) Handle {

	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		log := GetRequestLog(ctx)
		t0 := time.Now()

		hdl(ctx, w, r)

		t1 := time.Now()
		log.Info("Profile", "Duration", t1.Sub(t0))
	}
}
