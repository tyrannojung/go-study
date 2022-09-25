package main

import (
  "fmt"
  "unicode/utf8"
)

func main() {
  //문자열
  //큰 따옴표 "", 백스쿼드``
  //Golang : 문자 char 타입은 존재하지 않는다. -> rune(int32)잘 사용하기 때문에 alias를 만듬. rune 문자 코드 값으로 표현
  //문자 : ''작음따옴표 로 작성
  //자주 사용하는 escape : \\ , \', \", \a(콘솔벨), \b(백스페이스) ...

  //예제1
  var str1 string = "c:\\go_study\\src\\" // escape 사용해야함. --> c:\go_study\src\
  str2 := `c:\go_study\src\` // excape를 사용하지 않더라도 그대로 사용할 수 있음. str1 == str2

  fmt.Println("ex1 : ", str1)
  fmt.Println("ex1 : ", str2)

  //예제2
  var str3 string = "Hello, world!"
  var str4 string = "안녕하세요."
  var str5 string = "\ud55c\uae00" //유니코드

  fmt.Println("ex2 : ", str3)
  fmt.Println("ex2 : ", str4)
  fmt.Println("ex2 : ", str5)

  //예제3 (중요)
  //길이와 크기를 구하는걸 명확하게 지원
  fmt.Println("ex3 : ", len(str3)) // len함수는 바이트를 구해줌.
  fmt.Println("ex3 : ", len(str4))

  //예제4
  fmt.Println("ex4 : ", utf8.RuneCountInString(str3))
  //아래 동일
  fmt.Println("ex4 : ", utf8.RuneCountInString(str4)) // 해당함수를 조금더 많이 사용하는편.
  fmt.Println("ex4 : ", len([]rune(str4))) // len으로 실제 길이 구하는법


}
