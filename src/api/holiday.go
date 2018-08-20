package api

import (
	"model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Api struct {
	Flow func(model.CountryCodeInfo) model.HolidayInfo
}

func (api Api) HolidayHandler(c *gin.Context) {
	var countryCodeInfo model.CountryCodeInfo
	err := c.Bind(&countryCodeInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	holidays := api.Flow(countryCodeInfo)

	c.JSON(200, holidays)
}
