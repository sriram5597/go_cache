# In-Memory-Cache-Golang

### Introduction

This is a In-Memory cache server which will be used to cache integers with expiry timestamp using the api call. It is implemented with mutex to handle concurrent requests.


### Project Structure

main.go -> Starts http server in port 8000

src/cache -> Manages the Map store. It exposes get key and set key functions to other packages

src/api/routers -> Created and map handlers with api routes

src/api/handlers -> Contains all the handler functions for api, request/response payload schema


### API Endpoints

Endpoint: /get-key?key=200  -> To fetch the value stored for that key

Endping: /set-key -> To store the value with respective to the given key in the payload
Payload:
{
    "key": 200,
    "value": 400,
    "expiry": 10
}
