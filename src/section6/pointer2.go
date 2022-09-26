package main

import (
  "fmt"
)

func main() {
  //예제1
  i := 7
  p := &i

  fmt.Println("ex1 : ", i, *p, &i, p)

  *p++ // 주소값이 아니고 역참조이므로 연산가능
  fmt.Println("ex1 : ", i, *p, &i, p)

  *p = 7777 //포인터 변수 역 참조 값 변경
  fmt.Println("ex1 : ", i, *p, &i, p)

  i = 77 //원본 변경
  fmt.Println("ex1 : ", i, *p, &i, p)

}
