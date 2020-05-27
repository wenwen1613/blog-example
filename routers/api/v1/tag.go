package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/wenwen1613/blog-example/models"
	"github.com/wenwen1613/blog-example/pkg/e"
	"github.com/wenwen1613/blog-example/pkg/setting"
	"github.com/wenwen1613/blog-example/pkg/util"
)

const (
	resCode = "code"
	resMsg  = "msg"
	resData = "data"
)

//GetTags 查询标签
func GetTags(ctx *gin.Context) {
	name := ctx.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(ctx), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	ctx.JSON(http.StatusOK, gin.H{
		resCode: code,
		resMsg:  e.GetMsg(code),
		resData: data,
	})
}

//AddTag 添加标签
func AddTag(ctx *gin.Context) {
	name := ctx.Query("name")
	state := com.StrTo(ctx.Query("state")).MustInt()
	createdBy := ctx.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最大长度为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最大长度为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只能是0或1")

	code := e.SUCCESS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		resCode: code,
		resMsg:  e.GetMsg(code),
		resData: make(map[string]string),
	})

}

//DeleteTags 删除标签
func DeleteTags(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagById(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		resCode: code,
		resMsg:  e.GetMsg(code),
		resData: make(map[string]string),
	})
}

//EditTags 编辑标签
func EditTags(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()
	name := ctx.Query("name")
	modifiedBy := ctx.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最大长度为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagById(id) {
			data := make(map[string]interface{})
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		resCode: code,
		resMsg:  e.GetMsg(code),
		resData: make(map[string]string),
	})
}
