package main

import (
        "fmt"
    )



func main(){
    fullName := func(first_name, last_name string) string{
        return fmt.Sprintf("%s %s", first_name, last_name)
    }
    fullName = "Nicholas Karimi"
    welcomeMessage := sayHello("Bug", "Byte", fullName)
}


func sayHello(first, last string, fn func(string, string) string) string{
    fullName := fn(first, last)
    return fmt.Sprintf("Welcome %s", fullName)
}

