package controller

import (
	"github.com/gin-gonic/gin"
	"gobaucua/src/main/service"
	"net/http"
)

type SupportController struct {
	service service.ISupportService
}

func NewSupportController(service service.ISupportService) *SupportController {
	return &SupportController{
		service: service,
	}
}

// Ping godoc
// @Summary Ping System API
// @Description Check health system
// @Tags Support
// @Produce  json
// @Success 200 {object} map[string]interface{} "Successful"
// @Router /ping [get]
func (p *SupportController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// PingDB godoc
// @Summary Ping DB API
// @Description Check health DB System
// @Tags Support
// @Produce  json
// @Success 200 {object} map[string]interface{} "Successful"
// @Router /db [get]
func (p *SupportController) PingDB(c *gin.Context) {
	c.JSON(http.StatusOK, p.service.PingDB(c))
}

// PingRedis godoc
// @Summary Ping Redis API
// @Description Check health Redis System
// @Tags Supports
// @Produce  json
// @Success 200 {object} map[string]interface{} "Successful"
// @Router /redis [get]
func (p *SupportController) PingRedis(c *gin.Context) {
	c.JSON(http.StatusOK, p.service.PingRedis(c))
}

//
//// MigrationUp godoc
//// @Summary Migration API
//// @Description Up migration
//// @Tags Support
//// @Produce  json
//// @Param version query int64 false "Migration Version (default: 0)"
//// @Success 200 {object} response.Data "Successful"
//// @Failure 400 {object} response.DataError "Client Error"
//// @Router /api/migration/up [get]
//// @Security BearerAuth
//func (p *SupportController) MigrationUp(c *gin.Context) {
//	version := c.DefaultQuery("version", "0")
//	verInt64, _ := strconv.ParseInt(version, 10, 64)
//	m := &migration.Migration{
//		Dir:        global.AppConfig.Migration.DirPath,
//		TimeoutSec: global.AppConfig.Migration.TimeoutSec,
//		Sqlx:       database.GetDB(consts.DecisionEngineSchema),
//		TypeSQL:    helper.SQLServer,
//		Ctx:        c,
//		Version:    verInt64,
//	}
//	if err := m.Up(); err != nil {
//		response.SetError(c, http.StatusBadRequest, consts.InvalidRequest, err.Error())
//		return
//	}
//	response.OK(c, nil, consts.OK)
//}
//
//// MigrationDown godoc
//// @Summary Migration API
//// @Description Down migration
//// @Tags Support
//// @Produce  json
//// @Param version query int64 false "Migration Version (default: 0)"
//// @Success 200 {object} response.Data "Successful"
//// @Failure 400 {object} response.DataError "Client Error"
//// @Router /migration/down [get]
//// @Security BearerAuth
//func (p *SupportController) MigrationDown(c *gin.Context) {}
