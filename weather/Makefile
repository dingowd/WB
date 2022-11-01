up:
	docker build -f DockerPostgres . --tag postgres-con
	docker build -f DockerWeather . --tag weather
	docker-compose up -d
down:
	docker-compose down