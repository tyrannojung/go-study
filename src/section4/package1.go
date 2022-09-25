package main

/*
//선언 방법1
import "fmt"
import "os"

*/
//선언 방법 2
import (
  "fmt"
  "os"
)

func main() {
  // 패키지 : 코드 구조화(모듈화) 및 재사용
  // 모든 언어들은 발전하면서 api, component화가 잘 되어있다.
  // 결합도는 최대한 낮추고, 응집도는 최대한 높임 --> 재사용성이 좋음(클린코드)
  // Go : d패키지 단위의 독립적이고 작은 단위로 개발 --> 작은 패키지를 결합해서 프로그래밍을 작성할 것 권고
  // 패키지 이름 = 디렉토리 이름으로 작성된다.
  // 같은 패키지내 소스파일은 디렉토리명을 패키지 명으로 사용한다.
  // 네이밍 규칙 : 소문자 private, 대문자 : public
  // fmt와 같은 표준 패키지들은 go ROOt에 있따.
  // Go : main 패키지는 특별하게 인식 --> 컴파일러 공유 라이브러리라고인식x , 프로그램 시작점 start point로 생각


  var name string
  fmt.Println("이름은 ? : ")

  fmt.Scanf("%s", &name)
  fmt.Fprintf(os.Stdout, "Hi! %s\n", name)
}
