package timer

import (
	"context"

	"github.com/pkg/errors"
	"github.com/scukonick/eta/clients/predict/client/operations"
	"github.com/scukonick/eta/clients/predict/models"
	"github.com/scukonick/eta/external"
)

func (s *Service) GetETAs(ctx context.Context, cars []external.Car) ([]int, error) {
	predictSource := make([]models.Position, 0, len(cars))
	for _, car := range cars {
		predictSource = append(predictSource, models.Position{
			Lat: car.Lat,
			Lng: car.Lng,
		})
	}
	positionList := &models.PredictParamsBody{
		Target: models.Position{
			Lat: 56.003891,
			Lng: 37.428484706,
		},
		Source: predictSource,
	}
	predictParams := &operations.PredictParams{
		PositionList: positionList,
		Context:      context.Background(),
	}

	predictResp, err := s.predictService.Operations.Predict(predictParams)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get predictions")
	}

	if predictResp.Payload == nil {
		return nil, errors.Wrap(err, "predict resp payload is null")
	}

	if len(predictResp.Payload) != len(cars) {
		return nil, errors.Wrap(err, "invalid response length")
	}

	resp := make([]int, len(predictResp.Payload))
	for i, v := range predictResp.Payload {
		resp[i] = int(v)
	}

	return resp, nil
}
