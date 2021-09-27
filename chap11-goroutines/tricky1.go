package main

import (
  "fmt"
  "sync"
)

func printHi(x string){
  fmt.Println(x)
}


func Tricky1(){
  // runs function async. which will be delegated to treads running in 
  // background
  go printHi("hello from tricky1") 
 
  // main func closes before go routine get a chance to run
  fmt.Println("Hi2")
}

func Tricky2(){

  for i := 0; i < 5; i++ {
    go printHi(fmt.Sprintf("hello from tricky2: %d",i))
  }

  fmt.Println("Hi3")
}

func Tricky3(){
  var wg sync.WaitGroup

  // setup counter
  wg.Add(10)

  for i := 0; i < 10; i++ {
    // goroutine are run async and run in background treads.
    // they are lightweight and easy to setup
    go func(x int){
      fmt.Println("tricky3: ",x)
      wg.Done()
    }(i)// closures are by ref, you need to pass values your self
  }

  // wait till counter is done
  fmt.Println("start waiting")
  wg.Wait()
  fmt.Println("done waiting")


}
