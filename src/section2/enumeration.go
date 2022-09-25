package main

import "fmt"

func main() {
  //열거형
  //상수를 사용하는 일정한 규칙에 따라 숫자를 계산 및 증가시키는 묶음
  // const (
  //   Jan = 1
  //   Feb = 2
  //   Mar = 3
  //   Apr = 4
  //   May = 5
  //   Jun = 6
  // )
  //
  // fmt.Println(Jan)
  // fmt.Println(Feb)
  // fmt.Println(Mar)
  // fmt.Println(Apr)
  // fmt.Println(May)
  // fmt.Println(Jun)


  const (
    A = iota
    B
    C
  )

  fmt.Println(A, B, C)

  const (
    Jan = iota + 1
    Feb
    Mar
    Apr
    May
    Jun
  )

  fmt.Println(Jan)
  fmt.Println(Feb)
  fmt.Println(Mar)
  fmt.Println(Apr)
  fmt.Println(May)
  fmt.Println(Jun)


 // iota는 규칙적으로 증가하는데 중간값 생략가능.
  const (
    _ = iota
    D
    _
    F
  )

  const (
    _ = iota + 0.75 * 2
    DEFAULT
    _
    GOLD
    PLATINUM
  )

  fmt.Println("D ", D, "F ", F)
  fmt.Println("DEFAULT ", DEFAULT, "GOLD ", GOLD, "PLATINUM ", PLATINUM)

}
