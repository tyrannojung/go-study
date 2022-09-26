package main

import (
  "fmt"
)

func main() {
  //슬라이스 복사
  //copy(복사 대상, 원본)
  //make로 공간을 할당 후 복사해야 한다.
  //복사 된 슬라이스 값 변경해도 원본에는 영향 없음

  //예제1(복사)
  slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  slice2 := make([]int, 5)
  slice3 := []int{}

  copy(slice2, slice1)
  copy(slice3, slice1)

  fmt.Println("ex1 : ", slice2) //허용한 값까지만 복사
  fmt.Println("ex1 : ", slice3) //복사안됨


  //예제2
  a := []int{1, 2, 3, 4, 5}
  b := make([]int, 5)

  copy(b, a) //카피메소드를 사용하면 배열처럼 사용가능, 원본의 데이터가 바뀌지 않음
  b[0] = 7
  b[4] = 10

  fmt.Println("ex2 : ", a)
  fmt.Println("ex2 : ", b)
  fmt.Println()


  //예제3
  //주의! 부분적으로 슬라이스 추출은 참조 --> 원본 값 변경됨.
  c := [5]int{1, 2, 3, 4, 5}
  d := c[0:3] //추출해서 가져오는건 참조타입이다.

  d[1] = 7 // 원본값에 영향이 간다.

  fmt.Println("ex3 : ", c)
  fmt.Println("ex3 : ", d)


  //예제4
  e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  f := e[0:5:7]

  fmt.Println("ex4 : ", len(f), cap(f)) // cap에서 7을 주므로, 해당마지막 매개변수는 용량지정
  fmt.Println("ex4 : ", f)
}
