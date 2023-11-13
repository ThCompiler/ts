package example

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"ts"
)

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
