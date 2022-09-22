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

  //예제2
  // go는 break가 없어도 자동으로 바로 빠져나간다.
  // 그러므로 반대로 fallthrough를 사용하면 통과시킨다.
  // case문의 맨 마지막에는 사용 불가
  switch e := "go"; e {
  case "java":
    fmt.Println("Java!")
    fallthrough
  case "go":
    fmt.Println("go!")
    fallthrough
  case "python":
    fmt.Println("python!")
    fallthrough
  case "ruby":
    fmt.Println("ruby!")
  case "c":
    fmt.Println("c")

  }


}
