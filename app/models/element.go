package models

//题元

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"github.com/revel/revel"
	"regexp"
	"encoding/json"
	"fmt"
)

type Element struct {
	Id_             bson.ObjectId `bson:"_id"`
	Yiyong	int
	Content string
	Hash	string
	Aanda	string
}

var R string
var ReAsk *regexp.Regexp

func init() {
	R = `[^\\]({{(.*?)}})|({{(.*?)}})`
	ReAsk = regexp.MustCompile(R)
}

func GetAanda(s string)[]string {
	ss := ReAsk.FindAllStringSubmatch(s, -1)
	var r []string
	r = make([]string,0)
	for _, v := range ss {
		r = append(r, v[2])
	}
	return r
}

func GenHtml(content string, mark string) string {
	//r := ReAsk.ReplaceAllLiteralString(content, `<input type="text" />`)
	r_ := regexp.MustCompile(`({{(.*?)}})`)
	r := ReAsk.ReplaceAllStringFunc(content, func(s string) string {
		return r_.ReplaceAllLiteralString(s, `<input type="text" name="vv` + mark + `[]" />`)
	})
	return r
}

func Juge(db *mgo.Database, eid string, daan []string)([]string, bool) {
	var daann []string
	var els Element
	db.C(EL).Find(bson.M{"_id":bson.ObjectIdHex(eid)}).One(&els)
	if err := json.Unmarshal([]byte(els.Aanda), &daann); err != nil {
		fmt.Println(err, els.Aanda)
		return nil, false
	}
	var r = true
	for i, v := range daann {
		if v != daan[i] {
			r = false
			break
		}
	}
	return nil, r

}

