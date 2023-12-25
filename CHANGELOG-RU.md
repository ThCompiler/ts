## 0.0.1-alpha

Начальный релиз

### Добавлено

* Добавлены функции для запуска списка testcasов.
* Добавлены функции ```ts.ToTestValuesExpected(expected ...interfaces{})```,
```ts.TTVE(expected ...interfaces{})``` для проверки ожидаемых результатов без ошибок.
* Добавлены функции ```ts.ToTestValuesExpectedWithNilError(expected ...interfaces{})```,
  ```ts.TTVEWNE(expected ...interfaces{})``` для проверки ожидаемых результатов с nil ошибкой.
* Добавлены функции ```ts.ToTestErrorExpected(err error)```,
```ts.TTEE(err error)```  для проверки ожидаемых ошибок (ожидаемая ошибка в качестве последнего возвращаемого из функции параметра).
* Добавлены функции ```ts.ToTestPanicErrorExpected(msg interface{})```, ```ts.TTPEE(msg interface{})``` для проверки ошибки паники.
* Добавлены функции ```ts.ToTestCheckErrorExpected()```, ```ts.TTCEE()``` для проверки не nil ошибки.

### Github

* Добавлен CI для тестов и проведения релиза
