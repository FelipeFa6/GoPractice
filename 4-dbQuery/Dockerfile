FROM postgres:latest

ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=secret
ENV POSTGRES_DB=bank

COPY sql/DDL.sql /docker-entrypoint-initdb.d/

EXPOSE 5432
CMD ["postgres"]

