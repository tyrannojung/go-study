package main

import (
  "fmt"
)

func main() {
  //문자열 표현
  //문자열 : UTF-8 인코딩(유니코드 문자 삽입) -> 바이트 수 주의!
  //문자열은 배열로 취급한다. --> 반복문에서 순회 및 인덱스로 접근 가능.  아스키코드로 출력, 실제문자로 출력하려면 printf를 사용.

  //예제1
  var str1 string = "Golang"
  var str2 string = "World"
  var str3 string = "고프로그래밍"

  // 코드값 표시
  fmt.Println("ex1 : ", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
  fmt.Println("ex1 : ", str2[0], str2[1], str2[2], str2[3], str2[4])
  fmt.Println("ex1 : ", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5])

  //예제2
  //한글자 한글자 정확히 표시
  fmt.Printf("ex2 : %c %c %c %c %c %c\n", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
  fmt.Printf("ex2 : %c %c %c %c %c\n", str2[0], str2[1], str2[2], str2[3], str2[4])
  fmt.Printf("ex2 : %c %c %c %c %c %c\n", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5]) //한글 깨짐


  conStr := []rune(str3)
  fmt.Printf("ex2 : %c %c %c %c %c %c\n", conStr[0], conStr[1], conStr[2], conStr[3], conStr[4], conStr[5])// 한글 정상 출력


  //예제3
  for i, char := range str1 {
    fmt.Printf("ex3 : %c(%d)\t\n", char, i)
  }

  for i, char := range str2 {
    fmt.Printf("ex4 : %c(%d)\t", char, i)
  }

}
