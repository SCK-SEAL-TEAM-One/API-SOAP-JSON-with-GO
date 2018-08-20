package api

import (
	"model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Api struct {
	HolidayService func(model.CountryCodeInfo) (model.HolidayInfo, error)
}

func (api Api) HolidayHandler(c *gin.Context) {
	var countryCodeInfo model.CountryCodeInfo
	err := c.Bind(&countryCodeInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	holidays, err := api.HolidayService(countryCodeInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, holidays)
}
