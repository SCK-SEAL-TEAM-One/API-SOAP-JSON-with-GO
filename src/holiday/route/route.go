package route

import (
	apiLibrary "holiday/api"

	"github.com/gin-gonic/gin"
)

func NewRoute(api apiLibrary.Api) *gin.Engine {
	route := gin.Default()
	route.POST("/v1/holiday", api.HolidayHandler)
	return route
}
