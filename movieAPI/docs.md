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

### Make
A task runner.
`Make` helps with testing, building, cleaning, and installing Go projects. 

### Makefiles
`make` uses the `Makefile` as its source of commands to execute and these commands are defined as a rules in the `Makefile`. A single rule defines target, dependencies, and the recipe of the `Makefile`.
#### Terminology
1. `Target:` Targets are the main component of a Makefile. The make command executes the recipe by its target name. Commands like `build, run, and build_and_clean`, are called `targets`. _Targets_ are the interface to the commands you want to execute.
2. `Dependencies`: A target can have dependencies that need to be executed before running the target. For example, the `build_and_clean` command can have two dependencies: build and run.
3. `Recipe`: Recipes are the actual commands that will be executed when the target is run. A recipe can be a single command or a collection of commands. You can specify multiple commands in a target using a line break. In the `makeAPI` makefile, the recipe for the run target is `./${BINARY_NAME}`. *A recipe should always contain a tab at the start.*

>*Variables*  Makefiles also have a mechanism to use variables.These are useful when you want the same configs or outputs to be used for different targets. In this project makefile, I have added the `BINARY_NAME `variable, which is reused across different targets.
>The variable can be substituted by enclosing it `${<variable_name>}`.
```go
    BINARY_NAME=movieAPI

    run:
        ./${BINARY_NAME}
 ```
 Variables can be defined either by using `= or :=`

###### Some Useful Tips
- To make comments in a Makefile, you can simply add a `# `before any line.
- To disable printing the recipe while running the target command, use `@` before the recipe.
- Use the `SHELL` variable to define the shell to execute the recipe.
- Define the `.DEFAULT_GOAL` with the name of the target.

### Multi-platform Builds
Golang supports multi-platform builds but it needs multiple commands to build the binaries for different platforms, which means more time-consuming and repetitive steps when building binaries.
`Make` utility can be used to automate these tasks.
1. Linux
`GOARCH=amd64 GOOS=linux go build -o ./target/movie-api-linux main.go`
Will build a program for `Unix` based systems as the target.
Run
`./movie-api-linux`
2. Windows
`GOARCH=amd64 GOOS=windows go build -o ./target/movie-api-windows main.go`
Builds an executable program for `windows` based `os`. The executable rogram will be inside the `target` folder.
Running the windows executable on linux will throw the following error `zsh: exec format error: ./movie-api-windows` since that is not the target operating system.


### Cleaning up
`make clean`
`make lint`
