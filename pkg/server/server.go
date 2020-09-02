package server

import (
	"strings"

	"github.com/urfave/cli"
	"github.com/valyala/fasthttp"
)

// StartServer - starts server
func StartServer(_ *cli.Context) error {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		path := strings.ToLower(string(ctx.Path()))
		switch path {
		case "/v1/generate":
		default:
			ctx.SetConnectionClose()
		}
	}

	err := fasthttp.ListenAndServe(":80", requestHandler)
	if err != nil {
		return err
	}

	return nil
}
