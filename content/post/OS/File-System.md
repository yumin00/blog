---
title: "파일 시스템에 대해 알아보자"
date: 2024-07-10T19:39:48+09:00
draft: false
categories :
- OS
---

파일 시스템이란, 데이터를 효율적으로 저장하고 관리하기 위해 사용되는 구조와 규칙이다. 즉, 파일을 저장할 때 데이터를 어떻게 읽고 쓸 것인지에 대해 미리 규칙을 정한 것이 바로 파일 시스템이다.

파일 시스템은 파일과 디렉터리를 관리하는데, 파일 시스템에 대해 더 자세히 알아보기 전, 파일과 디렉터리에 대해 공부해보고자 한다.

# 파일과 디렉터리
## 파일
### 파일이란?
파일은 하드 디스크나 SSD와 같은 보조기억장치에 저장되어 있는 관련 데이터의 집합이다. 문서, 이미지, 비디오, 프로그램 등 다양한 형태의 데이터를 파일로 저장할 수 있다. 파일은 일반적으로 `파일.txt`와 같은 형태로 확장자를 통해 나타낸다.

### 파일 구성요소
파일은 다음과 같은 정보를 포함한다.
- 이름
- 파일을 실행하기 위한 정보
- 파일 관련 부가 정보(== 속성 == 메타데이터)

여기에서 파일 속성이란, 아래 사진처럼 파일에 대한 부가적인 정보들을 말한다. 

![image](https://github.com/yumin00/blog/assets/130362583/5581d8f6-93aa-4140-9e25-c1c632e3cd64)

### 시스템 호출
파일을 생성하거나 삭제하는 것과 같은 파일을 다루는 모든 작업을 운영체제에 의해 이루어진다. 즉, 운영체제에서 제공하는 시스템 호출을 통해 파일을 다루는 작업을 진행할 수 있다.

## 디렉터리(폴더)
### 디렉터리란?
디렉터리는 파일과 다른 디렉터리를 포함할 수 있는 컨테이너이다. 파일 시스템의 계층 구조를 형성한다.

디렉터리도 파일이다. 파일이 파일과 관련된 정보를 포함하고 있다면, 디렉터리는 디렉터리에 담겨져 있는 대상과 관련된 정보를 담고 있다.

디렉터리를 저장할 때, 디렉터리 엔트리 라는 자료구조를 사용한다. 파일 시스템을 어떤 것을 사용하냐에 따라 디렉터리 엔트리에 저장하는 정보는 달라지지만 일반적으로 테이블 형식으로 다음과 같이 디렉터리에 담겨져 있는 대상과 관련된 정보를 가진 테이블이 보조기억장치에 저장된다.

![image](https://github.com/yumin00/blog/assets/130362583/a1d23776-cd58-4a6f-968f-42ee5cfcc725)

### 디렉터리 구조
#### 1단계 디렉터리 (single level directory)
![image](https://github.com/yumin00/blog/assets/130362583/377f64dc-53a9-4d41-9550-c806b338beba)
모든 파일이 하나의 디렉터리 안에 존재하는 구조를 1단계 디렉터리라고 한다.

#### 트리 구조 디렉터리 (tree structured directory)
![image](https://github.com/yumin00/blog/assets/130362583/c123571b-7bfb-404c-b722-549cea657352)
트리 구조 디렉터리는 최상위 디렉터리(루트 디렉터리)가 있고, 그 아래 여러 서브 디렉터리가 있고, 그 아래 또 다른 서브 디렉터리를 가질 수 있는 구조이다.

이렇게 여러 디렉터리를 포함하다보니 생긴 개념이 바로 경로(path)이다.

### 시스템 호출
디렉터리를 다루는 작업도 파일과 마찬가지로 운영체제에서 제공하는 시스템 호출을 통해 이루어진다.

# 파일 시스템
파일 시스템은 파일과 디렉터리를 보조기억장치에 저장하고, 접근할 수 있게 하는 운영체제 내부 프로그램이다. 파일 시스템은 다양한 종류가 있고, 하나의 컴퓨터에서 여러 파일 시스템을 사용할 수 있다.

## 파일 저장
파일 시스템이 파일과 디렉터리를 관리할 때, 파티셔닝과 포매팅을 사용하여 저장한다.

### 파티셔닝
파티셔닝이란, 저장 장치를 여러 구역으로 나누어 파일과 디렉터리를 쉽게 저장할 수 있도록 하는 방법이다. 파티셔닝을 통해 나누어진 영역 하나하나를 파티션이라고 한다.

### 포매팅
포매팅이란 파일 시스템을 어떤 방식으로 사용할 것인지 결정하는 방법이다. 즉, 파일과 디렉터리를 어떻게 저장하고 어떻게 관리할 것인지 결정하고, 새로운 데이터를 쓸 준비를 작업하는 것이다. 그래서 포매팅을 진행하면 파일 시스템의 방법이 새로 결정되며 저장 장치를 완전히 삭제한다.

저장 장치의 파티션 별로 서로 다른 파일 시스템을 사용할 수 있다.

## 파일 할당 방법
![image](https://github.com/yumin00/blog/assets/130362583/107d80b5-2436-44b9-9ec8-04e199e5d20d)
운영체제는 파일과 디렉터리를 블록 단위로 읽고 쓴다. 즉, 파일 하나를 보조기억장치에 저장할 때 하나 이상의 블록에 걸쳐 저장된다. 파일을 보조기억장치에 할당할 때 연속 할당/불연속 할당(연결 할당, 색인 할당)을 사용할 수 있다.

### 연속 할당 (contiguous allocation)
![image](https://github.com/yumin00/blog/assets/130362583/d85d2530-326c-4cc4-9003-c5174e50be9a)
연속 할당은 말 그대로 보조기억장치 내의 연속적인 블록에 파일을 할당하는 방법이다.

연속으로 할당된 파일에 접근하기 위해서는 파일의 첫 번째 블록 주소와 블록 단위 길이만 알면 된다. 그렇기 대문에 연속 할당을 사용하는 파일 시스템에서는 디렉터리 엔트리에 파일 이름과 첫 번째 블록 주소와 블록 단위의 길이를 명시한다.

![image](https://github.com/yumin00/blog/assets/130362583/8ef000ac-aee5-4815-932e-772bda05de34)

[장점]
- 연속적으로 저장하기 때문에 구현이 단순함

[단점]
- 외부 단편화를 야기함

예를 들어, 분홍 파일이 삭제되었을 때 분홍 파일을 삭제된 블록 3부터 7까지는 이 안에 할당이 가능한 파일이 있어야만 할당될 수 있다. 이후에 길이가 6 이상인 파일만 저장된다면 해당 공간은 영원히 할당되지 못한다.

![image](https://github.com/yumin00/blog/assets/130362583/07fe0c4c-6f5d-4c74-8ad5-5d55c3fd56e6)

### 연결 할당 (linked allocation)
연결 할당이 바로 연속 할당의 외부 단편화를 해결할 수 있는 할당 방법이다.

연결 할당은 각 블록의 일부에 다음 블록의 주소를 저장하여 각 블록이 다음 블록을 가리키는 형태로 할당하는 방식이다.

![image](https://github.com/yumin00/blog/assets/130362583/067d3b52-a31d-4edd-a0ee-c91ad6f495c1)

이때 -1은 다음 블록이 없다는 표시자이다!

해당 방식에서도 디렉터리 엔트리에는 첫 번째 블록 주소와 길이만 있어도 어떤 파일이 어디에 저장되어 있는지 알 수 있다.

[장점]
- 외부 단편화 문제를 해결할 수 있다.

[단점]
- 반드시 첫 번째 블록부터 하나씩 차례대로 읽어야 한다.
  - 파일을 중간 부분부터 읽고 싶어도 다음 블록을 알기 위해서는 이전 블록을 알아야하기 때문에 반드시 첫 번째 블록부터 차례로 읽어야 한다.
  - == 임의 접근 속도가 매우 느리다
- 하드웨어 고장이나 오류 발생시 해당 블록 이후 블록은 접근할 수 없다.
  - 하나의 블록 안에 파일 데이터와 다음 블록 주소가 포함되어 있기 때문에 하나의 블록만이라도 문제가 발생하면 그 이후 블록에 접근할 수 없다.

### 색인 할당 (indexed alloction)
색인 할당은 파일의 모든 블록 주소를 색인 블록(index block)이라고 하는 하나의 블록에 모아 관리하는 방식이다.

![image](https://github.com/yumin00/blog/assets/130362583/db9d19f6-bad9-473b-93a5-1da05bddf32b)

색인 블록에 해당 파일이 할당된 블록의 주소를 차례로 저장해놓는 방식이다.

디렉터리 엔트리에는 파일의 색인 블록 주소만 잇다면, 해당 파일 테이터에 접근할 수 있다.

![image](https://github.com/yumin00/blog/assets/130362583/c70f2c77-6394-4120-b506-7496dab064de)

# 파일 시스템
## FAT 파일 시스템
FAT 파일 시스템은 연결 할당의 단점을 보완한 파일 시스템이다. FAT 파일 시스템은 파일 할당 테이블 (File Allocation Table) 을 사용하는 방식이다.

![image](https://github.com/yumin00/blog/assets/130362583/e7e71476-0cfc-4011-a64d-df8dd2a0ea72)

연결 할당에서 한 블록이 다음 블록의 주소를 가지고 있었다면, FAT 파일 시스템은 모든 블록 주소와 해당 블록의 다음 블록 주소를 데이터가 담겨져 있는 FAT를 사용하는 방식이다.

하드디스크의 한 파티션을 FAT 파일 시스템으로 포맷하면 해당 파티션은 다음과 같은 구조로 구성된다. FAT는 파티션의 앞부분에 만들어지고, 그 뒤에 루트 디렉터리가 저장되는 영역이 있고, 그 뒤에는 서브 디렉터리와 파일들을 위한 영역이 있다.

![image](https://github.com/yumin00/blog/assets/130362583/32e9c666-0c1c-4098-b0bc-de636df970f3)

FAT 파일 시스템에서는 블록을 하나씩 접근하여 다음 블록을 찾아야하는 연결 할당 방식과는 다르게 FAT를 통해 다음 블록의 주소를 확인할 수 있기 때문에 임의 접근의 성능이 개선된다.

옛날 마이크로소프트의 운영체제에서 사용되었으며, 현재까지는 USB 메모리, SD 카드와 같은 저용량 저장 장치용 파일 시스템으로 많이 사용되고 있다.

### 장점
- 단순한 구조로 구현과 이해가 쉽다.
- 대부분의 운영체제와 호환된다.
- 메타데이터와 관리 구조가 간단하여 디슼 공간을 효율적으로 사용할 수 있다.
- 단순한 테이블 구조로 파일의 할당과 해제 작업이 빠르다.

### 단점
- 파일 크기가 제한된다.
  - FAT16: 최대 파일 크기 2GB
  - FAT32: 최대 파일 크기 4GB
- 저널링 기능이 없어서, 갑작스러운 시스템 종료나 오류 발생 시 데이터 손상 가능성이 크다.

> 저널링 기능
>
> 저널링 기능이란, 파일 시스템의 변경 사항을 추적하기 위해 만들어진 기능으로, 삭제된 파일을 복구하는데 사용하기도 한다.

## 유닉스 파일 시스템
유닉스 파일 시스템은 색인 할당 기반을 하는 방식이다. 유닉스 파일 시스템에서는 색인 블록인 i-node(index node)를 사용한다.

각 파일은 i-node를 갖고, i-node에는 파일 속성 정보와 열다섯 개의 블록 주소가 저장될 수 있다.

![image](https://github.com/yumin00/blog/assets/130362583/ac80d7c2-022c-4ac1-89bc-948cb37dac8e)

유닉스 파일 시스템에서는 i-node마다 번호를 부여하고, 파티션 내 특정 영역에 i-node가 모여있다.

![image](https://github.com/yumin00/blog/assets/130362583/c9b3c26c-047c-4434-a647-94dc5b9e3d1c)

i-node에는 15개의 블록 주소를 저장할 수 있다고 했는데, 그렇다면 한 파일이 16개 이상의 블록을 가진 파일이라면 어떻게 저장을 해야할까? 유닉스 파일 시스템은 다음과 같은 방법으로 문제를 해결한다.

### 1. 블록 주소 중 열두 개에는 직접 블록 주소를 저장한다.
i-node가 저장할 수 있는 15개의 블록 주소 중 12개에는 파일 데이터가 저장된 블록 주소를 저장하는데, 이때 이 12개의 블록을 직접 블록이라고 한다.

![image](https://github.com/yumin00/blog/assets/130362583/d79c0186-e3de-4be2-976d-58722a02ee23)

만약, 파일이 12개 이하의 데이터가 가지고 있다면 1번에서 추가 작업이 필요하지 않다.

### 2. 1번으로 충분하지 않다면 13번째 주소에는 단일 간접 블록 주소를 저장한다.
13번째 주소에는 데이터의 주소인 블록을 저장하는 것이 아니라, 여러 개의 블록의 주소를 가지고 있는 단일 간접 블록의 주소를 저장한다.
![image](https://github.com/yumin00/blog/assets/130362583/d47b561a-6f5c-4d35-8e10-9196ee840df7)

### 3. 2번으로 충분하지 않다면 14번째 주소에 이중 간접 블록 주소를 저장한다.
14번재 주소에는 단일 간접 블록들의 주소를 가진 이중 간접 블록의 주소를 저장한다.
![image](https://github.com/yumin00/blog/assets/130362583/6f1303b2-9725-4323-9eaf-4e9984e36569)

### 4. 3번으로 충분하지 않다면 15번째 주소에는 삼중 간접 블록 주소를 저장한다.
15번째 주소에는 이중 간접 블록들의 주소를 가진 삼중 간접 블록의 주소를 저장한다.
![image](https://github.com/yumin00/blog/assets/130362583/43236f42-ad41-46f3-8f10-d0d1d55d8186)

삼중 간접 블록까지 이용하면 웬만한 크기의 파일은 모두 표헌할 수 있다!

### 장점
- 안정적이고 신뢰성이 높다.
- 파일에 대한 읽기,쓰기,실행 권한을 세분화하여 설정할 수 있다.
- 메타데이터를 효율적으로 관리하며, 디스크 공간을 최적화한다.

### 단점
- 다른 운영 체제와의 호환성 문제가 발생할 수 잇다.
- 초기 설정에 inode 수를 설정해야하는데, 수를 초과하면 더이상 파일을 생성할 수 없다.
