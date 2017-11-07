package main
 
 import (
     "io"
     "log"
     "net/http"
     "strconv"
     "fmt"
 )
 
 var iCnt int = 0;
 
 func helloHandler(w http.ResponseWriter, r * http.Request) {
     iCnt++;
     str := "Hello world ! friend(" + strconv.Itoa(iCnt) + ")"
     io.WriteString(w, str)
     fmt.Println(str)
 }
 
 func main() {
    fmt.Println("hellogo start ...")     
ht := http.HandlerFunc(helloHandler)
     if ht != nil {
         http.Handle("/hello", ht)
     }
     err := http.ListenAndServe(":8090", nil)
     if err != nil {
         log.Fatal("ListenAndServe: ", err.Error())
     }
 }
