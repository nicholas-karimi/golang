package main

import (
    "fmt"
    )

type Post struct{
    ID int
    Title string
    Content string
    Author string
}

var postById map[int]*Post
var postByAuthor map[string][]*Post

func store(post Post){

    postById[post.ID] = &post
    postByAuthor[post.Author] = append(postByAuthor[post.Author], &post)
}


func main(){
    postById = make(map[int]*Post)
    postByAuthor = make(map[string][]*Post)
    
    post1 := Post{ID: 1, Title: "GO made easy", Content: "Go is an easy programming language", Author: "Nicholas Karimi"}
    post2 := Post{ID: 2, Title: "HTML", Content: "Html is a programming language", Author: "The Primegean"}

    store(post1)
    store(post2)

    fmt.Println(postById[1])

    for _, post := range postByAuthor["Nicholas Karimi"] {
        fmt.Println(post)
    }
}
