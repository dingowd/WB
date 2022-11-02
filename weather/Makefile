up:
	docker build -f DockerWeather . --tag weather
	docker build -f DockerPostgres . --tag postgres-con
	docker-compose up -d
down:
	docker-compose down
start:
	docker-compose start postgres
	docker-compose start weather
stop:
	docker-compose stop