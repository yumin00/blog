---
title: "Atlantis의 동작 방식에 대해 알아보자"
date: 2024-12-27T19:06:20+09:00
draft: true
categories :
- Infra
---

![image](https://github.com/user-attachments/assets/a9d390b9-e631-46fb-b548-f51141ffacb3)

이번 글에서는 아틀란티스의 동작 방식에 대해 공부해보고자 한다.

## Atlantis
먼저 아틀란티스는 테라폼의 Pull Reqeust를 자동으로 관리할 수 있는 오픈 소스 도구이다. Github과 통합되어 코드 변경 사항을 적용하고 검토할 수 있도록 도와준다.

테라폼을 통해 인프라를 코드로 관리할 수 있는데 이를 다시 Github과 통합할 수 있는 아틀란티스트를 사용하여 변경 사항을 관리하고 변경 사항을 적용하기 전에 검토할 수 있는 기능을 제공한다.

즉, 테라폼은 인프라를 코드로 관리할 수 있게 해주며, 아틀란티스는 이러한 변경 사항을 팀원들과 협업하여 효율적으로 관리할 수 있도록 도와준다.

더 자세한 내용은 아틀란티스의 동작 방식에 대해 공부하며 알아보자.

## Atlantis 동작 방식
### 1. PR 생성
개발자가 테라폼 혹은 테라그런트를 수정한 뒤 Github으로 PR을 생성한다.


### 2. 웹훅 트리거
<img width="1019" alt="image" src="https://github.com/user-attachments/assets/42dd1d82-6bdf-413b-ad2b-40bc669f9087" />

참고로, PR 생성 시 아틀란티스가 동작할 수 있게 하기 위해서는 사진과 같이 Github의 세팅에서 웹훅을 등록해주고, 트리거를 설정해야 한다!

PR이 생성되면 Github은 설정되어 있는 아틀란티스 웹훅 URL로 PR의 변경사항을 담아 API를 호출한다.

### 3. 아틀란티스 동작
#### Locking
아틀란티스가 동작할 때 특징에는 Locking이 있다. 아틀란티스는 PR을 기반으로 동작하기 때문에 여러 사람이 동시에 작업하게 된다면 conflict가 발생할 수 있다.

아틀란티스는 동시에 여러 PR에서 변경 사항이 적용되어 충돌이 일어나는 것을 방지하기 위해 Locking을 제공한다. A 레포에서 a-branch, b-branch가 있다고 가정해보자. a-branch에서 plan이 이루어진다면 b-branch에서는 plan-apply를 할 수 없다. 

#### Terraform 실행 (Autoplanning)
<img width="947" alt="image" src="https://github.com/user-attachments/assets/020af58d-20a7-4dff-b539-d098e52a987b" />

아틀란티스는 plan 명령을 자동으로 실행하여 현재 상태와 변경 사항을 비교하여 리소스의 추가/수정/삭제를 결정하여 이를 사용자에게 보여준다.

Show Ouput을 통해서 테라폼의 변경 사항을 확인할 수 있다.

#### 변경 사항 적용
해당 사항을 적용하기 위해 사용자는 `attlantis apply` 명령을 실행하면 아틀란티스는 이를 수신하여 테라폼을 통해 변경 사항을 인프라에 적용한다.


#### Atuomerging
<img width="1093" alt="image" src="https://github.com/user-attachments/assets/93ac74bb-b2a3-42f3-bfbb-c450562cd963" />
아틀란티스는 automerging 기능을 통해 모두 성공적으로 적용되면 자동으로 PR을 merger할 수 있다.