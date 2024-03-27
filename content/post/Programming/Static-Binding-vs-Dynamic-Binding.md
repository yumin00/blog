---
title: "Static Binding vs Dynamic Binding"
date: 2024-03-27T10:33:26+09:00
draft: false
categories :
- Programming
---

# 정적 바인딩(Static Binding)과 동적 바인딩(Dynamic Binding)에 대해 알아보자

# 바인딩(Binding)
프로그램이 실행되기 위해서는 변수/함수가 메모리에 할당된다. 프로그램을 실행시키기 위해서 메모리에 할당된 변수/함수의 위치(주소)로 연결시키는 과정을 **바인딩**이라고 한다.

바인딩이 언제 진행되는지에 따라 정적 바인딩과 동적 바인딩으로 나뉠 수 있다.

# 정적 바인딩(Static Binding)
## 정적 바인딩이란?
정적 바인딩이란, 변수와 함수가 컴파일 타임에 바인딩되는 것을 말한다.

> 컴파일 타임
> 
> 컴파일이 이루어지는 시점 

즉, 컴파일 타임에 호출될 함수가 미리 정해지고, 호출될 함수로 점프할 주소가 미리 바인딩된다.

변수의 경우에는 해당 변수의 타입에 맞춰 메모리에 변수가 할당되고, 값이 변경되더라도 타입이 동일하다면 오류 없이 해당 주소에 값이 변경된다. 예를 들면 다음과 같다.

```go
var x int32 //int 타입에 맞춰 메모리 할당
x = 1  // 해당 주소 1 할당

x = x + 5 // 해당 주소에 1은 사라지고 6으로 할당

x = "hi" // int 타입이 아니기 때문에 오류
```

## 장점
- 컴파일 타임에 모든 것이 결정되기 때문에 런타임에 바인딩을 확인할 필요가 없어서 더 빠른 실행 속도를 제공한다.
- 컴파일 시에 타입 체크가 이루어지기 때문에 타입 관련 오류를 컴파일 시에 확인하고 수정할 수 있다.

## 단점
- 런타임에 바인딩이 결정되지 않기 때문에 코드 재사용성이나 확장성이 동적 바인딩에 비해 제한적일 수 있다.
- 동적 바인딩에 비해 다형성이 제한적이다.

# 동적 바인딩(Dynamic Binding)
## 동적 바인딩이란?
동적 바인딩이란, 변수와 함수가 런타임에 바인딩되는 것을 말한다.

> 런타임
> 
> 프로그램이 실행되는 시점

런타임에 실제로 사용된 객체의 클래스형에 의해 호출될 함수가 결정되고,이때 호출될 함수로 점프하는 과정인 바인딩이 진행된다.

동적 바인딩은 파이썬 예제를 보며 이해할 수 있다.


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

if __name__ == '__main__':
    yumin = Employee("yumin", 25, 100)
    print(yumin.say_hi())
```

해당 함수에서 yumin은 Employee로 할당되었지만, Employee의 부모 클래스는 Human이다. 즉, say_hi를 출력하면 Human 클래스의 메서드에 따라서, "hi, My name is yumin" 이 출력되어야 하지만,
동적 바인딩이 진행됨에 따라 값들이 미리 할당되지 않고 런타임에 실제로 사용된 타입에 따라 호출될 메서드가 결정되기 때문에 "Hi, I'm employee and my name is yumin"이 출력될 수 있는 것이다.

## 장점
- 런타임에 바인딩이 결정되기 때문에 유연한 코드 작성이 가능하다. (ex. 상속, 오버라이딩)
- 같은 인터페이스나 부모 클래스를 가진 서로 다른 객체들이 런타임에 결정되기 때문에 다형성을 쉽게 구현할 수 있다.

## 단점
- 런타임에 타입을 체크하고 어떤 메서드를 호출할지 결정하기 때문에 정적 바인딩에 비해 성능이 느릴 수 있다.
- 컴파일 시에 타입 검사가 이루어지지 않기 때문에 오류 발견이 늦어질 수 있다.