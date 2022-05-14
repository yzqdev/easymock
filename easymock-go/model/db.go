package model

import (
	"easymock/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Database struct {
	*gorm.DB
}
type BaseModel struct {
	ID        uint      `json:"id"gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var DB *gorm.DB

// Init Open mysql 连接
func init() {
	g := config.GetGlobal()

	//dsn := g.Mysql.User + ":" + g.Mysql.Pass + "@tcp(127.0.0.1:3306)/" + g.Mysql.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "host=localhost user=postgres password=" + g.Pgsql.Pass + " dbname= " + g.Pgsql.Name + " port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Group{}, &Mock{}, &Project{}, &AdminUser{})
	if err != nil {
		return
	}

	DB = db
	return
}

func GetDb() *gorm.DB {
	return DB
}
