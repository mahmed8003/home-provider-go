package server

import (
	"encoding/json"
	"home-provider/boom"

	"github.com/go-ozzo/ozzo-validation"

	routing "github.com/go-ozzo/ozzo-routing"
)

func errorHandler(c *routing.Context) error {

	err := c.Next()
	if err == nil {
		return nil
	}

	switch err.(type) {
	case *boom.Boom:
		boomError := err.(*boom.Boom)
		c.Response.WriteHeader(boomError.Code)
		c.Write(boomError)

	case validation.Errors:
		errors := err.(validation.Errors)
		boomError := boom.BadRequest("Invalid form data")
		boomError.Detail = errors
		c.Response.WriteHeader(boomError.Code)
		c.Write(boomError)

	case *json.SyntaxError:
		e := err.(*json.SyntaxError)
		boomError := boom.BadRequest("Invalid json")
		boomError.Detail = e.Error()
		c.Response.WriteHeader(boomError.Code)
		c.Write(boomError)

	default:
		return err
	}

	return nil
}
