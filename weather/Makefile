dockerbuild:
	docker build -f DockerWeather . --tag weather
	docker build -f DockerPostgres . --tag postgres-con
dockerup:
	docker-compose up -d
dockerdown:
	docker-compose down