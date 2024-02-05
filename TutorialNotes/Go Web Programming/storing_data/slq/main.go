package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

type Post struct{
    ID int
    Title string
    Content string
    Author string
}

// initalize the db

var Db *sql.DB

func init(){
    var err error
    Db, err = sql.Open("postgres", "user=admin dbname=gwp password=Incorrect sslmode=disable")

    if err != nil {
        panic(err)
    }

}

func Posts(limit int) (posts []Post, err error) {

    rows, err := Db.Query("select id, title, content, author from posts limit $1", limit)

    if err != nil {
        return
    }

    for rows.Next() {
        post := Post{}
        err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author)
        if err != nil {
            return
        }

        posts = append(posts, post)
    }
    rows.Close()
    return
}

func GetPost(id int) (post Post, err error) {
    post = Post{}
    err = Db.QueryRow("select id, title, content, author form posts where id = $1", id).Scan(&post.ID, &post.Title, &post.Content, &post.Author)
    return
}

func (post *Post) Create() (err error) {
    statement := "insert into posts (title, content, author) values ($1, $2, $3) returning id"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }

    defer stmt.Close()
    err = stmt.QueryRow(post.Title, post.Content, post.Author).Scan(&post.ID)
    return
}

func (post *Post) Update() (err error){
    resp, err := Db.Exec("update posts set title = $2, content = $3, author =  $4 where id = $1", post.ID, post.Title, post.Content, post.Author)
    if err != nil {
        return err
    }
    fmt.Printf("%+v\n", resp)
    return nil
}

func (post *Post) Delete() (err error){
    res, err := Db.Exec("delete from posts where id = $1", post.ID)
    if err != nil {
        return err
    }
    fmt.Printf("%+v\n", res)
    return nil
}

func main(){
    post := Post{Title: "API made easy", Content: "Learn API the rigth way", Author: "Nicholas Karimi"}
    post2 := Post{Title: "Bear Republic", Content: "Love from the bear republic", Author: "Kenny Paul"}
    post3 := Post{Title: "Gin vs Fiber", Content: "Best webframework for Go fullstack development", Author: "Nicholas Karimi"}


    fmt.Println(post)
    post.Create()
    post2.Create()
    post3.Create()
    fmt.Println(post)
}

