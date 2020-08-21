package _controllers

import (
	"Microservice/models"
	"Microservice/stock"
	"bytes"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

func GetPrice(c echo.Context) (err error)  {

	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	u := new(models.Company)
	er := c.Bind(u)
	if er  != nil {
		panic(err)
	}
	r := stock.Price(u.Ticker)
	priceMap := map[string]string{"price": r}

	return c.JSON(http.StatusOK, priceMap)
}