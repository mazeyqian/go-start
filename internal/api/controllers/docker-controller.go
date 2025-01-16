package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/persistence"
	http_err "github.com/mazeyqian/go-gin-gee/pkg/http-err"
)

func GetTag(c *gin.Context) {
	rep := persistence.GetDockerRepository()
	tagName, err := rep.GetTagName("mazeyqian", "go-gin-gee", "api")
	if err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"tagName": tagName})
	}
}
