package main

import (
  "fmt"
)

// 패키지 로드시 가장 먼저 호출
//가장 먼저 초기화 되는 작업 작성 시 유용.

func init() {
  fmt.Println("Init Method Start!")
}

func main() {
    fmt.Println("Main Method Start!")
}
