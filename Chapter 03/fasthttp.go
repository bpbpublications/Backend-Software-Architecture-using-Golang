package main

import (
    "github.com/valyala/fasthttp"
)

func requestHandler(ctx *fasthttp.RequestCtx) {
    ctx.WriteString("Hello, World!")
}

func main() {
    fasthttp.ListenAndServe(":8080", requestHandler)
}
