package main

import (
  "fmt"
)

func start(t string) string {
  fmt.Println("start : ", t)
  return t
}

func end(t string) {
  fmt.Println("end : ", t)
}

func a() {
  defer end(start("b")) //defer문의 바로 붙어있는 함수만 먼저 시작.(중첩함수는 주의. defer문은 왠만하면 함수 하나로 쓴다.)
  fmt.Println("in a")
}

func main() {
  //예제1

  a()
}
