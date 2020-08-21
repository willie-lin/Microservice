package _controllers

import (
	"Microservice/models"
	"Microservice/stock"
	"bytes"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

// GetNews...

func GetNews(c echo.Context) (err error)  {
	// read the body content
	var bodyBytes []byte
	if c.Request().Body != nil{
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	u := new(models.Company)
	//
	er := c.Bind(u)
	if er != nil {

		panic(err)
	}
	r := stock.News(u.Ticker)
	return c.JSON(http.StatusCreated, r)
}