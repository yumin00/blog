---
title: "[과제] CSV 대규모 데이터 처리 파이프라인"
date: 2024-07-22T19:14:59+09:00
draft: true
categories :
- DataEngineering
---

약 191만 개의 행과 17개의 열을 가진 CSV 파일을 처리하여
각 도시에서 가장 많이 팔린 Maker를 추출하고 (parquet 형식)[]으로 저장하는 파이프라인을 구축하는 과제를 해결하는 과정에 대해 정리해보았다.

# 1. 어떻게 데이터를 처리할 것인가?
먼저, CSV 파일을 처리하기 위해 파이프라인을 어떤식으로 구성할지 결정해야 한다. 어떤 방식을 선택할지에 대한 기준이 필요한데 나의 기준은 다음과 같다.

- 대규모 데이터를 처리할 수 있어야 한다.
- 러닝 커브가 낮아야 한다.

내가 고려한 방법은 두 가지이다.

## Python
- Python의 라이브러리는 CSV 파일을 효율적으로 읽고, 변환 및 분석하는 데 탁월하다.
- 대규모 병렬 처리 작업에 대한 제한이 있다.

## Go
- 성능이 우수하고 실행 속도가 빠르다.
- 데이터 처리 및 분석에 특화된 라이브러리가 부족하다.


Go 언어를 사용하면 실행 속도가 빠르지만, 데이터 처리 관련 라이브러리가 부족하여 파이썬의 대규모 데이터 처리 라이브러리를 사용한다면 Go만큼의 성능이 나올 것이라고 판단하여 파이썬을 선택했다.

파이썬에 Dask라는 가상 데이터프레임 라이브러리가 존재하는데, 대규모 데이터셋을 병렬로 처리하여 대용량 데이터 처리 성능을 향상시키고, 병목 현상을 줄여준다고 한다. 그래서, 파이썬의 Dask를 사용하여 코드를 작성해보았다.

## 2. 코드 작성
```
import dask.dataframe as dd


if __name__ == "__main__":
    csv_file_path = 'data.csv'
    parquet_file_path = 'result.parquet'

    ddf = dd.read_csv(csv_file_path, assume_missing=True)

    # 각 City에서 가장 많이 팔린 Maker를 추출
    # groupby 연산 후 size()로 각 Maker의 판매 수를 세고 idxmax()를 통해 가장 많이 팔린 Maker를 추출
    most_sold_maker = ddf.groupby('City')['Make'].apply(lambda x: x.value_counts().idxmax(), meta=('Make', 'str'))

    # 결과를 데이터프레임으로 변환
    result_df = most_sold_maker.compute().reset_index()
    result_df.columns = ['City', 'Most_Sold_Maker']

    # 결과를 parquet 형식으로 저장
    result_df.to_parquet(parquet_file_path, engine='fastparquet', index=False)

    print("Processing completed and saved to Parquet file.")
```

parquet 형식의 파일이 제대로 저장되었는지 확인하기 위해 아래와 같이 parquet 형식의 파일을 csv로 변환하는 코드를 구현했다. 실제로 제대로 실행된 것을 확인할 수 있다.

```python
import pandas as pd
import dask.dataframe as dd



if __name__ == "__main__":
    csv_file_path = 'data.csv'
    parquet_file_path = 'result.parquet'

    ddf = dd.read_csv(csv_file_path)

    # 각 City에서 가장 많이 팔린 Maker를 추출
    # groupby 연산 후 size()로 각 Maker의 판매 수를 세고 idxmax()를 통해 가장 많이 팔린 Maker를 추출
    most_sold_maker = ddf.groupby('City')['Make'].apply(lambda x: x.value_counts().idxmax(), meta=('Make', 'str'))

    # 결과를 데이터프레임으로 변환
    result_df = most_sold_maker.compute().reset_index()
    result_df.columns = ['City', 'Most_Sold_Maker']

    # 결과를 parquet 형식으로 저장
    result_df.to_parquet(parquet_file_path, engine='fastparquet', index=False)

    print("Processing completed and saved to Parquet file.")

    df = pd.read_parquet('result.parquet')
    df.to_csv('projects-0701.csv')
````

<img width="515" alt="image" src="https://github.com/user-attachments/assets/c088bf48-0ca6-41dc-84e5-995c79df476b">