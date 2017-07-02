package main

import "net/http"
import "os"
import "strconv"

func main() {
  pwd, _ := os.Getwd()
  port := 8000
  i := 1

  if len(os.Args) >= i+2 && (os.Args[i] == "-p" || os.Args[i] == "--path") {
    port, _ = strconv.Atoi(os.Args[i + 1])
    i += 2
  }

  if len(os.Args) >= i+1 {
    pwd = os.Args[i]
    i += 1
  }

  fs := http.FileServer(http.Dir(pwd))
  http.Handle("/", fs)
  http.ListenAndServe("0.0.0.0:" + strconv.Itoa(port), nil);
}
