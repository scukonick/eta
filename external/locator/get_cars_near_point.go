package locator

import (
	"context"

	"github.com/pkg/errors"
	"github.com/scukonick/eta/clients/cars/client/operations"
	"github.com/scukonick/eta/external"
)

func (s *Service) GetCarsNearPoint(ctx context.Context, lat, lng float64) ([]external.Car, error) {
	resp, err := s.carsService.Operations.GetCars(&operations.GetCarsParams{
		Lat:        lat,
		Lng:        lng,
		Limit:      10,
		Context:    ctx,
		HTTPClient: s.httpClient,
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to get cars near")
	}

	result := make([]external.Car, 0, len(resp.Payload))
	for _, r := range resp.Payload {
		result = append(result, external.Car{
			ID:  r.ID,
			Lat: r.Lat,
			Lng: r.Lat,
		})
	}

	return result, nil
}
