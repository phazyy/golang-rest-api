package middleware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/inconshreveable/log15.v2"
)

var (
	green  = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white  = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red    = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue   = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	cyan   = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset  = string([]byte{27, 91, 48, 109})
)

// Logger : Custom middleware for unifing application and gin logs
func Logger(log log15.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", log)
		c.Next()

		statusColor := colorForStatus(c.Writer.Status())
		methodColor := colorForMethod(c.Request.Method)

		output := []string{
			statusColor, strconv.Itoa(c.Writer.Status()), reset,
			"|", methodColor, reset, c.Request.Method,
			"\t", c.Request.URL.Path,
		}

		log.Info("gin: " + strings.Join(output, " "))
	}
}

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return green
	case code >= 300 && code < 400:
		return white
	case code >= 400 && code < 500:
		return yellow
	default:
		return red
	}
}

func colorForMethod(method string) string {
	switch method {
	case "GET":
		return blue
	case "POST":
		return cyan
	case "PUT":
		return yellow
	case "DELETE":
		return red
	default:
		return reset
	}
}
