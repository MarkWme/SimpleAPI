package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
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

func get1KBFile(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/octet-stream")
	ctx.Response.Header.Set("Content-Disposition", "attachment; filename=1kb.bin")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprint(ctx, strings.Repeat("X", 1024))
	log.Println("get1KBFile endpoint")
}

func get1MBFile(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/octet-stream")
	ctx.Response.Header.Set("Content-Disposition", "attachment; filename=1mb.bin")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprint(ctx, strings.Repeat("X", 1048576))
	log.Println("get1MBFile endpoint")
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
		case "/api/get1KBFile":
			get1KBFile(ctx)
		case "/api/get1MBFile":
			get1MBFile(ctx)
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
