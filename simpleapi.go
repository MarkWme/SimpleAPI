package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

type version struct {
	Application string
	Version     string
}

var ready bool = true

func getVersion(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	json.NewEncoder(ctx).Encode(getVersionValue())
	log.Println("getVersion endpoint:", getVersionValue())
}

func podTerminate() {
	log.Println("podTerminate endpoint: Starting 30 second waiting period ...")
	ready = false
	time.Sleep(30 * time.Second)
	log.Println("Waiting period complete")
}

func podReady(ctx *fasthttp.RequestCtx) {
	log.Println("podReady endpoint:", ready)
	if ready {
		ctx.SetStatusCode(fasthttp.StatusOK)
		fmt.Fprintf(ctx, "OK")
	} else {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}
}

func main() {
	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/api/getVersion":
			getVersion(ctx)
		case "/api/podTerminate":
			podTerminate()
		case "/api/podReady":
			podReady(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}
	log.Println("Starting HTTP server on port 3000")
	fasthttp.ListenAndServe(":3000", m)
}

func getVersionValue() version {
	return version{"Simple API Server", "1.0.0"}
}
