package main 


import (
	"fmt"
)

func waitForDbs(numDBs int, dbChan chan struct{}){
	for i :=0; i < numDBs; i++{
		<-dbChan
	}
}

func test(numDBs int){
	dbChan := getDatabasesChannel(numDBs)
	fmt.Printf("Waiting for %v databases...\n", numDBs)
	waitForDbs(numDBs, dbChan)
	fmt.Println("All databases are available!")
}

func main() {
	test(3)
	test(4)
	test(5)

}


func getDatabasesChannel(numDBs int) chan struct{} {
	ch := make(chan struct{})
	go func(){
		for i :=0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online \n", i+1)
		}
	}()
	return ch
}