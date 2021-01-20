# Port manager

Port manager is my attempt of monorepo micro-services.

The whole project is meant to read `ports.json` from HTTP POST body and stream it to another micro-service.

## How it works

The project is split in 2 parts. ClientAPI and PortDomain.

PortDomain micro-service is responsible for any operations on ports (`upsert` and `get`) through GRPC.

While the ClientAPI micro-service is exposing the PortDomain service to outer world via HTTP.

### ClientAPI

ClientAPI exposes 2 endpoints:

```
GET http://localhost:4555/ports?id=4
POST http://localhost:4555/ports
```

Get method only gets the port with given id.

Post method is the upsert method which is writen in a way that can read the JSON like `ports.json` from root of the project. Handler is written in a way that it can read JSON of any size. It is reading JSON line by line and streaming to the PortDomain through GRPC stream.

### PortDomain

PortDomain exposes the GRPC server with 2 methods, `get` and `upsert`. Upsert is streaming method.

## Layout

The layout is derived from [here](https://github.com/golang-standards/project-layout).

The project is run from `cmd` directory. There are 2 binaries, `clientapi` and `portdomainsvc`.

Each micro-service has their own implementation in `internal` package while the code in `cmd` acts only as a glue.

In `deployment` directory are dockerfiles for each of the services. This can be simplified but it was not my focus. The database queries (service layer) of `PortDomain` was also not one of main focuses.

## Running the project

There are 3 ways of running the project.

* docker-compose

	Running with docker-compose is simple
	
	```bash
	docker-compose up
	```

	And the ClientAPI will be exposed on port 4555.

* hybrid (my approach when developing some of the services)

	Run just the parts of the services with docker-compose and the rest you start natively.
	
	```bash
	docker-compose up -d portdomain-database
	
	# add another terminal tab
	PORT=5001 DB='host=127.0.0.1 port=5434 user=root password=root dbname=db sslmode=disable' go run cmd/portdomainsvc/main.go
	
	# add another terminal tab
	PORT=4555 SVC=127.0.0.1:5001 go run cmd/clientapi/main.go
	```
	
	This way you can trash the database quicky.
	
* full native

	For this approach you need to have postgres database running.
	
	Then you need to create the database and load the schema.
	```bash
	psql #connect to your local databse with your credentials
	CREATE DATABSE ports;
	
	# now exit the psql (you can do it with CTRL-d) and load the schema
	psql -d ports < internal/portdomain/db/schema.sql 
	
	# start the PortDomain service (don't forget to input your database credentials in connection string)
	PORT=5001 DB='host=127.0.0.1 port=5432 user=username password=username? dbname=ports sslmode=disable' go run cmd/portdomainsvc/main.go
	
	# now again opet the new terminal (or tab) and run the api
	PORT=4555 SVC=127.0.0.1:5001 go run cmd/clientapi/main.go
	```

## Sending the json

You can send the json with following command `curl http://localhost:4555/ports -d @ports.json`.

And you can now get some of the ports by id `curl http://localhost:4555/ports?id=3`
