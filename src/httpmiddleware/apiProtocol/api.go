package apiProtocol

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func MiddlewareHandler(c *gin.Context) {

	bodyLogWriter := &bodyLogWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: c.Writer,
	}
	c.Writer = bodyLogWriter

	t := time.Now()
	log.Println(t)
	// ---
	c.Next()
	// ---

	//responseBody := bodyLogWriter.body.String()
	//
	//log.Println(responseBody)
}
