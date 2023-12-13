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

[Golang-badge]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Golang-url]: https://golang.org/
[Gin-badge]: https://img.shields.io/badge/Gin-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Gin-url]: https://gin-gonic.com/
[Echo-badge]: https://img.shields.io/badge/Echo-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Echo-url]: https://echo.labstack.com/
[PostgreSQL-badge]: https://img.shields.io/badge/PostgreSQL-336791?style=for-the-badge&logo=postgresql&logoColor=white
[PostgreSQL-url]: https://www.postgresql.org/
[Redis-badge]: https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white
[Redis-url]: https://redis.io/
[Kafka-badge]: https://img.shields.io/badge/Apache%20Kafka-231F20?style=for-the-badge&logo=apache-kafka&logoColor=white
[Kafka-url]: https://kafka.apache.org/
[gRPC-badge]: https://img.shields.io/badge/gRPC-00ADD8?style=for-the-badge&logo=go&logoColor=white
[gRPC-url]: https://grpc.io/
[Docker-badge]: https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white
[Docker-url]: https://www.docker.com/
[Swagger-badge]: https://img.shields.io/badge/Swagger-85EA2D?style=for-the-badge&logo=swagger&logoColor=black
[Swagger-url]: https://swagger.io/
[Grafana-badge]: https://img.shields.io/badge/Grafana-F46800?style=for-the-badge&logo=grafana&logoColor=white
[Grafana-url]: https://grafana.com/
[Prometheus-badge]: https://img.shields.io/badge/Prometheus-E6522C?style=for-the-badge&logo=prometheus&logoColor=white
[Prometheus-url]: https://prometheus.io/