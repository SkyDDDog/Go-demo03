package controller

import (
	"demo03/com/west2/common"
	"demo03/com/west2/entity"
	"demo03/com/west2/service"
	"demo03/com/west2/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
	"time"
)

type NoteController struct {
}

// 注册接口
func NoteControllerRegister(noteGrp *gin.RouterGroup) {
	noteController := &NoteController{}
	noteGrp.Use().POST("", noteController.createNote)
	noteGrp.Use().PUT("/done/:id", noteController.DoneNote)
	noteGrp.Use().PUT("/undo/:id", noteController.UndoNote)
	noteGrp.Use().PUT("/done", noteController.DoneAllNote)
	noteGrp.Use().PUT("/undo", noteController.UndoAllNote)
	noteGrp.Use().DELETE("/:id", noteController.DeleteNote)
	noteGrp.Use().DELETE("/done", noteController.DeleteDoneNote)
	noteGrp.Use().DELETE("/undo", noteController.DeleteUndoNote)
	noteGrp.Use().DELETE("/all", noteController.DeleteAllNote)
	noteGrp.Use().GET("", noteController.GetAllNote)
	noteGrp.Use().GET("/done", noteController.GetDoneNote)
	noteGrp.Use().GET("/undo", noteController.GetUndoNote)
	noteGrp.Use().GET("/:content", noteController.SearchNote)

}

// @Summary		新增待办
// @Description	添加一条待办事项
// @Tags			NoteController
// @Accept			json
// @Produce		json
// @param			string	body		string	false  "note"
// @Success		200		{object}	string
// @Router			/note [post]
func (controller NoteController) createNote(c *gin.Context) {
	// 定义map或结构体
	note := entity.NewNote()
	var m map[string]interface{}
	// 反序列化
	c.ShouldBindBodyWith(&note, binding.JSON)
	c.ShouldBindBodyWith(&m, binding.JSON)

	note.State = "0"
	t := util.Strval(m["ddl"])
	timeLayoutStr := "2006-01-02 15:04:05"
	note.EndTime, _ = time.Parse(timeLayoutStr, t)
	service.Service.CreateNote(note)
	//fmt.Println(user)
	common.CommonSuccess(c)
}

// @Summary		完成待办
// @Description	将 一条 代办事项设置为已完成
// @Tags			NoteController
// @param			string	path		string	false  "note"
// @Success		200		{object}	string
// @Router			/note/done/{id} [PUT]
func (controller NoteController) DoneNote(c *gin.Context) {
	id := c.Param("id")
	service.Service.UpdateNoteState(id, "1")
	common.CommonSuccess(c)
}

// @Summary		设为待办
// @Description	将 一条 已完成事项设置为待办
// @Tags			NoteController
// @param			string	path		string	false  "note"
// @Success		200		{object}	string
// @Router			/note/undo/{id} [PUT]
func (controller NoteController) UndoNote(c *gin.Context) {
	id := c.Param("id")
	service.Service.UpdateNoteState(id, "0")
	common.CommonSuccess(c)
}

// @Summary		完成待办
// @Description	将 多条 代办事项设置为已完成
// @Tags			NoteController
// @Success		200		{object}	string
// @Router			/note/done [PUT]
func (controller NoteController) DoneAllNote(c *gin.Context) {
	service.Service.UpdateAllNoteState("1")
	common.CommonSuccess(c)
}

// @Summary		设为待办
// @Description	将 多条 已完成事项设置为待办
// @Tags			NoteController
// @Success		200		{object}	string
// @Router			/note/undo [PUT]
func (controller NoteController) UndoAllNote(c *gin.Context) {
	service.Service.UpdateAllNoteState("0")
	common.CommonSuccess(c)
}

// @Summary		删除事项
// @Description	删除 一条 事项
// @Tags			NoteController
// @param			string	path		string	false  "id"
// @Success		200		{object}	string
// @Router			/note/{id} [DELETE]
func (controller NoteController) DeleteNote(c *gin.Context) {
	id := c.Param("id")
	service.Service.DeleteNoteById(id)
	common.CommonSuccess(c)
}

// @Summary		删除事项
// @Description	删除 所有已经完成 事项
// @Tags			NoteController
// @Success		200		{object}	string
// @Router			/note/done [DELETE]
func (controller NoteController) DeleteDoneNote(c *gin.Context) {
	service.Service.DeleteNoteByState("1")
	common.CommonSuccess(c)
}

// @Summary		删除事项
// @Description	删除 所有待办 事项
// @Tags			NoteController
// @Success		200		{object}	string
// @Router			/note/undo [DELETE]
func (controller NoteController) DeleteUndoNote(c *gin.Context) {
	service.Service.DeleteNoteByState("0")
	common.CommonSuccess(c)
}

// @Summary		删除事项
// @Description	删除 所有 事项
// @Tags			NoteController
// @Success		200		{object}	string
// @Router			/note/all [DELETE]
func (controller NoteController) DeleteAllNote(c *gin.Context) {
	service.Service.DeleteAllNote()
	common.CommonSuccess(c)
}

// @Summary		查看事项
// @Description	查看所有事项
// @Tags			NoteController
// @Param 	query string true "pageNum" 第几页
// @Param 	query string true "pageSize" 一页几条数据
// @Success		200		{object}	string
// @Router			/note [Get]
func (controller NoteController) GetAllNote(c *gin.Context) {
	pageNum := c.DefaultQuery("pageNum", "1")
	pageSize := c.DefaultQuery("pageSize", "5")
	num, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	result := service.Service.GetAllNote(num, size)
	common.Success(result, c)
}

// @Summary		查看事项
// @Description	查看所有已完成事项
// @Tags			NoteController
// @Param 	query string true "pageNum" 第几页
// @Param 	query string true "pageSize" 一页几条数据
// @Success		200		{object}	string
// @Router			/note/done [Get]
func (controller NoteController) GetDoneNote(c *gin.Context) {
	pageNum := c.DefaultQuery("pageNum", "1")
	pageSize := c.DefaultQuery("pageSize", "5")
	num, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	result := service.Service.GetNoteByState("1", num, size)
	common.Success(result, c)
}

// @Summary		查看事项
// @Description	查看所有待办事项
// @Tags			NoteController
// @Param 	query string true "pageNum" 第几页
// @Param 	query string true "pageSize" 一页几条数据
// @Success		200		{object}	string
// @Router			/note/done [Get]
func (controller NoteController) GetUndoNote(c *gin.Context) {
	pageNum := c.DefaultQuery("pageNum", "1")
	pageSize := c.DefaultQuery("pageSize", "5")
	num, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	result := service.Service.GetNoteByState("0", num, size)
	common.Success(result, c)
}

// @Summary		查询事项
// @Description	输入关键词查询事项
// @Tags			NoteController
// @Param   path  string true  "content" 搜索关键词
// @Param 	query string true "pageNum" 第几页
// @Param 	query string true "pageSize" 一页几条数据
// @Success		200		{object}	string
// @Router			/note/done/{content} [Get]
func (controller NoteController) SearchNote(c *gin.Context) {
	content := c.Param("content")
	pageNum := c.DefaultQuery("pageNum", "1")
	pageSize := c.DefaultQuery("pageSize", "5")
	num, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	result := service.Service.SearchNote(content, num, size)
	common.Success(result, c)
}
