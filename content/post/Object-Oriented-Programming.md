---
title: "객체 지향 프로그램밍에 대해 알아보자"
date: 2024-03-25T20:17:23+09:00
draft: false
---

# 객체 지향 프로그래밍 Object-Oriented Programming(OOP)
객체 지향 프로그래밍이란, 프로그램을 독립된 객체들의 집합이라고 보는 방식이다. 각 객체들은 서로 메시지를 전달하며 상호작용하고, 각각의 객체는 고유한 데이터와 해당 데이터를 조작하는 메소드로 구성되어 있다.

## 캡슐화
객체 지향 프로그래밍에서는 캡슐화라는 개념이 있다. 이는 data와 method를 함께 묶는 것을 의미한다.

data와 method를 다음과 같이 Class로 묶을 수 있다.

```python
class Student:
    def __init__(self, name, age, score):
        self.name = name
        self.age = age
        self.score = score
        self.school = "A"
```

이와 같이 class를 사용하면 Student라는 객체를 쉽게 만들어낼 수 있으며, `self.school` 처럼 초기값을 설정할 수도 있다.

class 안에는 다양항 method를 정의할 수 있다. 그리고 해당 class에서 생성된 모든 객체는 class 에 정의되어 있는 모든 method를 호출할 수 있다.

method를 정의함으로써, 데이터를 조작할 수 있게 되는 것이다. 즉, 메서드를 통해 데이터를 조작하므로써 데이터를 간접적으로 접근하도록 하여 데이터를 보호한다.
```python
class Student:
    def __init__(self, name, age, score):
        self.name = name
        self.age = age
        self.score = score
        self.school = "A"
    
    def say_hi(self):
        return f"Hi, my name is {self.name}"
```

## 상속(Inheritance)
상속이란, 자식 class가 부모 class의 값을 상속 받는 것이라고 이해할 수 있다. 이는 코드 중복을 줄일 수 있다.

```python
class Human:
    def __init__(self, name):
        self.name = name
        self.legs = 2
        self.arms = 2

class Student(Human):
    def __init__(self, name, age, score):
        super().__init__(name)
        self.age = age
        self.score = score

    def say_hi(self):
        return f"Hi, my name is {self.name}"

if __name__ == '__main__':
    yumin = Student("yumin", 25, 100)
    print(yumin.legs)
```

Human이라는 부모 Class를 상속 받는 Student 를 구현함으로써, Student class에서는 Human의 정보를 상속 받을 수 있다. 즉, Student class를 갖는 객체의 legs를 출력하면 2가 나온다.

이렇게 상속이라는 개념을 사용함으로써, 중복되는 코드를 방지할 수 있다는 장점이 있다.

## 추상화
추상화란, method의 구현 세부 정보를 숨기는 것이라고 할 수 있다. 즉 interface를 지정함으로써 세부 정보를 보여주지 않고 method를 사용할 수 있게 하는 것이다.

추상화의 장점은 내부 로직이 변경되더라도, 사용자는 interface만 사용하였기 때문에 코드를 수정하지 않아도 된다는 것이다.

## 다형성 (Polymorphism)
다형성이란, 같은 부모 class 안에 있어서 같은 메서드를 호출할 수 있다고 하더라도, 서로 다른 동작을 할 수 있도록 하는 것을 말한다.

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

위와 같이 서로 같은 Human 부모 클래스를 상속받아 say_hi라는 메서드를 사용할 수 있지만 서로 다른 구현 방식을 가지고 있는 것이다. 이를 method overriding 이라고 한다.

이떄 중요한 것은, 부모 클래스와 결과값이 같아야한다. 즉, 부모 클래스에서 say_hi는 string을 return하기 때문에 자식 클래스에서도 string을 return해야 한다.

오버라이딩이 존재함에도 불구하고, 메소드가 어떻게 작동해야하는지에 대한 규칙이 있기 때문에 클래스의 핵심은 그대로 있고 구현 방식과 모양을 다양하게 할 수 있다는 장점이 있다. 