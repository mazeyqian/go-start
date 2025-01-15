package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/mazeyqian/go-gin-gee/internal/pkg/models/alias2data"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/persistence"
	http_err "github.com/mazeyqian/go-gin-gee/pkg/http-err"
)

func GetDataByAlias(c *gin.Context) {
	rep := persistence.GetAlias2dataRepository()
	alias := c.Query("alias")
	log.Println("GetDataByAlias alias", alias)
	if data, err := rep.Get(alias); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("data not found"))
		log.Println(err)
	} else {
		if !data.Public {
			// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403
			http_err.NewError(c, http.StatusForbidden, errors.New("the client does not have access rights to the content"))
		} else {
			c.JSON(http.StatusOK, gin.H{"data": data})
		}
	}
}

func CreateAlias2data(c *gin.Context) {
	rep := persistence.GetAlias2dataRepository()
	var alias2dataInput models.Alias2data
	_ = c.BindJSON(&alias2dataInput)
	if err := rep.Add(&alias2dataInput); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, alias2dataInput)
	}
}

func CountAlias2data(c *gin.Context) {
	rep := persistence.GetAlias2dataRepository()
	alias := c.Query("alias")
	// log.Println("CountAlias2data alias", alias)
	count, err := rep.CountByAlias(alias)
	if err != nil {
		log.Println(err)
		http_err.NewError(c, http.StatusNotFound, errors.New("unable to retrieve count"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}
