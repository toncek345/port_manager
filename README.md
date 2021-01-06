### running the project

* start db from docker
  ```
  docker-compose up -d
  ```

* start domain svc
  ```
  cd cmd/portdomainsvc
  PORT=5001 DB='host=127.0.0.1 port=5434 user=root password=root dbname=db sslmode=disable' go run main.go
  ```
* start api
  ```
  cd cmd/clientapi
  PORT=5000 SVC=127.0.0.1:5001 go run main.go
  ```

* send json
  ```
  curl http://localhost:5000/ports -d @ports.json
  ```
