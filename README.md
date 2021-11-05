# Stock manager

A little stub to a stock manager to keep track of items in your stock. It is a **REST API** and an example **CLI UI**.

There is a lot of known issues, like:

* No proper error handling when trying to remove more items than exists in storage.
* _very_ minimal logging.
* No persistant database.
* No auth whatsoever anywhere.

**Quick start:**

Requirements: `docker` and `docker-compose` installed, port 4000 not used by another application.

Build: `docker-compose build api api-tests cli`  
Start: `docker-compose up -d`

Run CLI commands via docker-compose:

`docker-compose run --rm cli I5`  
`docker-compose run --rm cli S3`  
`docker-compose run --rm cli L`  

Take a look at the Swagger in your browser over at `http://localhost:4000`

Clean up: `docker-compose down`

## REST API

`cd api`

The REST API is written in Go and uses [Fiber](https://gofiber.io/) to power the REST stuff.

### Unit tests

Some example tests are written and can be run by: `go test -v ./...`

### Integration tests

Integration tests are written in [Robot Framework](https://robotframework.org/). Run by: `docker-compose run --rm api-tests`

### Developer notes

To (re-)generate Swagger docs, run `swag init`.

## CLI

`cd cli`

The CLI is a tiny Python script. To try to show a little structure, functions have been moved to a lib-file. To run locally you need `python3` as well as python requests library (`pip3 install requests`). The API URL is consumed as an ENV variable.

Example of running the CLI locally:

`export API_URL=http://localhost:4000`  
`./cm I5`  
`./cm S3`  
`./cm L`

To omit the `./` put `cm` and `smlib.py` in your PATH.
