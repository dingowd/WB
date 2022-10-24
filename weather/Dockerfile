FROM postgres:latest
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=masterkey
ADD migrations/createDB.sql /
ADD migrations/migrate.sql /