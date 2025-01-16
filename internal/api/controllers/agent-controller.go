package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/mazeyqian/go-gin-gee/internal/pkg/models/agent"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/persistence"
	http_err "github.com/mazeyqian/go-gin-gee/pkg/http-err"
)

func AgentMock(c *gin.Context) {
	rep := persistence.GetAgentRepository()
	var res models.Response
	err := c.BindJSON(&res)
	if err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		return
	}
	data, _ := rep.Mock(&res)
	c.JSON(res.StatusCode, data)
}

func AgentRecord(c *gin.Context) {
	rep := persistence.GetAgentRepository()
	method := c.Query("method")
	url := c.Query("url")
	data1 := c.Query("data")
	var req models.RecordRequestOrResponse
	req.MethodOrStatusCode = method
	req.URL = url
	req.Data = data1
	fileName, err := rep.Record(&req)
	if err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		return
	}
	data := models.ResponseData{
		Code:    200,
		Message: "success",
		Data:    fileName,
	}
	c.JSON(200, data)
}
