package tests

import (
	"github.com/revel/revel"
	)

type CuserTest struct {
	revel.TestSuite
}

func (t *CuserTest) Before() {
	println("Set up")
}

func (t *CuserTest) TestThatSignupLoadOk() {
	t.Get("/user/signup")
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *CuserTest) After() {
	println("Tear down")
}
