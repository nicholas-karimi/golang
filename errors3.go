package main

import (
    "errors"
    "fmt"
)

func DevideByZeror(a, b float64)(float64, error){
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main(){
    results, err := DevideByZeror(10, 0)

    if err != nil {
        fmt.Println("Error", err)
        return
    }
    fmt.Println("results: ", results)
}
