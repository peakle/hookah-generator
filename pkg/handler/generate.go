package handler

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// Generate - `generate` api route
func Generate(ctx *fasthttp.RequestCtx) {
	_, _ = fmt.Fprintf(ctx, "hookah")
}
