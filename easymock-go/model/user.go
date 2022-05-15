package model

import (
	"easymock/utils"
	"fmt"
	"github.com/gookit/color"
	"github.com/rs/xid"
)

type AdminUser struct {
	BaseModel
	Username string `json:"username"`
	Uid      string `json:"uid"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

func QueryByUsername(username string) (result AdminUser) {

	db := GetDb()
	db.Where("username = ?", username).First(&result)
	color.Red.Println(result)
	return

}
func SaveUser(data *AdminUser) {
	db := GetDb()
	node, err := utils.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	guid := xid.New()
	// Generate a snowflake ID.
	id := node.Generate()
	data.ID = id
	data.Uid = guid.String()
	db.Create(data)

}
func GetUserCheck(username string) bool {
	db := GetDb()
	obj := db.Find(&AdminUser{}).
		Where("username = ?  ", username).
		Where("status in (?)", []string{"1", "2"})
	var count int64
	obj.Count(&count)

	if count > 0 {
		return true
	} else {
		return false
	}
}
func QueryCheckToken() {

}
