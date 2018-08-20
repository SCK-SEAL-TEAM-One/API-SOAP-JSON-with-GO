package api

import (
	"holiday/model"
	"holiday/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Api struct {
	HolidayService service.IHolidayService
}

func (api Api) HolidayHandler(c *gin.Context) {
	var countryCodeInfo model.CountryCodeInfo
	err := c.Bind(&countryCodeInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	holidays, err := api.HolidayService.SendToHolidayWebService(countryCodeInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, holidays)
}
