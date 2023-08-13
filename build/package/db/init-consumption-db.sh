#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	\copy consumptions FROM '/docker-entrypoint-initdb.d/consumption.csv' DELIMITER ',' CSV HEADER;
EOSQL
