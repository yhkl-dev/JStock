package jwtauth

import (
	"JStock/src/interfaces/utils"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type TokenCheck struct {
}

func NewTokenCheck() *TokenCheck {
	return &TokenCheck{}
}

func (s *TokenCheck) OnRequest(ctx *gin.Context) error {
	if ctx.Request.URL.Path == "/v1/user/login" {
		return nil
	}
	token := ctx.Request.Header.Get("authoritaion")
	if token == "" {
		goft.Throw("Unauthorized", 503, ctx)
	}
	j := utils.NewJWT()
	_, err := j.ResolveToken(token)
	if err != nil {
		goft.Throw("authorized error", 503, ctx)
	}
	return nil
}

func (s *TokenCheck) OnResponse(result interface{}) (interface{}, error) {
	return result, nil
}
