### Intro
A simple golang api for a bank

### List of dependencies
1. gorilla mux
`go get github.com/gorilla/mux`
2. pq
`go get github.com/lib/pq` 

### Docker
`docker run --name bank-postgres -e POSTGRES_PASSWORD=gobank -p 5433:5432 -d postgres`
Were binding the postgres connection to port 5433 because we have an active connection on port 5432.
`docker ps`
check running containers
`docker images`
show a list of docker images.