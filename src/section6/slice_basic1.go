package main

import (
  "fmt"
)

func main() {
  //길이고정x 가변 --> 동적으로 크기가 늘어남, 래퍼런스(참조 값) 타입
  //슬라이스 (길이 & 용량) 크기가 동적으로 할당 가능

  //2가지 선언방법 1. 배열처럼 선언, 2. make 함수 : make(자료형, 길이 용량(생략시 길이))

  //예제1
  //배열은 크기를 선언해주었지만,  slice는 선언하지 않는다. 즉 할당되지 않은 사이즈를 설정하지 않으면 slice.
  var slice1 []int
  slice2 := []int{}
  slice3 := []int{1, 2, 3, 4, 5}
  slice4 := [][]int{
    {1, 2, 3, 4, 5},
    {6, 7, 8, 9, 10},
  }

  //slice2[0] = 1 //가변형이므로, 값이 초기화되지 않은 상태에서는 수정할 수 없다.
  slice3[4] = 10 // 값 수정 가능
  fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
  fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
  fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
  fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)


  //예제2
  // make로 만드는경우 0으로 초기화됨. 길이와 용량이 할당되었으므로, 수정가능
  // 만약 용량을 초과하면 메모리의 제할당이 일어나 성능의 저하가 일어나며, 너무 적게 쓰면 메모리의 낭비가 일어남
  // 길이와 용량을 엔지니어 입장에서 성능을 고려해서 디테일하게 사용할 수 있게 GO가 제공한다.
  var slice5 []int = make([]int, 5, 10) //int타입, 길이가5, 용량10 (용량 생략하면 길이값)
  var slice6 = make([]int, 5) //해당 방법을 가장 많이 사용
  slice8 := make([]int, 5)
  slice7 := make([]int, 5, 100)

  slice6[2] = 7// 삽입 (make 수정가능)

  fmt.Printf("%-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
  fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
  fmt.Printf("%-5T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
  fmt.Printf("%-5T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8)


  //예제3
  var slice9 []int //슬라이스(길이와 용량 0) --> 해당상태를 nil이라 부름

  if slice9 == nil {
    fmt.Println("This is Nill Slice")
  }

}
