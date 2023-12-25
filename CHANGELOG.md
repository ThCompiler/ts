## 0.0.1-alpha

This is the initial release.

### Supported

* Check expected results ```ts.ToTestValuesExpected(expected ...interfaces{})``` or ```ts.TTVE(expected ...interfaces{})```.
* Check expected results with nil error ```ts.ToTestValuesExpectedWithNilError(expected ...interfaces{})```
  or ```ts.TTVEWNE(expected ...interfaces{})```.
* Check expected errors (expected error as last returned parameters) .
  ```ts.ToTestErrorExpected(err error)``` or ```ts.TTEE(err error)```
* Check panic error ```ts.ToTestPanicErrorExpected(msg interface{})``` or ```ts.TTPEE(msg interface{})```.
* Check not nil error ```ts.ToTestCheckErrorExpected()``` or ```ts.TTCEE()```.
