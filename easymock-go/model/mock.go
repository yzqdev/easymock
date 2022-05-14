package model

import (
	"github.com/gookit/color"
)

type Mock struct {
	BaseModel
}

func (mock *Mock) QueryProjectList() (result map[string]interface{}) {

	return nil
}
func (mock *Mock) QueryAddProject(article Project) (flag bool) {
	db := GetDb()
	color.Red.Println(article)
	db.Table("article").Create(&article)

	return true
}
func (mock *Mock) QueryUpdateProject(article Project) (flag bool) {
	db := GetDb()

	color.Red.Println(article)
	db.Table("article").Save(&article)

	return true
}
func (mock *Mock) DeleteProject(id int) (flag bool) {
	db := GetDb()
	var article Project
	db.Table("article").Where("id= ?", id).Delete(&article)
	return true
}
func (mock *Mock) QueryGetProjectById(id int) (result Project) {
	db := GetDb()

	db.Table("article").Where("id=?", id).First(&result)
	color.Blueln(result)
	return
}
func (mock *Mock) GetAllPartCount() (flag map[string]interface{}) {
	sql := "SELECT SUM(part_count) as all_part_count ,SUM(view_count) as all_view_count FROM article"
	db := GetDb()
	data := map[string]interface{}{
		"all_part_count": 456, "all_view_count": 123,
	}
	db.Exec(sql).Scan(&data)
	return data
}
