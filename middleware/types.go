package middleware

import (
	"net/http"

	"golang.org/x/net/context"
)

type Handle func(ctx context.Context, w http.ResponseWriter, r *http.Request)
