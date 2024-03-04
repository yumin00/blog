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

## Kafka Sensor & Trigger DAG 생성
이를 위해, Kafka에서 메시지를 consume하고 필요한 DAG를 tirgger하는 DAG를 생성해보자.

`~dags` 폴더에 `kafka_consumer_to_trigger.py` 파일을 생성해주었다.

<img width="962" alt="image" src="https://github.com/yumin00/blog/assets/130362583/f203e880-5757-4155-a828-409feab7b88d">

- 토픽에서 메시지를 consume하여 다른 DAG를 trigger하는 DAG의 기본값을 설정할 수 있다.


<img width="622" alt="image" src="https://github.com/yumin00/blog/assets/130362583/5a3e3e14-cccb-4864-ad2c-0214e3b6b013">

- 해당 DAG는 특정 시간에만 실행되는 것이 아니라, 한 topic을 계속 바라보다가 message가 생성되면 동작할 수 있도록 `AwaitMessageTriggerFunctionSensor`를 사용했다.
- AwaitMessageTriggerFunctionSensor
    - topics: consume할 토픽을 여러개 설정할 수 있다.
    - apply_function: 메시지를 consume했을 때 실행할 함수를 설정할 수 있다.
    - event_triggered_function: `apply_function`에서 return한 값은 `event_triggered_function`의 파라미터로 사용된다.
    - kafka_config_id: `AwaitMessageTriggerFunctionSensor` consumer의 config 값을 설정할 수 있다.

이때, `kafka_config_id`는 다음과 같이 설정할 수 있다.
