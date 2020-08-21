package main

import (
	"Microservice/server"

)

func main() {
	//// create a new echo instance
	//e := echo.New()
	//
	//e.GET("/", func(c  echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!!!")
	//})

	//e.Logger.Fatal(e.Start(":1234"))

	s := server.NewServer()

	s.Start(":1323")
	


}
