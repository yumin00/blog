---
title: "GopherCon"
date: 2023-08-05T10:12:26+09:00
draft: true
---

# GopherCon Korea Day 1

## 1. Go 테스트
(문서&코드) jmote/go-test

### Go 테스트 기초
패지지명 : xxx_test.go

함수명 : TestXxx

```go
func TestXxx(t *testing.T) {
}w
```

명령어 : go test

testify
- testing 보다 간결
- asser/require/mock/suite
```go
func TestAbs(t *testing.T) {
	assert.Equal(t, 1, ..)
}
```

- 화이트박스 테스트 : 기본
- 블랙박스 테스트 : import한 디렉토리 테스트


서브테스트 : t.Run()을 이용한 테스트 내의 테스트
table-driven : t.Run()

테스트의 유형
- test
  - pass/fail
- benchmark
  - 성능 측정
  - benchtime 
- fuzz
  - 자동 생성된 랜덤값 테스트
  - 쿼리문 검증에 활용 가능
  - *testing.F
- example
  - 원하는 대로 작동하는지 테스트, 출력값 검증(test과 비슷하지만 제한적임)

error : 다 실행
fatal : 에러가 있으면 바로 중단
asser : 다 실행
require : 에러가 있으면 바로 중단

Failfast
- 테스트 함수 중 하나라도 fail 이면 중단

Skip
- 오래 걸리는 t.shkp으로 건너뛸 수 있음
- short와 함께 사용

Parallel
- 함수 순차 실행

### 커버리지 100%

## 2. Golang 도입
###
- 해결하기 어려운 상황
  - ralils 서버로 감당하기 어려운 트래픽 발생

### go 장단점
- 장점
  - 컴파일 속도가 빠르고 효율적
  - 성능 튜닝
- 단점
  - 에러 처리 번거로움

### 성능 이슈
- starvation mode of mutex
  - 동시성 제어가 안됨
  - mutex mdoe
    - normal mode
- container 환경에서 cpu throttling
  - uber ~
- 불필요한 메모리 할당 최소화
  - make로 초기크기와 최대크키 설정 가능
  - 할당된 메미로 재사용 : 
- gc tuning
  - gc freqeuncy가 높을 때
    - daangn/autopprof
- fluentbit plugin

## 3. 컨텍스트를 이용한 상태 관리
### context

context란 무엇인가? 왜 사용하는가?

- context.Background()
  - 프로그램의 메인 함수에서 생성해서 활용해야 됨
  - 모든 ctx의 부모 역할
- context.TODO()

- context.WithCancel()
  - context 취소
  - 취소 시그널 전파 가능
- context.WithDeadline()
  - 데드라인이 되면 타임아웃
- context.WithValue()

### usage
[작업 영역에 대한 정의 및 제한 설정]


### conclusion


## 4. API 서버 테스트코드
### Golang 테스트 개요


### 테스트를 어렵게 만드는 요소(유지보수하기 어려운 코드)
- 복잡한 매개변수와 반환 값
- 고루틴 & 채널
- 외부 의존성 ; db, 외부 api
- 전역 상태 관리
- 코드 구조와 설계 : Go Clean Architecture 참고하는 것을 추천
  - clean architecure 원칙
    - 독립성
    - ui, 데이터베이스 등 외부요소에 상관없이 테스트

### Layer 별 테스트 Usecase
- docker
- testmain
- in memory database

### 어려운 테스트 Usecase
- go-cmp
  - 테스트를 무시하고 작성
- 어쩔 수 없이 비대해지는 함수
  - 내부 함수 mocking
- 고루틴 & 채널
  - 채널 반환

거의 모든 코드에 대해서 테스트 코드를 짜는 편인가요?

원래 개발 사이클에

유지보수 분석 개발 테스트 유지보수
이 사이클이라서
테스트를 안한 코드는 사실 내보내면 안됩니다
그래서 TDD BDD 같은 방식으로 테스트의 중요성을 매우 강조하는 편입니다

저런 테스트 코드를 사용하면, 모든 경우의 수를 다 분석해주는건가요?

다 자기가 짠 테스트 케이스를 넣는거고 거기에 대해서 테스트 커버리지를 측정하는 편입니다

저는 코드가 잘 작동하는지 확인할 때 포스트맨으로 하는편인데 이거는 테스트라고 볼 수 없는건가요?

볼 순 있는데, 보고서 뽑는다던지 테스트 케이스 정리한다던지 이런거에서 테스트 커버리지 작성이 거의 불가능한 편이라서 차라리 코드로 작성하여
해당 커버리지를 한번에 뽑아보는 편이 낫고, 협업시에 해당 테스트 케이슫 ㅗ공유할 필요가 있어서 앵간하면 코드로 다 표혀낳는게
고랭의 패러다임이며, 특징인거같습ㄴ디ㅏ

그럼 모든 파일에 대해서 test 파일을 만들고 파일에 있는 함수에 대한 테스트 코드를 모두 작성하나요?

aasdsd_test.go 에서 해당 들어있는 함수들에 대해 적당히 다 테스트해야합니다. 예를 들어 연결 이되었는가? 메소드가 실행이되서 디비에 접근했는가?

## 5. Go와 K8S로 만드는 Datacenter 자동화
### kubernetes operator pattern
- yaml 로 되어있는 server custom resource

인프라 자동화 같은거는 구글 클라우드에서 제공하고 있나요?

구글 클라우드 강튼 퍼블릭 ㅡㅋㄹ라우드들은 전부 그냥 상품만 제공할 뿐이지 모든 자동화는 테라폼이라는 다른 서드파티 프로그램으로 사용해야하고
그걸로 제가 다 짜서 사용하고 있는데 정확히 모든 자동화라고 보기는 어렵고 반자동화로 보면 될 것 같스빈다
모듈 단위로 전부 코드로 작성하고 그 코드들을 배포 및 삭제를 보내는 형식이고,
조금 더 자동화하기 위해서 깃헙이랑 붙여서 깃헙에서 뭔가 트리거를 주면 인프라르 만들고 부시는 작업을 달아놓긴 했으나, 정확히 자동화라고 보기 어려운 상황입니다

## 6. OLAP 데이터베이스 개발
### 성능 최적화 대상 쿼리
- 적당히 빠른 쿼리란?
  - p90에서 p95 사이의 쿼리

### pprof
- 프로파일링
  - 프로그램 실행시간, 메모리 사용량, 함수 호출 빈도 측정
  - 프로그램 최적화를 보조하기 위해 사용
- pprof
  - 프로파일링 데이터 분석 시각화

- pprof package

- trace

- gotraceui


자기네 디비 서비스하는데 쿼리가 느린거 빠른거 중에서 적당히 빠른 것들을 최적화 시키고싶다.
근데 다 고랭으로 되어있어서 어떻게할까 생각하다가
우린 이 방법을 사용해서 추적해서 뭐를 최적화 했다에서
이 방법! 에 대해서 설명중입니다.

데이터 분석같이 실시간ㄴ으로 들어오는 데이터들을 저장하고 파싱해서 사용하기 위한 데이터베이스같아요
거기서 \필요한 데이터 /\ㅋ뽑을려고 쿼리날리는데 그게 많이 이러지고이ㅆ다 그래서 느리먀ㅕㄴ 안되닉ㅏ 최적화한다..

## 7. eBPF 도구를 이용해 Go 어플리케이션 추적하기
### eBPF
운영체제 기능을 확장하는 프로그래밍 언어 / 런타임
- 운영체제 안에서 실행
- 커널 내부에서 안정하게 실행되도록 제약

### eBPF 이벤트 소스
- 동적 추적
  - 내가 원하는 함수를 추적
- 정적 추적
  - 
- kprobe, kretprobe
  - 커널 코드 임의의 위치에 동적으로 이벤트 설정
  - 특정 커널 함수의 진입, 반환 추적
- uprobe, uretprobe
  - 사용자 공간 프로그램 임의의 위치에 동적으로 이벤트 설정
  - 특정 사용자 공간 함수의 진입, 반환 추적
- tracepoint
  - 커널 코드 특정 위치에 정적으로 설정된 이벤트
  - 커널 개발자들이 중요하다가 판단하는 위치에 설정
- usdt
  - 사용자 공간 프로그램에 정적으로 미리 설정된 이벤트
  - 사용자 개발자들이 ~
  - 활성화/비활성화 가능

### bpftrace

### Go 어플리케이션 추적 제약 사항


# GopherCon Korea Day 2
## 1. Golang으로 서버 모니터링 툴 개발
### 모니터링이란?
- 자원의 상태를 관찰
- 서버 : cpu 사용량, memory 사용량, disk I/O

### 왜 golang으로 만드는가
- cross-compile 가능
- binary 파일로 떨어지기 때문에 가벼움
- 비동기 프로그램

### 메모리 모니터링
- influxDB UI

## 2. AWS Lambda in Go (with Kafka)
### 기존 아키텍쳐 개선점
- 서버리스
- 큐
  - 하나의 이벤트를 여러 서비스에서, 여러 차례 재사용

### 새로운 아키텍처
- API GateWay
  - 가장 앞단에서 슬랙 요청을 받음
- AWS Lambda
  - 최소한의 전처리 후, msk로 이벤트 전달
- Amaxon MSK
  - 이벤트 consume 후 전달

## 3. 프로메테우스는 어떻게 쿠보네티스와 메트릭을 자동으로 가지고 올까요?
### GCP
- vertex
- google bigquery
- google kubernetes engine

### 프로메테우스
### 쿠버네티스

## 4. 버그 없는 프로그램 만들기 - qmffhrcpdls
### Unit Testing
- 관련된 모듈, api를 모두 함께 테스트를 진행함
- Test suite 사용
  - testify library 일부

### Mutation Testing
- test를 돌릴 떄 test가 자체적으로 프로그램을 작은 방법으로 수정(synthetic change)하는 방법
- 각 변형 버전을 mutant라고 하며, mutant 상태일때에 fail하는 테스트가 있는지 확인 -> fail 하는 테스트가 없다면?
- Test하는 코드도 test를 해야하는 것 아닌가?
- mutation testing : test 코드를 test하는 코드
- statuemnet mutation, value mutation, decision mutation
- go-mutesting

### E2E Testing
- 실제 사용자 경험을 시뮬레이션 하며 테스팅 하는 방법
- QA와도 밀접한 관걔가 있음

### Fuzz Testing
- 랜덤한 입력값을 메서드나 프로그램에 주입하였을 때 출력값을 검증하고 테스트하는 테스팅 방법
- edge case 테스트 가능

## 5. 시나리오 테스트
### 인수테스트
문제
- 매번 짜야하는 코드가 많음
- 어던 인수테스트가 있는지 한 눈에 보기 어려움
- 언어 변겨을 학 ㅔ된다면 사람이 손으로 옮겨줘야 함

### 시나리오 테스트


## 6. 서버 레이턴시 개선
- goroutine
  - 트래픽이 몰려도 안정적으로 서빙 가능

- fiber
  - net/http와 호환되지 않지만, 프레임워크 기능을 많이 필요로 하지 않음
  - zero aloocation
    - 각 요청마다 사용되는 fiver.Ctx를 pooling
    - http header, query parameter와 같은 값들을 버퍼로부터 복사 없이 사용할 수 있도록 제공
    - heap 할당을 줄여줌

## 7. GC in Golang
### GC란?
- 메모리 관리 / 메모리 할당 해제에 초점

### concurrent mark and sweep (cms)
- mark : 쓰이는 객체, 쓰이지 않는 객체는 찾아 마크
- sweep : 모든 힙을 돌아다니면서 쓰이지 않는 힙은 할당 해제

### mark - tri color abstraction
- black : 접근할 수 있는 객체, 지워져서는 안됨
- gray : root로부터 접근 가능하지만 아직 모든 자신 객체들이 mark되지 않아서 mark 단계가 끝나지 않았음을 암시
- white : mark되지 않은 객체. sweep 단계에서 white 객체들은 해제
- mark 가 끝나면 grey 객체는 없음
- atomic operation
  - 하나의 operation만 실행해야 됨

### Write Barrier
- mark가 끝났는데 화이트 객체에 참조했을 경우,
- wirte barrier는 객체가 참조된 순간 grey로 만듦
- 잘못 할당 해제하는 상황을 방지함
- 문제
  - grey 객체가 계속 생성됨
  - 이를 막기 위해 mark 단계를 2개로 나누고 STW를 실행함
- black이 됐는데, 참조가 없어진 경우 : out of memeory 상황이지만, 잘못된 객체를 할당 해제한 게 아니기 때문에 괜찮다고 생각

### Memory Allocation 메모리 할당
- 메모리를 어떻게 잘 관리할 것인가?
- GC : non-coping, non-moving - 복사하지 않고, 움직이지 않는다.
-> 발생하는 문제 : 메모리 단편화(memory fragmentation)
- 메모리 단편화 해결 방법 : TCMalloc
  - thread cache를 사용
  - 객체 할당을 thread cache에

- P, M, G
  - P : manager resources for scheduler
  - M : os leel thread
  - G : 고루틴
  - mcache 사용

- 작은 메모리 : 아주 작은 객체(tiny object)는 따로 관리
- small object : 대부분의 객체는 각 개체마다 class를 정해서 해당 class에 할당 - 비슷한 크기의 객체드르이 비슷한 메모리에 위치
- large object : 큰 객체는 heap 에서 직접 곤리
- mcache span이 모두 꽉차있다면, 중앙 힙에서 새로운 span을 할당 받음


- mark bit / freelist 
- bitmap

### Triggering GC
- tradeoff
- gc를 자주 돌리면 메모리 양은 많아지겠지만, cpu 사용량이 너무 많아짐
- gc를 얼마나 자주 실행할 것인지 설정할 수 있게 함
  - GOGC & Soft Memory Limit

### performance
- background worker
- user goroutine이 gc mark를 도와줌
- cpu를 대략 30% 정도 씀

### preemption for STW
- tight loop problem : 한 고루틴이 for 문을 엄청 많이 돌 때
  - 다른 고루틴이 이 고루틴이 끝날 때까지 기다려야함
    - 해결 방법 1 : runtime.Goshed() : 다른 고루틴에게 스케줄 양보
    - 해결 방법 2 : cooperative preemption

cooperative preemption


- non-cooperative async preemption
  - signals을 통해 구현됨 - goroutine에서 signal을 알려줌
  - SIGURG
  - Safepoint : 리소스스를 념거서 다른 고루틴이 돌아도 문제가 없음을 알려주는 것

### implementations
- GC 시작 조건 호가인
  - 함수 시작 부분에 acquirem() 함수를 호출한 후 발생함
- GC 모드 설정
- STW 이후 스윕 단계 종료
- 스케줄러 조정
- 마킹 단계 시작
- stw 해제하고, Mark 진행
- gc 시작의 완료 후 정리
  - 시작 완료와 이후 정리 작업이 이루어짐
  - 유저의 고루틴이 다시 시작되고 필요한 세마포어가 해제됨
- worker들이 전부 mark를 실행했고, mark 단계를 끝낸다
- sweep 진행