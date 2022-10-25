FROM ubuntu:latest
ADD cmd/command/weather /service/
ADD config/config.toml /service/config/
ADD docs/swagger.json /service/docs/
ADD docs/swagger.yaml /service/docs/
ADD templates/cities.html /service/templates/
ADD templates/detail.html /service/templates/
ADD templates/details.html /service/templates/
ADD templates/short.html /service/templates/
ADD templates/shorts.html /service/templates/
CMD ["/bin/bash"]