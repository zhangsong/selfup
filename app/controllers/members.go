package controllers

import (
	"github.com/revel/revel"
	//"github.com/dchest/captcha"
	//"strings"
	//"errors"
	"gopkg.in/mgo.v2/bson"
	"github.com/zhangsong/selfup/app/models"
)

type Members struct {
	*revel.Controller
}

//最新会员
func (c Members) Index() revel.Result {
	var members []models.User
	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	db.C(models.USERS).Find(nil).Sort("-username").Limit(40).All(&members)
	membersCount, _ := db.C(models.USERS).Find(nil).Count()
	return c.Render(members, membersCount)
}

//临时
func (c Members)Delete(username string) revel.Result {
	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	err := db.C(models.USERS).Remove(bson.M{"username":username})
	if err != nil {
		return c.RenderError(err)
	}
	return c.Redirect("/members/index")

}
