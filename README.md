# Pick-up point service
### A service for admins and employees, where they can manage pick-up point orders.

## Getting started
### Installation
1. Clone the repository
 ```sh
   git clone https://github.com/agadilkhan/pickup-point-service.git
   ```
2. Install go packages
 ```sh
   go mod tidy
   ```

### Migrate
#### Create .env file and fill values given in .env.example file
1. To migrate the data and tables on auth service
 ```
    make auth-migrateup
  ```
2. To migrate the data and tables on user service
 ```
    make user-migrateup
  ```
3. To migrate the data and tables on pickup service
 ```
    make pickup-migrateup
  ```

### Launch
#### Create config.yaml file inside each folder that are located inside the config folder. And fill values given in config.yaml.example.
1. To launch the auth-service:
 ```
    make start-auth
  ```
2. To launch the user-service:
 ```
    make start-user
  ```
3. To launch the pickup-service:
 ```
    make pickup-auth
  ```

## Built With
* [![Golang][Golang-badge]][Golang-url]
* [![Gin][Gin-badge]][Gin-url]
* [![gRPC][gRPC-badge]][gRPC-url]
* [![PostgreSQL][PostgreSQL-badge]][PostgreSQL-url]
* [![Redis][Redis-badge]][Redis-url]
* [![Kafka][Kafka-badge]][Kafka-url]
* [![Docker][Docker-badge]][Docker-url]
* [![Swagger][Swagger-badge]][Swagger-url]
* [![Grafana][Grafana-badge]][Grafana-url]
* [![Prometheus][Prometheus-badge]][Prometheus-url]