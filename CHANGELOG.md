## 0.0.1-alpha

This is the initial release.

### Added

* Added functions for run list of testcases.
* Added functions ```ts.ToTestValuesExpected(expected ...interfaces{})``` or ```ts.TTVE(expected ...interfaces{})``` for checking expected results without error.
* Added functions ```ts.ToTestValuesExpectedWithNilError(expected ...interfaces{})```
  or ```ts.TTVEWNE(expected ...interfaces{})``` for checking expected results with nil error.
* Added functions ```ts.ToTestErrorExpected(err error)``` or ```ts.TTEE(err error)``` for checking expected errors (expected error as last returned parameters).
* Added functions ```ts.ToTestPanicErrorExpected(msg interface{})``` or ```ts.TTPEE(msg interface{})``` for checking panic error.
* Added functions ```ts.ToTestCheckErrorExpected()``` or ```ts.TTCEE()``` for checking not nil error.

### Github

* Added ci and release github actions
