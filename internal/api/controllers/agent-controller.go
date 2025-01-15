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
