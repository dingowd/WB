FROM scratch
ADD cmd/command/weather /
ADD config/config.toml /config
ADD docs/swagger.json /docs
ADD docs/swagger.yaml /docs
ADD templates/cities.html /templates
ADD templates/detail.html /templates
ADD templates/details.html /templates
ADD templates/short.html /templates
ADD templates/shorts.html /templates
CMD ["/weather"]