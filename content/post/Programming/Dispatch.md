---
title: "Dispatch에 대해 알아보자"
date: 2024-03-25T19:49:22+09:00
draft: false
categories :
  - Programming
---

# Dispatch에 대해 알아보자
# Dispatch
OOP(객체 지향 프로그래밍)의 특징은 특정한 객체의 메서드를 호출하더라도, 실제로 실행되는 메서드는 상황에 따라 달라질 수 있다는 것이다. 즉, 다형성에 의해 상황에 따라 실제로 실행되는 메서드는 항상 달라진다. (더 자세한 내용은 [여기](https://yumin.dev/p/%EA%B0%9D%EC%B2%B4-%EC%A7%80%ED%96%A5-%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%A8%EB%B0%8D%EC%97%90-%EB%8C%80%ED%95%B4-%EC%95%8C%EC%95%84%EB%B3%B4%EC%9E%90/)에서 확인할 수 있다.)

```python
```python
class Human:
    def __init__(self, name):
        self.name = name
        self.legs = 2
        self.arms = 2
    
    def say_hi(self):
        return f"hi, My name is {self.name}"

class Student(Human):
    def __init__(self, name, age, score):
        super().__init__(name)
        self.age = age
        self.score = score

    def say_hi(self):
        return f"Hi, I'm student and my name is {self.name}"


class Employee(Human):
    def __init__(self, name, age, score):
        super().__init__(name)
        self.age = age
        self.score = score

    def say_hi(self):
        return f"Hi, I'm employee and my name is {self.name}"
```

해당 코드에서는 Human이라는 부모 클래스가 있고, 해당 클래스를 상속받는 Studen, Employee 클래스가 있다. 그렇다면 아래 코드에서는 어떤걸 출력하게 될까? 

```python
yumin.say_hi()
```

정답을 알 수 없다. 이다! 객체 지향 프로그래밍에서는 동적 바인딩이 진행되기 때문에 상황에 따라서 실제로 실행되는 메서드가 결정되기 때문에 당장 저 코드만으로는 알 수 없다. 

이렇게, 어떤 메서드를 실행시킬 것인지 결정하여 그것을 실행시키는 매커니즘을 바로 dispatch 라고 한다.

더 알아보니, 언어마다 dispatch 를 어떻게 사용하는지가 다르다.

## Python Dispatch
파이썬에서 디스패치는 특정 조건이나 상황에 따라서 어떤 메서드를 호출할지 결정하는 매커니즘이다. 즉, 파이썬에서 디스패치를 오버로딩을 위한 것이다.

> 오버로딩(OverLoading)
> 
> 오버로딩이란, 같은 이름의 메서드가 있더라고 매개변수 혹은 타입이 다르면 같은 이름을 사용하여 메서드를 정의하는 것이다.

파이썬에서 디스패치는 싱글 디스패치(Single Dispatch)와 멀티 디스패치(Multiple Dispatch)가 있다.

### 싱글 디스패치
싱글 디스패치란, 함수의 첫 번째 인자 타입에 따라 실행될 메서드를 결정하는 것이다. 이때, 파이썬에는 `functools.singledispatch` 를 사용하여 구현할 수 있다.

```python
from functools import singledispatch

@singledispatch
def say_hello(arg):
    print(f"Hi, My name is {arg}")

@say_hello.register(int)
def _(arg):
    print(f"Hi, I'm {arg} years old")

@say_hello.register(list)
def _(arg):
    print(f"Hi, I have {len(list)} books")

say_hello("yumin")   # Hi, My name is yumin
say_hello(23)  # Hi, I'm 23 years old
say_hello(['전공책1', '전공책1'])  # Hi, I have 2 books 
```

해당 코드처럼 `say_hello`라는 함수가 있더라도, 첫 번째 인자 타입이 무엇인지에 따라 실제로 실행되는 메서드는 달라질 수 있다.

### 멀티 디스패치
멀티 디스패치란, 두 개 이상의 인자의 타입을 기준으로 적절한 함수를 호출하는 것이다. 파이썬에서는 `multipledispatch` 패키지를 사용하여 구현할 수 있다.

```python
# multipledispatch 패키지 설치 필요: pip install multipledispatch
from multipledispatch import dispatch

@dispatch(int, int)
def add(x, y):
    return x + y

@dispatch(str, str)
def add(x, y):
    return x + y

print(add(1, 2))  # 3
print(add("Hello, ", "world!"))  # Hello, world!
```

해당 코드처럼, `add`라는 함수는, 어떤 인자를 받느냐에게 따라 결정되는 메서드가 달라진다.

## Swift Disaptach
swift에서 디스패치는 정적 디스패치와 동적 디스패치가 있다.

- 정적 디스패치: 컴파일 타임 때 메서드 호출이 결정된다.
- 동적 디스패치: 런타임 때 어떤 메서드를 호출할지 결정된다.

이 두 개념은 정적 바인딩 / 동적 바인딩과 거의 비슷한 개념임을 알 수 있다.

그러면다면, 같은 디스패치이지만, 파이썬과 스위프트에서는 다르게 정의하고 사용하는 것일까? 그 이유는 바로 언어의 타입에 있다.

## 동적 언어 / 정적 언어
먼저, 파이썬은 동적 언어로, 런타임 때 어떤 메서드를 호출할지 결정하는 동적 디스패치를 자연스럽게 지원한다.
또한, 파이썬의 동적 특성을 강화시키기 위해 타입별로 다른 함수를 실행시킬 수 있도록 싱글 디스패치 / 멀티 디스패치를 지원한다.

하지만 스위프트는 정적 언어로, 자연스럽게 컴파일 때 메서드를 호출하는 정적 디스패치를 지원한다. 하지만 더 높은 수준의 유연성과 다형성을 위해 런타임 때 어떤 메서드를 호출할지 결정하는 동적 디스패치도 제공한다.