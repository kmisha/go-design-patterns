# fan-in-pattern-go

An example of Fan-in pattern implementation

## Task

В нашей системе имеется несколько моделей данных:

- Product
- User
- Order
- Cart

> Данные модели, могут содержать любые атрибуты, главное они должны содержать атрибут id имеющий тип uuid.

**Задача:**
Мы начали разработку нового сервиса статистики, отлеживающего различные параметры нашей системы. Изначально при выполнении изменения любой из моделей отправляются в сортов: канал _productUpdatesCh_, _userUpdatesCh_, _orderUpdatesCh_, _cartUpdatesCh_ соотв.

Мы хотим отследить количество изменений в нашей системе. Для этого необходимо разработать функцию, объединяющую эти каналы в один канал обновлений. Данный канал используется для получения всех событий-изменений о всех сущностях.

Необходимо разработать функцию, реализующую паттерн fan-in, fan-out, примерный интерфейс функции

`MergeData(ctx context.Context, chans …SomeInterface) chan UnifiedStruct.`

## Requirements

- the system manages a product, a user, an order, and a cart (entities). We can create, update, and delete these entities.
- the system sends all changes to channels
- the system tracks the number of changes

## How to run

`go run cmd/main.go
