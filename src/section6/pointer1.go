package main

import (
  "fmt"
)

func main() {
  //포인터
  //Go : 포인터 지원(C)
  //변수의 지역성, 연속된 메모리 참조 ..., 힙, 스택
  //파이썬, 자바(Jre), 컴파일러, 인터프리터 내부적으로 포인터가 동작된다.
  //포인터지원x (파이썬, C#, JAVA등)
  //주소의 값은 직접 변경 불가(잘못된 코딩으로 인한 버그 방지)
  //*(에스터리스크)로 사용
  //nil 로 초기화(nil == 0)


  //예제1
  var a *int // 방법1
  var b *int = new(int) //방법2 정석

  fmt.Println(a) // 할당이 없을땐, nil로 초기화
  fmt.Println(b) // new로 해서 제대로 할당된 경우 주소값을 가지고 있음.


  i := 7

  fmt.Println("ex1 : ", i, &i)

  a = &i // & 주소값 전달
  b = &i

  //*a = 77

  fmt.Println("ex2 : ", a, &i)
  fmt.Println("ex2 : ", &a)
  fmt.Println("ex2 : ", *a) //역참조

  fmt.Println("ex2 : ", b, &i)
  fmt.Println("ex2 : ", &b)
  fmt.Println("ex2 : ", *b) //역참조


  var c = &i
  d := &i

  *d = 10


  fmt.Println("ex2 : ", c, &i)
  fmt.Println("ex2 : ", &c)
  fmt.Println("ex2 : ", *c) //역참조

  fmt.Println("ex2 : ", d, &i)
  fmt.Println("ex2 : ", &d)
  fmt.Println("ex2 : ", *d) //역참조
}
