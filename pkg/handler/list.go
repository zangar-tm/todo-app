package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zangar-tm/todo-app"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	//call service method
	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (handler *Handler) getAllLists(c *gin.Context) {

}

func (handler *Handler) getListById(c *gin.Context) {

}

func (handler *Handler) updateList(c *gin.Context) {

}

func (handler *Handler) deleteList(c *gin.Context) {

}
