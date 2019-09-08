package locator

import (
	"net/http"

	"github.com/go-openapi/strfmt"
	"github.com/scukonick/eta/clients/cars/client"
)

type Service struct {
	httpClient  *http.Client
	carsService *client.CarsService
}

func NewService(httpClient *http.Client, endpoint string) *Service {
	x := client.NewHTTPClientWithConfig(strfmt.Default,
		&client.TransportConfig{
			Host:     endpoint,
			Schemes:  []string{"https"},
			BasePath: "/fake-eta",
		})

	return &Service{
		httpClient:  httpClient,
		carsService: x,
	}

}
