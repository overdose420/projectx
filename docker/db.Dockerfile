FROM mysql:8.0.23

EXPOSE 3350

COPY ./database/*.sql /docker-entrypoint-initdb.d/