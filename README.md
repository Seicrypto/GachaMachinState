# Limited Pool Gacha Machin State

Practice design of a gacha machine using the state pattern in Go.

## Limited Pool Gacha Diagram

### Gacha State Flow

```mermaid
---
title: Limited Pool Gacha State Pattern
---

stateDiagram
%% Declare interfaces
    Idle: Idle State
    Dispensing: Dispensing State
 state ifEmpty <<choice>>
    SoldOut: Sold Out State
    Maintenance: Maintenance State

%% Relationships
Idle --> Dispensing : Receive order
Dispensing --> ifEmpty : Order processed
ifEmpty --> Idle : if gacha NOT empty
ifEmpty --> SoldOut : if gacha empty
SoldOut --> Maintenance : Charging or changing content
Maintenance --> Idle : Maintenance finished
```

### Gacha Interface Class

```mermaid
---
title: Gacha Abstract Class
---

classDiagram
%% Declare
class State {
    <<interface>>
    HandleRequest(*Context)
}
```

## Order Diagram

If it were official project, please use [Transition](https://github.com/qor/transition) and [GORM](https://github.com/go-gorm/gorm).

```mermaid
---
title: Order State Pattern
---

stateDiagram
%% Declare interfaces
    New: NewOder State
    Paid: Paid State
    Cancel: Cancel State

%% Relationships
[*] --> New : Initail oreder
New --> Paid : Payment finished
Paid --> Cancel : Cancel or reject
New --> Cancel : Cancel or payment rejected
Paid --> [*] : Order completed
Cancel --> [*] : Order canceled
```
