// Package mid yea yea yea.
package mid

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/foundation/web"
	"go.uber.org/zap"
)

func Logger(log *zap.SugaredLogger) web.Middleware {

	// Logger middleware function.
	m := func(h web.Handler) web.Handler {

		f := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			log.Infow("request started", "method", r.Method, "path", r.URL.Path,
				"remoteaddr", r.RemoteAddr)

			err := h(ctx, w, r)

			log.Infow("request completed", "method", r.Method, "path", r.URL.Path,
				"remoteaddr", r.RemoteAddr, "status", "statuscode")

			return err
		}

		return f
	}

	return m
}
