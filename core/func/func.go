package _func

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func DoFunc(c *Context) (*resty.Response, error) {
	cli := resty.New()
	req := cli.R()
	req.Method = c.Method()
	req.URL = c.URL()
	req.SetHeaders(map[string]string{
		"x-fc-service-name": c.f.Service,
	})
	return req.Send()
}

type Func struct {
}

func New() *Func {
	return &Func{}
}

func (s *Func) HTTPTrigger(c *gin.Context) {
	serviceName := c.Param("service_name")
	functionName := c.Param("function_name")
	// 检查函数是否存在

	resp, err := DoFunc(NewContext(serviceName, functionName).WithReq(c.Request))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Data(resp.StatusCode(), resp.Header().Get("Content-Type"), resp.Body())
}
