package main

import "fmt"

func main() {
  //Swit
  //예제 1
  a := 7
  switch {
    case a < 0:
      fmt.Println(a, "는 음수")
    case a == 0:
      fmt.Println(a, "는 0")
    case a > 0:
      fmt.Println(a, "는 양수")
  }

  //예제 2
  switch b:= 27; {
  case b < 0:
      fmt.Println(b, "는 음수")
    case b == 0:
      fmt.Println(b, "는 0")
    case b > 0:
      fmt.Println(b, "는 양수")
  }

  //예제 3
  // c의 위치를 잘봐야함. 바로 같다체크(c == go)
  switch c := "go"; c {
  case "go":
    fmt.Println("Go!")
  case "java":
    fmt.Println("Java!")
  default:
    fmt.Println("일치하는 값 없음")
  }

  //예제 4
  switch c := "go"; c + "lang" {
  case "golang":
    fmt.Println("golang")
  case "java":
    fmt.Println("Java!")
  default:
    fmt.Println("일치하는 값 없음")
  }

  //예제 5
  switch i, j := 20, 30; {
  case i < j:
    fmt.Println("i는 j보다 작다")
  case i == j:
    fmt.Println("i와 j는 같다")
  case i > j:
    fmt.Println("i는 j보다 크다")
  }

}
