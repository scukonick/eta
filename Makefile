all:
	mkdir -p clients/cars
	mkdir -p clients/predict
	swagger generate client -f car-swagger.yml -t clients/cars
	swagger generate client -f predict-swagger.yml -t clients/predict
