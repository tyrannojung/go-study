## 💡 Introduction
- GO STUDY

## DAY3

- GO 특징
  - 모호하거나, 애매한 문법 금지(go철학)
  - 후치 연산자 허용 i++ but 전치 연산자 비허용 ++i x 문법애러
  - 증감연산 반환값 x sum := i++ 사용불가
  - 포인터(Pointer 허용, 연산 비허용)
  - 문장 끝 세미콜론(;) 주의 (자동으로 끝에 세미콜론 삽입)
  
- GO 코드서식지정
  - 터미널명령어 : gofmt -w ~~.go --> 코드가 뒤죽박죽이라도 자동 변환 

- 조건문(If)
  - 반드시 Boolean 검사 -> 1, 0 (사용불가 : 자동 형 변환 불가)
  - 조건 소괄호 사용안됨 && 조건 한문장일때 괄호생략 안됨
  
- 조건문(Switch) 
  - switch 뒤 표현식(Expression) 사용 가능 & case 뒤 표현식(Expression) 사용 가능
  - break가 default이므로, fallthough 존재
  - case 여러가지 나열가능
  
- 반복문(For)
  - Go에서 제공되는 '유일한' 반복문
  - session3 for 여러 사례들을 한번씩 보자
  

## DAY2

- GO 장점
  - 간결한 문법 및 단순함(코드정리 지원)
  - 병행 프로그래밍 지원
    - go의 핵심 장점
  - 정적 타입 및 동적 실행
  - 간편한 협업 지원
  - 컴파일 및 실행속도가 빠름
  - 제네릭 및 예외처리 미지원
    - 장점이자 단점이 될 수 있음

- 변수
  - var
    - 숫자 첫글자x, 대소문자 구분o, 문자, 숫자, 밑줄, 특수기호 사용 가능
    - 선언했는데 사용하지 않으면 애러가 나옴
    - var a string = "Test1"
    - var b = "Test2" // go 내부적으로 자동 컨버팅
    - 선언을 했는데, 초기화 하지 않으면 기본값 int = 0 string = null
  - 짧은선언
    - ex) shortVar1 := 3 
    - 함수 안에서는 사용 (전역x) 1회성 사용. - 메소드가 끝나면 메모리에서 지워진다.
    - 제한된 범위의 함수 내에서 사용 할 경우 코드 가독성↑
  - const
    - 사용 초기화, 한 번 선언 후 값 변경 금지, 고정 된 값 관리 용
  - iota
    - 일정한 규칙에 따라 숫자를 계산 및 증가시키는 묶음
    - 규칙적으로 증가. 중간값 생략가능. 규칙성은 유지.


## DAY1

- 기본다운로드 
  - https://go.dev/, https://atom.io/

- ATOM Package 설정
  - go-plus : https://atom.io/packages/go-plus
  - terminal : https://atom.io/packages/platformio-ide-terminal
  - script : https://atom.io/packages/script/

- 기본폴더구조
  - bin
    - 소스코드를 컴파일해서, 실행가능한 파일을 만들면, bin폴더에 저장이 자동으로 된다.
  - pkg
    - 외부 패키지 라이브러리
  - src
    - 소스파일

- 기본명령어
  - go run .go파일(ctrl + shift + b)
    - 빌드하지 않고 바로시작(단위테스트 느낌)

  - go build .go파일
    - 실행가능한 파일을 만드는것(테스트개념)

  - go install
    - bin폴더를 가면 session1.exe파일이 만들어짐, go는 폴더이름을 프로젝트이름으로 인식, 개발이 완료됬다고 판단 배포하는 명령어 느낌(패키지에 만약 라이브러리가 있으면, 의존관계에 있는 것 까지 다 빌드후에 인스톨을 해줌)
