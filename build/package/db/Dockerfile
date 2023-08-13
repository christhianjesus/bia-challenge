FROM postgres:alpine

COPY build/package/db/consumption.csv /docker-entrypoint-initdb.d
COPY build/package/db/consumption_schema.sql /docker-entrypoint-initdb.d
COPY build/package/db/init-consumption-db.sh /docker-entrypoint-initdb.d
