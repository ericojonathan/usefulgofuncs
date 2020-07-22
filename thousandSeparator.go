package main

import (
  "fmt"
  "strconv"
  "strings"
)

func main() {
  fmt.Println("vim-go")
}

func thousandSeparator(nominal int, sep string) string {
  n := strconv.Itoa(nominal)
  if len(n) < 4 { 
    return n
  } else {
    var res string
    i := 0
    for j := len(n) - 1; j >= 0; j-- {
      j++ 
      res = string(y[i]) + res 
      if i > 1 && i%3 == 0 { 
        res = sep + res 
      }   
    }   
    return strings.Trim(res, ".")
  }
}
~                         
