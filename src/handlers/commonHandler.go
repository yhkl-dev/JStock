package handlers

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

type ResultFunc func(message string, code string, result interface{}) func(output Output)
type Output func(c *gin.Context, v interface{})

type JSONResult struct {
	Message string      `json:"message"`
	Code    string      `json:"code"`
	Result  interface{} `json:"result"`
}

func NewJSONResult(message string, code string, result interface{}) *JSONResult {
	return &JSONResult{Message: message, Code: code, Result: result}
}

var ResultPool *sync.Pool

func init() {
	ResultPool = &sync.Pool{
		New: func() interface{} {
			return &JSONResult{Message: "", Code: "", Result: nil}
		},
	}
}

func Return(c *gin.Context) ResultFunc {
	return func(message string, code string, result interface{}) func(output Output) {
		r := ResultPool.Get().(*JSONResult)
		defer ResultPool.Put(r)
		r.Message = message
		r.Code = code
		r.Result = result
		return func(output Output) {
			output(c, r)
		}
	}
}

func OK(c *gin.Context, v interface{}) {
	c.JSON(200, v)
}

func OK4String(c *gin.Context, v interface{}) {
	c.String(200, fmt.Sprintf("%v", v))
}

func Error(c *gin.Context, v interface{}) {
	c.JSON(400, v)
}
