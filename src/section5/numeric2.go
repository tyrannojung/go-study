package main

import (
  "fmt"
)

func main() {
  //데이터 타입 : 숫자형
  //정수형 문자 출력
  //예제1
  //아스키(영문)

  var char1 byte = 72 //10진수
  var char2 byte = 0110 // 8진수
  var char3 byte = 0x48 // 16진수

  //예제2
  //유니코드(한글)
  var char4 rune = 50556 //유니코드
  var char5 rune = 0142574 //44032(8진수)
  var char6 rune = 0xC57C //44032(16진수)

  //Printf 포매팅 + 뒤에 숫자 맵핑
  //%c는 문자맵핑
  fmt.Printf("%c %c %c\n", char1, char2, char3)

  //%d 숫자형(10진수)
  fmt.Printf("%d %d %d\n", char1, char2, char3)

  // 10진수, %o 8진수, %x 16진수
  fmt.Printf("%d %o %x\n", char1, char2, char3)

  fmt.Printf("%c %c %c\n", char4, char5, char6)

  //%d 숫자형(10진수)
  fmt.Printf("%d %d %d\n", char4, char5, char6)

  // 10진수, %o 8진수, %x 16진수
  fmt.Printf("%d %o %x\n", char4, char5, char6)

}
