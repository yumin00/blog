---
title: "Airflow 튜토리얼"
date: 2024-03-04T20:59:35+09:00
draft: false
categories :
- DataEngineering
---

# Airflow 튜토리얼
## Airflow란?
에어플로우에 대해 간단히 설명하자면, 방향성을 가진 테스크들을 순차적으로 진행시킬 수 있는 관리자라고 할 수 있다. 좀 더 자세한 내용은 여기에서 확인할 수 있다.

## Airflow 튜토리얼
### 🛠️ 설치 (Mac)
간단한 설치를 위하여 에어플로우에서 제공하고 있는 docker-compose 파일을 사용하고자 한다.
```shell
curl -LfO 'https://airflow.apache.org/docs/apache-airflow/2.6.1/docker-compose.yaml'
```

위 명령어를 통해 에어플로우 docker-compose.yaml 파일을 받을 수 있다.

docker-compose.yaml 파일이 생성되었다면, 다음 명령어를 통해 도커 컨테이너 서비스를 띄울 수 있다.
```shell
docker compose up -d
```

모두 실행되면, `http://localhost:8080/` 를 통해 에어플로우에 접근할 수 있다.

<img width="875" alt="image" src="https://github.com/yumin00/blog/assets/130362583/cb9fce73-a17e-4c33-8011-8b56612a595e">
로그인 페이지에서는 기본적으로 airflow / airflow 로 로그인이 가능하다.
w
### 📝 Hello World DAG 생성
여기에서 이야기한 것처럼, DAG에 정의되는 TASK들은 오퍼레이터를 통해 정의된다.

간단하게 PythonOperator를 사용하여 hello world를 출력하는 task를 가지는 DAG를 생성해보자.

DAG를 생성하기 위해 위에서 생성한 레포에서 `~dags` 폴더에 생성하고자 하는 DAG인 `hello_world.py` 파일을 생성해주었다.

<img width="756" alt="image" src="https://github.com/yumin00/blog/assets/130362583/62e9f4ac-bbf7-4391-bb00-d048b1d3ceea">

- `with DAG` 를 통해 DAG를 정의할 수 있다. DAG 이름과 start_date, tags 등 기본 값을 설정할 수 있다.
- `as dag:`를 통해 TASK를 정의할 수 있다. 정의한 테스크는 `PythonOperator`를 통해 "Hello World!"를 출력한다.

해당 DAG는 하나의 테스크이므로 순서를 정의하지 않아도 된다. 만약, 테스크가 여러개라면 `t1 >> t2` 와 같은 형식으로 순서를 정의해야하 한다.

### 🖥️ Web UI를 통한 DAG 확인
작성한 DAG는 에어플로우 UI를 통해서 확인할 수 있다.
<img width="1710" alt="image" src="https://github.com/yumin00/blog/assets/130362583/bb3eaa45-06c9-47ad-8993-b9bb9328a4ac">

해당 DAG에 들어가보면 더 자세한 내용을 확인해볼 수 있다.

- Graph : 설정한 TASK를 그래프로 확인할 수 있다.

<img width="408" alt="image" src="https://github.com/yumin00/blog/assets/130362583/3c3c15d6-1e0c-4702-b5e7-c88287c2ac16">

- DAG의 실행 결과를 확인할 수 있다.

<img width="305" alt="image" src="https://github.com/yumin00/blog/assets/130362583/b6689df3-a927-4949-88a2-63a9dc406dd8">
