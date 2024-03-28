package main

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		//seting aou own context on top of the cancel one
		timeoutContext, cancel := context.WithTimeout(ctx.Request.Context(), time.Second*2)
		defer cancel()
		//if we cancel the req then the ctx will cancel req
		// req, err := http.NewRequestWithContext(ctx.Request.Context(), http.MethodGet, "http://yahoo.com", nil)
		// on top of the defout context now he have our own wraper
		req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://yahoo.com", nil)
		if err != nil {
			panic(err)
		}

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			panic(err)
		}

		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		ctx.Data(200, "text/html", data)

	})
	r.Run("localhost:3000")

}
