package model

import (
	"github.com/gookit/color"
)

type Group struct {
	Id int `json:"id"`
}

func (group *Group) QueryProjectList() (result map[string]interface{}) {

	return nil
}
func (group *Group) QueryAddProject(article Project) (flag bool) {
	db := GetDb()
	color.Red.Println(article)
	db.Table("article").Create(&article)

	return true
}
func (group *Group) QueryUpdateProject(article Project) (flag bool) {
	db := GetDb()

	color.Red.Println(article)
	db.Table("article").Save(&article)

	return true
}
func (group *Group) DeleteProject(id int) (flag bool) {
	db := GetDb()
	var article Project
	db.Table("article").Where("id= ?", id).Delete(&article)
	return true
}
func (group *Group) QueryGetProjectById(id int) (result Project) {
	db := GetDb()

	db.Table("article").Where("id=?", id).First(&result)
	color.Blueln(result)
	return
}
func (group *Group) GetAllPartCount() (flag map[string]interface{}) {
	sql := "SELECT SUM(part_count) as all_part_count ,SUM(view_count) as all_view_count FROM article"
	db := GetDb()
	data := map[string]interface{}{
		"all_part_count": 456, "all_view_count": 123,
	}
	db.Exec(sql).Scan(&data)
	return data
}
