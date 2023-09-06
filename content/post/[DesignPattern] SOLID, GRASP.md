---
title: "[Design Pattern] SOLID, GRASP 원칙"
date: 2023-09-06T13:00:16+09:00
draft: true

categories :
- Architecture
---


소프트웨어 개발은 요구사항에 맞춰 계속 변화하고 개발되어야 한다. 따라서 소프트웨어는 유연해야 하고 쉽게 변화할 수 있어야 한다.

디자인패턴이란 좋은 코드를 설계하기 위한 설계 디자인 방법론이다. 즉, 좋은 설계란 시스템에 새로운 요구사항이나 변경사항이 생겼을 때 영향을 
받는 범위가 적은 구조를 말한다.

좋은 코드를 설계하기 위해서 객체지향 방법론에서는 SOLID 원칙이 있다.

> 객체지향
> 
> 응집도는 높게 / 결합성은 낮게

# SOLID 원칙
SOLID 원칙이란, 객체지향 설계에서 지켜줘야할 5개의 원칙 (SRP, OCP, LSP, ISP, DIP) 를 말한다.

## 1. SRP(Single Responsibility Principle) 단일 책임 원칙
- 소프트웨어의 설계 부품(클래스, 함수 등)은 단 하나의 책임을 가져가야 한다.

## 2. OCP(Open-Closed Principle) 개방-폐쇄 원칙
- 기본 코드는 변경하지 않으면서(closed), 기능을 추가할 수 있도록(open) 설계해야 한다.

이는 확장에는 개방적이고, 수정에는 폐쇄적이어야 한다는 것이다. 이를 만족하는 설계를 위해서는 **캡슐화** 를 진행해야 한다.

> 캡슐화
> 
> 클래스 안에 서로 연관있는 기능들을 하나의 캡슐로 만들어 데이터를 보호하는 것

## 3. LSP(Liskov Substitution Principle) 리스코프 치환 원칙
- 자식 클래스는 부모 클래스에서 가능한 행위를 수행할 수 있어야 한다.

## 4. ISP(Interface Segregation Principle) 인터페이스 분리 원칙
- 한 클래스에서 사용하지 않는 인터페이스는 구현하지 말아야하고, 거대한 인터페이스보다 여러 개의 구체적인 인터페이스가 낫다.

## 5. DIP(Dependency Inversion Principle) 의존 역전 원칙
- 의존 관계를 맺을 때, 변화하기 쉬운 것보다 변화하기 어려운 것에 의존해야 한다.

# GRASP 원칙
GRASP 원칙(General Responsibility Assignment Software Patterns) 이란, 각 객체에 책임을 부여하는 원칙을 말하는 패턴이다.


## 1. High Cohesion(높은 응집력)
- 각 객체가 밀접하게 연관된 책임만 가지도록 구성
- 한 객체, 한 시스템이 자신에게 주어진 책임만 수행하도록 구성
- 자신의 기능을 수행하기 위해 다른 객체나 시스템을 참조하는 일이 적으며 자연스럽게 Low Coupling 이 된다.

## 2. Low Coupling(낮은 의존성)
- 시스템 간에 상호 의존도가 낮게 책임을 부여
- 시스템의 재사용성을 높이고, 시스템 관리를 편하게 한다.

## 3. Information Expert
- 객체란 상태와 행동을 함께 가지는 단위로 표
- 책임을 수행하는데 필요한 데이터를 가지고 있는 객체에 책임을 부여하는 것

## 4. Polymorphism(다형성)
- if else / switch 와 같은 논리 조건을 사용한다면, 새로운 변화가 있을 경우 해당 논리 구조를 수정해야 한다. 다형성 패턴은 객체의 타입을 검사하여 타입에 따라 여러 대안을 수행하는 조건인 
  논리를 사용하지 말라고 경고한다.
- 다형성을 이용해 타입을 분리하고, 변화하는 행동을 각 타입의 책임으로 할당해라.

## 5. Pure Fabrication
- Information Expert에 근거하여 필요한 데이터를 가지고 있는 객체에 책임을 부여했는데, high cohesion / low coupling 에 위반된다면 어떻게 해야하는가?
- Pure Fabrication : 인위적으로 한 클래스를 만들어 문제가 되는 책임을 모아 high cohesion을 갖게해라

## 6. Indirection
- 두 객체 사이의 의존성을 피하고 싶다면, 그 사이에 다른 매개체를 통해 전달

## 7. Protected Variations
- 안정된 인터페이스를 정의하여 사용

