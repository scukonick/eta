package timer

import (
	"net/http"

	"github.com/go-openapi/strfmt"
	"github.com/scukonick/eta/clients/predict/client"
)

type Service struct {
	httpClient     *http.Client
	predictService *client.PredictService
}

func NewService(httpClient *http.Client, endpoint string) *Service {
	x := client.NewHTTPClientWithConfig(strfmt.Default,
		&client.TransportConfig{
			Host:     endpoint,
			Schemes:  []string{"https"},
			BasePath: "/fake-eta",
		})

	return &Service{
		httpClient:     httpClient,
		predictService: x,
	}
}
