# Consumption API #

This repository implements an API to get consumption information from different services on Golang Echo.

## Install

Make sure you are using Docker

```
docker-compose -f deployments/docker-compose.yml --project-directory . up
```

## Routes

```
curl --location --request GET 'http://localhost:80/api/consumption?meters_ids=1&start_date=2023-06-01&end_date=2023-07-10&kind_period=monthly'
curl --location --request GET 'http://localhost:80/api/consumption?meters_ids=1&start_date=2023-06-01&end_date=2023-06-26&kind_period=weekly'
curl --location --request GET 'http://localhost:80/api/consumption?meters_ids=1&start_date=2023-06-01&end_date=2023-06-10&kind_period=daily'
```

The service provides a single endpoint with three different parameters:  
* ***meters_ids***
Integer or list of integers

* ***start_date***
String with the following format YYYY/MM/DD

* ***end_date***
String with the following format YYYY/MM/DD

* ***kind_period***
String can be any of these: monthly, daily, weekly
