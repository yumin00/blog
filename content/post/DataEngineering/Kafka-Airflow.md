---
title: "Airflow 에서 Kafka 활용하기"
date: 2024-03-04T21:53:20+09:00
draft: true
categories :
- DataEngineering
---

# Airflow 에서 Kafka 활용하기
에어플로우는 배치성 작업에 사용되기 때문에, 스트리밍 서비스와 결합을 위해 Kafka를 사용하게 될 수 있다.

이번 글에서는, 에어플로우에서 카프카를 효과적으로 사용하기 위한 방법에 대해 이야기해보고자 한다.

에어플로우 설치 및 간단한 DAG를 생성하는 방법은 여기에서 확인할 수 있다.

## ✏️ Kafka Sensor & Trigger DAG 생성
특정 토픽에 생성되는 메시지를 consume하는 DAG를 생성하고자 한다.
해당 DAG는 메시지가 consume될 때마다 실행시키기 위해 `AwaitMessageTriggerFunctionSensor` 센서를 사용했다.

카프카에서 메시지를 consume하여 작동하는 DAG가 많아진다면, 각 DAG들은 카프카에 메시지가 생성될 때마다 실행되어야 한다.
그러면, 카프카와 결합하는 DAG들은 센서를 사용해야하기 때문에 배치성 작업에 사용하는 에어플로우의 성격을 잃을 수 있다.

그렇기 때문에, 메시지를 consume하는 DAG를 만들고, 해당 DAG가 consume한 메시지의 성격에 따라 필요한 DAG를 trigger 하고 trigger 된 DAG에게 카프카를 통해 메시지를 전달할 수 있도록 구현해보고자 한다.
아키텍처는 다음과 같다.

<img width="497" alt="image" src="https://github.com/yumin00/blog/assets/130362583/2e737eaf-3536-4a5a-bc05-ed126010f648">

Kafka에서 메시지를 consume하고 필요한 DAG를 tirgger하는 DAG를 생성해보자.

`~dags` 폴더에 `kafka_consumer_to_trigger.py` 파일을 생성했다.

- 토픽에서 메시지를 consume하여 다른 DAG를 trigger하는 DAG의 기본값을 설정할 수 있다.

<img width="622" alt="image" src="https://github.com/yumin00/blog/assets/130362583/5a3e3e14-cccb-4864-ad2c-0214e3b6b013">

- 해당 DAG는 특정 시간에만 실행되는 것이 아니라, 한 topic을 계속 바라보다가 message가 생성되면 동작할 수 있도록 `AwaitMessageTriggerFunctionSensor`를 사용했다.
- AwaitMessageTriggerFunctionSensor
    - topics: consume할 토픽을 여러개 설정할 수 있다.
    - apply_function: 메시지를 consume했을 때 실행할 함수를 설정할 수 있다.
    - event_triggered_function: `apply_function`에서 return한 값은 `event_triggered_function`의 파라미터로 사용된다.
    - kafka_config_id: `AwaitMessageTriggerFunctionSensor` consumer의 config 값을 설정할 수 있다.

### ⚙️ Kafka Config 설정
<img width="853" alt="image" src="https://github.com/yumin00/blog/assets/130362583/ff599ecc-0933-47b0-bf41-1c32b076c85f">

에어플로우 Web에서 Admin > Connections 에서 Kafka config를 설정할 수 있다. AwaitMessageTriggerFunctionSensor 를 통해 메시지를 consume할 consumer의 config를 설정하는 것이다.


<img width="1755" alt="image" src="https://github.com/yumin00/blog/assets/130362583/c502294d-dea7-42ed-b416-b79aa2f5d4a5">

- Connection Id: 설정한 task_id ex) hello_world_message_consumer
- Connection Type: Apache Kafka
- Config Dic: `bootstrap.servers`와 같은 kafka config 값을 넣어주면 된다.

### 📥 kafka_consumer_to_dag_trigger
- apply_function: 메시지를 consume했을 때, key가 'hello.world'면, 'hello.world'라는 토픽에 메세지를 produce한 뒤, key를 return한다.

<img width="871" alt="image" src="https://github.com/yumin00/blog/assets/130362583/732a24d6-9b8b-48f4-b861-7be0d2d0752b">

- event_triggered_function: key를 파라미터로 받고, key가 'hello.world'면, `print_hello_world` DAG 를 실행시킨다.

<img width="752" alt="image" src="https://github.com/yumin00/blog/assets/130362583/7fe779c1-96ec-4dba-9458-cec4799ef40f">

## 🔗 Kafka Consume & Process DAG 생성
위 DAG를 통해 실행되는 DAG를 구현해보자.

해당 DAG는 hello.world 에서 메시지를 consume하고, 해당 메시지를 print하는 DAG이다.

<img width="648" alt="image" src="https://github.com/yumin00/blog/assets/130362583/fc6df4e6-f12d-4ddc-88bd-6e791de87e82">

consume하는 Taks는 PythonOperator 혹은 Kafka Operator를 사용할 수 있다. 먼저, 각 테스크 간의 메시지 통신을 위해 PythonOperator 를 사용할 것이다.

- consume_message

<img width="940" alt="image" src="https://github.com/yumin00/blog/assets/130362583/e1a382a0-1561-4885-82d7-5b811458958e">

- print_hello_world

<img width="615" alt="image" src="https://github.com/yumin00/blog/assets/130362583/326c406a-4b7a-41dc-9bc3-288d5816fb61">

각 Task 간의 메시지 통신을 위해 xcom을 사용할 수 있다. `xcom_push` 를 통해 메시지와 key를 넣을 수 있고 `xcom_pull`를 통해 키 값으로 메시지를 가져올 수 있다.

