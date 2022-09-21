package main

import "fmt"

func main() {
  //상수
  //const 사용 초기화, 한 번 선언 후 값 변경 금지, 고정 된 값 관리 용
  const a string = "Test1"
  const b = "Test2" // 내부적으로 자동 컨버팅이 된다.
  const c int32 = 10 * 10
  //const d = getHeight() 함수가 항상 같은 값을 줄 수 있을지 모르므로 애초에 컴파일 오류
  const e = 35.6
  const f = false

  /*
    애러 발생
    const g string
    g = "Test3"
    선언과 동시에 초기화가 되지 않음
  */

  fmt.Println("a : ", a, "b : ", b, "c : ", c, "e : ", e, "f : ", f)


}
