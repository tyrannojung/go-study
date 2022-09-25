//라이브러리 접근제어(1)

package lib3

import (
  "fmt"
)

func init() {
  fmt.Println("lib Package > init start!")
}

func Jung(c int32) bool {
  return c > 10
}
