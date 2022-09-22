package main

import "fmt"

func main() {
  //예제1
  //여러가지 값을 나열해서 조건식에 맞는지 확인 가능
  a := 30 / 15
  switch a {
  case 2, 4, 6: //i가 2, 4, 6인 경우
     fmt.Println("a ->", a, "는 짝수")
  case 1, 3, 5: //i가 1, 3, 5인 경우
    fmt.Println("a -->", a, "는 홀수")
  }

}
