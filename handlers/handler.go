package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Data holds all data from request
type Data struct {
	Param  *Param  `json:"param,omitempty"`
	Header *Header `json:"header,omitempty"`
	Body   *Body   `json:"body,omitempty"`
}

// Param holds url params
type Param struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Header holds request Header data
type Header struct {
	Token string `json:"token,omitempty"`
}

// Body holds request body data
type Body struct {
	Message string `json:"message,omitempty"`
}

// GetID is our example for handler with GET method
func GetID(c echo.Context) error {
	data := Data{
		Param: &Param{
			ID:   c.Param("id"),
			Name: c.QueryParam("name"),
		},
	}

	return c.JSON(http.StatusOK, data)
}

// GetAll is our example for handler with POST method
func GetAll(c echo.Context) error {
	data := Data{
		Param: &Param{
			ID:   c.Param("id"),
			Name: c.QueryParam("name"),
		},
		Header: &Header{
			Token: c.Request().Header.Get("token"),
		},
		Body: &Body{
			Message: c.FormValue("message"),
		},
	}

	return c.JSON(http.StatusOK, data)
}
