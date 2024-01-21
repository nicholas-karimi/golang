### Introduction

This is a simple `golang` `CRUD` `API` for a movie store. We implement the follwoing endpoints:
1. `/movies`
this endpoint will return all movies in the store.
2. `/movie/{id}`
the endpoint returns a specific movie that matches the supplied id by the user.
3. `/movies`
this will act as the create movie endpoint.
4. `/movie/{id}/`
endpoint to delete a movie.

#### Dependencies
A list of all 3rd party packages utilized in this project.
1. _Gorilla Mux_
`go get "github.com/gorilla/mux"`


### Build for target device
1. Linux
`GOARCH=amd64 GOOS=linux go build -o ./target/movie-api-linux main.go`
Will build a program for `Unix` based systems as the target.
Run
`./movie-api-linux`
2. Windows
`GOARCH=amd64 GOOS=windows go build -o ./target/movie-api-windows main.go`
Builds an executable program for `windows` based `os`. The executable rogram will be inside the `target` folder.
Running the windows executable on linux will throw the following error `zsh: exec format error: ./movie-api-windows` since that is not the target operating system.