package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	id         = "id"
	name       = "name"
	createdBy  = "created_by"
	state      = "state"
	modifiedBy = "modified_by"
)

//Tag 结构体
type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//GetTags by params.
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

//GetTagTotal by params.
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

//ExistTagByName by name.
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select(id).Where("name=?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

//ExistTagById 根据id判断是否存在.
func ExistTagById(id int) bool {
	var tag Tag
	db.Select(id).Where("id=?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

//AddTag by params.
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		CreatedBy: createdBy,
		State:     state,
	})
	return true
}

//BeforeCreate 在创建之前执行.
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

//BeforeUpdate 在更新之前操作.
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

//DeleteTag 删除标签.
func DeleteTag(id int) bool {
	db.Where("id=?", id).Delete(&Tag{})
	return true
}

//EditTag 编辑标签,根据ID更新.
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id=?", id).Updates(data)
	return true
}
