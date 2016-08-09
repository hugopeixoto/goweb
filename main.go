package main

import "net/http"
import "os"

func main() {
  pwd, _ := os.Getwd()

  fs := http.FileServer(http.Dir(pwd))
  http.Handle("/", fs)
  http.ListenAndServe("0.0.0.0:8000", nil);
}
