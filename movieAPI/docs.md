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
