package server

import (
	"github.com/qb0C80aE/clay/middleware"
	"github.com/qb0C80aE/clay/router"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/submodules"
)

func Setup(db *gorm.DB) *gin.Engine {
	submodules.HookSubmodules()
	r := gin.Default()
	r.Use(middleware.SetDBtoContext(db))
	router.Initialize(r)
	return r
}
