package main

import "fmt"

func main() {
  //Go : 모호하거나, 애매한 문법 금지(go철학)
  //후치 연산자 허용 i++ but 전치 연산자 비허용 ++i x 문법애러
  //증감연산 반환값 x sum := i++ 사용불가
  //포인터(Pointer 허용, 연산 비허용)
  //주석(//, /**/) 상요법 숙지


  //예제1
  sum, i := 0, 0

  for i <= 100 {
    //sum += i++ // 예외발생
    sum += i
    i++
    //++i 예외발생(전위 증감)
  }
  fmt.Println("ex1",  sum)

}
