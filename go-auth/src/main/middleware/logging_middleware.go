package middleware

import (
	"bytes"
	"context"
	"github.com/BevisDev/backend-template/logger"
	"github.com/BevisDev/backend-template/utils"
	"github.com/gin-gonic/gin"
	"goauth/src/main/global"
	"io"
	"net/http"
	"strings"
	"time"
)

type ResponseWrapper struct {
	gin.ResponseWriter
	body     *bytes.Buffer
	status   int
	duration float64
}

func (w *ResponseWrapper) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			state     = c.GetHeader("state")
			startTime = time.Now()
		)
		if utils.IsNilOrEmpty(state) {
			state = utils.GenUUID()
		}
		// write state in header response
		c.Writer.Header().Set("state", state)

		// store state in context
		ctx := context.WithValue(c.Request.Context(), "state", state)
		c.Request = c.Request.WithContext(ctx)

		// ignore log some content-type
		ignoreBody := isIgnoreBody(c.Request.Header)

		// log request
		var reqBody string
		if !ignoreBody {
			reqBytes, _ := io.ReadAll(c.Request.Body)
			reqBody = string(reqBytes)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBytes))
		}

		global.Logger.LogRequest(&logger.RequestLogger{
			State:  state,
			URL:    c.Request.URL.String(),
			Time:   startTime,
			Query:  c.Request.URL.RawQuery,
			Method: c.Request.Method,
			Header: c.Request.Header,
			Body:   reqBody,
		})

		// wrap the responseWriter to capture the response body
		respBuffer := &bytes.Buffer{}
		writer := &ResponseWrapper{
			ResponseWriter: c.Writer,
			body:           respBuffer,
		}
		c.Writer = writer

		// process next
		c.Next()

		// log response
		var (
			respHeaders = c.Writer.Header()
			duration    = time.Since(startTime)
			respBody    string
		)
		ignoreBody = isIgnoreBody(respHeaders)
		if !ignoreBody {
			respBody = writer.body.String()
		}

		global.Logger.LogResponse(&logger.ResponseLogger{
			State:       state,
			Status:      c.Writer.Status(),
			DurationSec: duration,
			Header:      respHeaders,
			Body:        respBody,
		})
	}
}

func isIgnoreBody(headers http.Header) bool {
	contentType := headers.Get("Content-Type")
	return strings.HasPrefix(contentType, "image") ||
		strings.HasPrefix(contentType, "video") ||
		strings.HasPrefix(contentType, "audio")
}
