package mid

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/foundation/web"
	"go.uber.org/zap"
)

func Logger(log *zap.SugaredLogger) web.Middleware {
	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			log.Infow("request started", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

			err := handler(ctx, w, r)

			log.Infow("request completed", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

			return err
		}

		return h
	}

	return m
}
