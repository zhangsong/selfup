package controllers

import (
	"github.com/revel/revel"
	"github.com/dchest/captcha"
	//"net/http"
)

type Captcha struct {
	*revel.Controller
}

//gopher /signup
func (c Captcha) Index() revel.Result {
	captcha.Server(captcha.StdWidth, captcha.StdHeight).ServeHTTP(c.Response.Out,c.Request.Request)
	return nil
}
