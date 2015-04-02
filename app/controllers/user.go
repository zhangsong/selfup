package controllers

import (
	"github.com/revel/revel"
	"github.com/dchest/captcha"
	"strings"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"github.com/zhangsong/selfup/app/models"
)

type User struct {
	*revel.Controller
}

func (c User) Signup() revel.Result {
	if models.IsLogin(c.Session) {
		return c.Redirect("/")
	}
	captchaId := captcha.New()
	return c.Render(captchaId)
}
func (c User) Signupdo() revel.Result {

	if c.Request.Method !="POST" {
		return c.RenderError(errors.New("no found"))
	}
	u := strings.TrimSpace(c.Params.Get("username"))
	c.Validation.Required(u).Message("请输入用户名")
	c.Validation.MinSize(u, 3).Message("用户名最少为3个")
	c.Validation.MaxSize(u, 15).Message("用户名不能超过15个字符")

	c.Validation.MinSize(strings.TrimSpace(c.Params.Get("password")), 6).Message("密码不能少于6个字符")
	c.Validation.Email(c.Params.Get("email")).Message("请输入正确的邮箱格式")
	if !captcha.VerifyString(c.Params.Get("captchaId"), c.Params.Get("captcha")) {
		c.Flash.Error("验证码错误")
		return c.Redirect("/user/signup")
	}
	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	user := models.User{}
	err := db.C(models.USERS).Find(bson.M{"username":u}).One(&user)
	if err == nil {
		c.Flash.Error("此用户已经注册")
		return c.Redirect("/user/signup")
	}
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/user/signup")
	}
	id := bson.NewObjectId()
	err = db.C(models.USERS).Insert(&models.User{
		Id_:id,
		Username:u,
		Password:c.Params.Get("password"),
		//Salt
		Email:c.Params.Get("email"),
	})

	if err != nil {
		return c.RenderError(errors.New("注册失败"))
	}
	return c.Redirect("/user/signin")
}

func (c User)Signin() revel.Result {
	if models.IsLogin(c.Session) {
		return c.Redirect("/")
	}
	captchaId := captcha.New()
	return c.Render(captchaId)
}
func (c User)Signindo() revel.Result {
	if c.Request.Method !="POST" {
		return c.RenderError(errors.New("错误的登陆请求"))
	}
	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	user := models.User{}
	err := db.C(models.USERS).Find(bson.M{"email":c.Params.Get("email"),"password":c.Params.Get("password")}).One(&user)
	if err != nil {
		c.Flash.Error("用户名或密码错误")
		return c.Redirect("/user/signin")
	}

	if !captcha.VerifyString(c.Params.Get("captchaId"),c.Params.Get("captcha")){//登陆
		c.Flash.Error("验证码错误"+c.Params.Get("captchaId"))
		return c.Redirect("/user/signin")
	}
	models.Login(c.Session, &user)
	return c.Redirect("/user/signin")
}

func (c User)Logout() revel.Result {
	models.Logout(&c.Session)
	//c.Session = make(revel.Session)
	return c.Redirect("/")
}
