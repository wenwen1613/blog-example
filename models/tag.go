package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	id         = "id"
	name       = "name"
	createdBy  = "created_by"
	state      = "state"
	modifiedBy = "modified_by"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select(id).Where("name=?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagById(id int) bool {
	var tag Tag
	db.Select(id).Where("id=?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		CreatedBy: createdBy,
		State:     state,
	})
	return true
}

func (tag *Tag) BeforeCreate(scope gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Tag) BeforeUpdate(scope gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}