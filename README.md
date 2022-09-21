## 💡 Introduction
- GO STUDY

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
