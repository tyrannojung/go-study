package main

import (
  "fmt"
  "section4/lib"
)

func main() {
  // go에서 접근제어자는 소문자 대문자로 구분함.
  fmt.Println("10보다 큰 수? : ", lib.CheckNum(4))
}
