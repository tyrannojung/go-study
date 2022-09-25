package main

import "fmt"

func main() {
  const a, b int = 10, 99
  const c, d, e = true, 0.84, "test"
  // 괄호안에 있는 변수들은 어떤 영향을주는 상태 ?
  const (
    x, y int16 = 50, 90
    i, k       = "Data", 7776
  )

  fmt.Println(a, b, c, d, e)
  fmt.Println(x, y, i, k)

}
