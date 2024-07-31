---
title: "Kafkad의 Replication에 대해 알아보자"
date: 2024-03-29T15:15:30+09:00
draft: true
categories :
- DataEngineering
- Kafka
---

브로커에 장애가 발생하면 어떻게 될까?
프로듀서, 컨슈마가 데이터를 프로듀서하거나 가져갈 수 없게 됨.

파티션 관려 데이터: message offset
log-end-offset과 current-offset의 차이 consumer log가 발생할 수 있음.

다른 브로커에서 파티션을 만들면 장애를 해결할 수 ㅣㅇ씅ㄹ까?
장애가 발생한 브로커를 대신해서 다른 브로커에서 동일한 ㅌ파티션을 만들어서 해결?
-> 메시지를 보내거나 저장할 수는 있음 / 그러면 메시지, 오프셋 정보가 날아감 

replication 기능 제공: 파티션을 미리 복제해서 장애를 대비하는 기술
실제 작동하는 파티션; 리더
복제한 파티션: 팔로워

replication factor를 설정할 수 있음 (leader + follower partition의 개수)

replicas = leader partion + follower partition 

실제로 프로듀서와 컨슈머는 리더와만 통신을 함. 팔로워는 복제만 함.
팔로워는 브로터 장애시 안정성을 제공하기 위히 제공 / commig log를 데이터를 가져오기 요청(Fetch Request)으로 복제


2.4버전부터, 팔로워 파티션에서 데이터를 리드할 수 있는 기능을 제공함 (컨슈머)

리더가 장애가 나면 어떻게 될까?
카프카 클러스터가 팔로워 중 새로운 리더를 선출함. 클라이언트는 자동으로 새 리더로 전환됨. 나머지 팔로워는 그대로 리드해서 복제해감

파티션 리더에 대한 자동 분산
하나의 브로커에만 파티션의 리더가 몰려있다면 어떻게 될까? : 특정 브로커에만 클라이언트들이 잔뜩 붙여서 부하가 집중되어 핫스팟이 됨. 문제가 발생할 여지가 높아짐

이를 방지하기 위해 파티션 리더에 대한 자동 분산의 옵션을 제공하고 있음

부하를 분산시키고, 모든 하드웨어 장비에 리소스를 골고루 제공 안정적인 서비스와 빠른 레이턴시를 보장해줌

auto.leader.rebalance.enable: 기본값은 enable - default로 켜져있음
leader.imbalance.check.inverval.seconds: 300초마다 리더의 불균형 상태를 체크
leader.imbalacne.per.broker.percentage; 기본값 10 / 브로커들의 불균형 상태가 10%가 넘으면 다시 리밸런시을 하겠다.

rack awareness
브로커는 분산 시스템으로 구성되어 있음. 브로커들이 올라가는 장비들을 하나의 rack에만 넣고 실행시키면, rakc에 공급되던 전원이 나가면 (만에 하나) 모든 브로커가 셧다운 됨. 정상적 작동이 안됨.
rack 혹은 availiabel zone 상의 브로커들에 동일한 rack name을 지정할 수 ㅣ있음
rack 간에 균현을 유지하며 rack 장애를 대비해야됨
topic 생성 시에 auto data balacer / self balancing cluster 동작 때만 자동으로 실행됨.


In-Sync Replicas (ISR)
리더가 장애가 나면 새로운 리더를 선출하는데 사용하는 것
high water mark라고 하는 지점까지 동일한 replicas의 목록

fully-replicated committed: replica.lag.max.message 이하만큼 따라잡은 위치 ==  high water mark
ISR: 리더의 메시지를 잘 따라잡고 있는 팔로워
Out of sync follwer: 리더의 메시지를 replica.lag.max.message 이상 차이나는 팔로워


리더에 장애가 발생하면 isr 중에 새로운 리더를 선출함.

replica.lag.max.message 사용시 문제점
메시지가 항상 일정한 비율로 들어올 때는 ISR이 정상적으로 동작한다.
만약, 메시지 유입량이 갑자기 늘어나서 팔로워들이 리더의 메시지를 따라자기 못하면, OSR로 변경된다.
실제로 팔로워들은 정상적으로 작동하고 잠깐 지연이 발생한 것인데, OSR로 판단되어 에러 혹은 불필요한 리트라이를 유발할 수 있다.

팔로워가 리더에게 fetch 요청을보내는 interval을 체크하는 replica.log.time.max.ms로 판단해야한다.
따라서 카프카에서는 replica.lag.time.max.ms 옵션만 제공하고 있다.

ISR은 리더가 관리한다.
리더가 팔로워 중에 느린 팔로워가 있으면 isr에서 해당 팔로워를 제거하고 주키퍼에게 이를 알린다. 그러면 주키퍼는 파티션 메타데이터에 대한 변경사항을  contorller에게 알린다.(컨트롤러는 브로커 중 하나)

controller란?
주키퍼를 통해 브로커가 살아있는지 죽어있는지 모니터링
리더와 레플리카 정보를 클러스터 내 다른 브로커들에게 전달해줌
주키퍼에서 replicas 정보의 복사본을 유지한 다음 더 빠른 액세스를 위해 클러스터의 모든 브로커들에게 동일한 정보를 캐시함.
리더 election (어떤 브로커를 선택할건지 결정)
컨트롤러가 장애가 나면, 다른 active broekr중에 재선출함.(이는 주키퍼가 결정)


position
last committed offset, curren position, high water mark, log end offset

- last committed offset(current offset): 컨슈머가 최종으로 커밋한 오프셋
- curren position:consumer가 읽어간 위치(커밋 하기 전)
- high water mark: isr간에 복제된 오프셋
- logendoffset:프로듀서가 메시지를 보내서 저장된 로그의 맨 끝 오프셋

committed: isr 목록의 모든 replicas가 메시지를 성공적으로 가져오면 committed
consumer는 committed 메시지만 읽을 수 있음
committed 메시지는 동일한 오프셋
브로커가 다시 시작할 때 committed 메시지 목록을 유지하기 위해서 브로커의 모든 파티션에 대한 마지막 committed offset은 replication-offset-checkpoint라는 파일에 기록됨


replicas 동기화 과정
리더 팔로워가 가장 마지막으로 committed한 메시지의 오프셋을 추적하고, 해당 오프셋을 replication-offset-checkpoint 파일에 기록한다.
새 리더가 선출된 시점을 오프셋으로 표시
브로커 복구중에 메시지를 체크포인트로 자른 다음 현재 리더를 따르기 위해 사용디ㅗㅁ
컨트롤러가 새 리더를 선택하면 leader epoch를 업데이트하고 해당 정보를 isr 목록의 모든 구성원에게 전달
leader-epoch-checkpoint 파일에 체크포인트를 기록

message commit 과정
팔로우어ㅔ서 리드로 패치만 수행

만약 3개의 레플리카들이 offset 5번까지 복제가 완료되어 있다고 가정해보자

1. 프로듀서가 메시지를 전송하면, 리더가 offset 6에 메시지를 추가 
2. 팔로워들이 fetcher thread가 독립적으로 fetch를 수행하고, 가져온 메시지를 Offset 6번에 write
   (자바 스레드)
3. 팔ㄹ로워들은 계속 패치하고 null을 받음. 리더는 팔로워들이 잘 패치했음을 파악하고 high water mark를 6으로 이동시킴
4. 또 패치하고 하이 워터 마크를 받아서 자신의 오프셋에도 하이워터마크를 이동시킴.

replicas 리스트 관리
commit이 무스 ㄴ의미? -> 팔ㄹ오워들이 가져가서 ack도 보낸 상태.

replica recovery
m3는 커밋이 된 적이 없는데 어떻게 m3를 가지고 있지? -> 한번 들어온 데이터를 지우지 않음.

죽었다가 살아나면 데이터 지우고, 리더와 패치함.

가용성과 내구성 중 선택
3개의 레플리카, 에크 올, 민 2 -> 3-2 :1 한대의 장애만 허용할 수 있다.

