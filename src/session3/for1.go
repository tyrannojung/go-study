package main

import "fmt"

func main() {
  //반복문 : for
  //Go에서 유일하게 제공되는 반복문
  //다양한 사용법 숙지

  //예제1
  for i := 0; i < 5; i++ {
    fmt.Println("ex1 : ", i)
  }

  // 에러 발생1
  /* 괄호 위치 주의
  for i := 0; i < 5; i++
  {

  }
  */

  //에러 발생2
  /* 괄호 필요
  for i := 0; i < 5; i++
    fmt.Println("ex1 : ", i)
  */

  //예제2 (무한 루프)
  // for {
  //   fmt.Println("ex2 : Hello Golang!")
  //   fmt.Println("ex2 : Infinite loop!")
  // }


  //예제3 (Range용법)
  // 첫번째 인자 index 만약 index표시안하고 싶으면 묵음(_)처리한다.
  loc := []string{"Seoul", "Busan", "Incheon"}
  for _, name := range loc {
    fmt.Println("ex3 : ", name)
  }
}
