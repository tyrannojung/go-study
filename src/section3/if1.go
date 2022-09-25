package main

import "fmt"

func main() {
  //IF문 : 반드시 Boolean 검사 -> 1, 0 (사용불가 : 자동 형 변환 불가)
  //소괄호 사용 x

  var a int = 20
  b := 20

  //예제1
  if a >= 15 {
    fmt.Println("15이상")
  }

  //예제2
  if b >= 25 {
    fmt.Println("25이상")
  }

  //애러발생 1
/*
  if b >= 25
  {
    fmt.Println("25이상")
  }
*/
// 위처럼 go언어는 25뒤에 ;(세미클론을 붙이므로 ) 애러. 괄호는 꼭 25 뒤부터 시작.

  //에러 발생 2
/*
    if b >= 25
      fmt.Println("25이상")

*/
  // 다른언어 같은 경우 조건 뒤 한문장이면 되지만, go는 안됨


  //에러발생 3
/*
  if c := 1; c {
    fmt.Println("True")
  }
*/
// 1은 true가 아니므로 에러
  if c := true; c {
    fmt.Println("True")

  }


  // 예제3

  if c := 40; c >= 35 {
    fmt.Println("35이상")
  }

  //c += 20  애러발생. c undifiend 짧은선언이므로 소멸
}
