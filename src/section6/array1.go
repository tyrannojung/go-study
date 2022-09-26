package main

import (
  "fmt"
)

func main() {
  //배열
  //배열은 용량과 길이가 같다.
  //배열 vs 슬라이스의 차이점이 중요하다.
  //배열 : 길이고정, 값타입(복사 전달), 전체 비교연산자 사용가능
  //슬라이스 : 길이가변, 참조 타입(참조 값 전달), 비교 연산자 사용불가
  //대부분 슬라이스를 사용한다.

  //cap() : 배열, 슬라이스 용량
  //len() : 배열, 슬라이스 개수


  //예제1
  var arr1 [5]int
  var arr2 [5]int = [5]int{1, 2, 3, 4, 5} // 배열은 개수와, 안에 들어가는 데이터 타입도 맞아야 한다.
  var arr3 = [5]int{1, 2, 3, 4, 5}
  arr4 := [5]int{1, 2, 3, 4, 5}
  arr5 := [5]int{1, 2, 3}
  arr6 := [...]int{1, 2, 3, 4, 5} //길이에 확신이 없을때. 배열 크기 자동 맞춤
  arr7 := [5][5]int{
    {1, 2, 3, 4, 5},
    {6, 7, 8, 9, 10},
  }

  arr1[2] = 5

  fmt.Printf("%d %v\n", len(arr1), arr1) // int는 기본값이 0임
  fmt.Printf("%d %v\n", len(arr2), arr2)
  fmt.Printf("%d %v\n", len(arr3), arr3)
  fmt.Printf("%d %v\n", len(arr4), arr4)
  fmt.Printf("%d %v\n", len(arr5), arr5)
  fmt.Printf("%d %v\n", len(arr6), arr6)
  fmt.Printf("%d %v\n", len(arr7), arr7)

  //예제2
  arr8 := [5]int{1, 2, 3, 4, 5}
  arr9 := [5]int{
      1,
      2,
      3,
      4,
      5,
  }
  arr10 := [...]string{"kim", "lee", "park"}
  fmt.Printf("%-5T %d %v\n", arr8, len(arr8), arr8) // %-5T는 fmt.print안 크기 및 자료형을 보여줌
  fmt.Printf("%-5T %d %v\n", arr9, len(arr9), arr9)
  fmt.Printf("%-5T %d %v\n", arr10, len(arr10), arr10)


}
