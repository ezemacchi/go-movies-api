FROM postgres

ENV POSTGRES_USER="admin"
ENV POSTGRES_PASSWORD="admin"
ENV POSTGRES_DB="movies_db"

COPY movies.sql /docker-entrypoint-initdb.d/
