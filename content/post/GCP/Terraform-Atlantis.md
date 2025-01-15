---
title: "Atlantis의 동작 방식에 대해 알아보자"
date: 2024-12-27T19:06:20+09:00
draft: false
categories :
- Infra
---

![image](https://github.com/user-attachments/assets/a9d390b9-e631-46fb-b548-f51141ffacb3)

이번 글에서는 아틀란티스의 동작 방식에 대해 공부해보고자 한다.

## Atlantis
### Atlantis는 왜 생겨났을까?
테라폼을 사용하면 인프라를 코드로 관리할 수 있지만, 다음과 같은 문제가 발생한다.

1. 버전 관리의 어려움
- 로컬에서만 테라폼을 사용할 경우, 누가 어떤 변경을 했는지 추적이 어려움
- 백엔드로 GCS를 사용해도 변경 이력 관리가 제한적임
- 코드 버전 관리가 체계적으로 이루어지지 않음

2. 협업과 리뷰의 한계
- 변경사항을 적용하기 전 검토가 어려움

3. Lock 관리의 복잡성
- 분산 환경에서 여러 사용자가 동시에 변경할 때 충돌이 발생할 수 있음
- 워크스페이스별 Lock 관리가 번거롭고 실수 발생할 수 있음


이러한 문제들을 해결하기 위해 Atlantis가 등장했다. 아래에서 Atlantis에 대해 더 자세히 알아보자.

### Atlantis란?
먼저 아틀란티스는 테라폼의 Pull Reqeust를 자동으로 관리할 수 있는 오픈 소스 도구이다. 즉 Terraform Pull Request Automation이다.
Github과 통합되어 코드 변경 사항을 적용하고 검토할 수 있도록 도와준다.

테라폼을 통해 인프라를 코드로 관리할 수 있는데 이를 다시 Github과 통합할 수 있는 아틀란티스트를 사용하여 변경 사항을 관리하고 변경 사항을 적용하기 전에 검토할 수 있는 기능을 제공한다.

즉, 테라폼은 인프라를 코드로 관리할 수 있게 해주며 아틀란티스는 이러한 변경 사항을 팀원들과 협업하여 효율적으로 관리할 수 있도록 도와준다.

### Atlantis의 주요 기능
1. **중앙집중식 관리**
   - 모든 테라폼 변경사항을 Git 기반으로 관리
   - 변경 이력과 작업자 추적이 용이
   - 통합된 로그 관리 가능

2. **효율적인 협업 프로세스**
   - PR 기반의 리뷰 프로세스 제공
   - 자동화된 plan 실행으로 변경사항 미리보기
   - 팀 협업을 위한 표준화된 워크플로우 제공

3. **안전한 Lock 관리**
   - 중앙화된 Lock 관리로 충돌 방지
   - 자동화된 Lock 획득/해제
   - 작업 상태의 실시간 확인 가능

4. **확장성**
   - 웹훅 기반으로 커스터마이징 용이
   - 다양한 도구들과의 통합 지원 (예: 테라그런트, 코스트 관리 도구 등)

더 자세한 내용은 아틀란티스의 동작 방식에 대해 공부하며 알아보자.

## Atlantis 동작 방식
### 1. 웹훅 트리거
<img width="1019" alt="image" src="https://github.com/user-attachments/assets/42dd1d82-6bdf-413b-ad2b-40bc669f9087" />

PR 생성 시 아틀란티스가 동작할 수 있게 하기 위해서는 사진과 같이 Github의 세팅에서 웹훅을 등록해주고, 트리거를 설정해야 한다.

### 2. PR 생성
개발자가 테라폼을 수정한 뒤 Github으로 PR을 생성한다.

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