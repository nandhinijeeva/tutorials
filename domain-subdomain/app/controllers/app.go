package controllers

import (
	"aahframework.org/aah.v0"
	"github.com/go-aah/tutorials/domain-subdomain/app/models"
)

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Index method is application home page.
func (a *AppController) Index() {
	data := aah.Data{
		"Greet": models.Greet{
			Message: "Tutorials  - Domain, Subdomain, and Wildcard Subdomain",
		},
	}

	a.Reply().Ok().HTML(data)
}
