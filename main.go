package main

import "fmt"

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, Grigory!")
}

func main() {
  http.HadleFunc("/", handler)
  fmt.Println(":8080", nil)
}
