package codeinformation

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/danilotadeu/r-customer-code-information-provider/app"
	"github.com/gofiber/fiber/v2"
)

type apiImpl struct {
	apps *app.Container
}

// NewAPI codeinformation function..
func NewAPI(g fiber.Router, apps *app.Container) {
	api := apiImpl{
		apps: apps,
	}

	g.Get("/", api.CodeInformation)
}

func (p *apiImpl) CodeInformation(c *fiber.Ctx) error {
	b, err := ioutil.ReadFile("response.xml")
	/*
		TODO: Receber XML da request e tratar os erros da request
		TODO: Logar erros
	*/
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(http.StatusOK).Type("xml").SendString(string(b))
}
