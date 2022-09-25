package main

import (
  "fmt"
  solo "section4/lib3"
)

// 패키지 로드시 가장 먼저 호출
//가장 먼저 초기화 되는 작업 작성 시 유용.

var num int32

func init() {
  num = 30
  fmt.Println("Main int start!")
}

func main() {
  fmt.Println("10 보다 큰수? : ", solo.Jung(num))
}
