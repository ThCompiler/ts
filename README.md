# Testcases (ts)

## About

That is simple wrap [testify's](https://github.com/stretchr/testify) Suite to write testcases and run them. 

It's some kind of experiment. 

Support:
* Check expected results ```ts.ToTestValuesExpected(expected ...interfaces{})``` or ``` ts.TTVE(expected ...interfaces{})```.
* Check expected results with nil error ```ts.ToTestValuesExpectedWithNilError(expected ...interfaces{})``` 
or ``` ts.TTVEWNE(expected ...interfaces{})```.
* Check expected errors (expected error as last returned parameters) .
```ts.ToTestErrorExpected(err error)``` or ``` ts.TTEE(err error)```
* Check panic error ```ts.ToTestPanicErrorExpected(msg interface{})``` or ``` ts.TTPEE(msg interface{})```.
* Check not nil error ```ts.ToTestCheckErrorExpected()``` or ``` ts.TTCEE()```.

## Examples

```go
package test

import (
	"testing"
	"github.com/ThCompiler/ts"
)

func sum(a, b int64) int64 {
	return a + b
}

type SumSuite struct {
	ts.TestCasesSuite
	ActFunc func(args ...interface{}) []interface{}
}

func (s *SumSuite) SetupTest() {
	s.ActFunc = func(args ...interface{}) []interface{} {
		res := sum(args[0].(int64), args[1].(int64))
		return []interface{}{res}
	}
}

func (s *SumSuite) TestCorrectNumber() {
	s.RunTest(
		s.ActFunc,
		ts.TestCase{
			Name:     "Zero With One",
			Args:     ts.TTA(int64(0), int64(1)),
			Expected: ts.TTVE(int64(1)),
		},
		ts.TestCase{
			Name:     "Two With One",
			Args:     ts.TTA(int64(2), int64(1)),
			Expected: ts.TTVE(int64(3)),
		},
		ts.TestCase{
			Name:     "Ten With One",
			Args:     ts.TTA(int64(10), int64(1)),
			Expected: ts.TTVE(int64(11)),
		},
		ts.TestCase{
			Name:     "One With One",
			Args:     ts.TTA(int64(1), int64(1)),
			Expected: ts.TTVE(int64(2)),
		},
	)
}

func TestSumSuite(t *testing.T) {
	suite.Run(t, new(SumSuite))
}
```

This example is located in the `example` folder.

## Future

- [ ] Add documentation.
- [ ] Add Some feature to work with gmock (now, i has no idea how it will be work and be comfortable for using).
- [ ] Improve the way to set the function under test.
