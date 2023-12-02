package ts

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// MockTestFunction is a interface of function to initialize mocks for a test case
type MockTestFunction func(ctrl *gomock.Controller) []interface{}

// TestCase is a structure describing the test case
type TestCase struct {
	Name      string           // Name of test
	Args      []interface{}    // Args of testing function
	Expected  TestExpected     // Info about expected return of testing function
	InitMocks MockTestFunction // For Mock initialization !!!future
}

// TestCasesSuite is a wrapper over a suite.Suite with added methods for run test cases.
type TestCasesSuite struct {
	suite.Suite
}

// RunTest is a function to run a list of test cases for the tested function
func (s *TestCasesSuite) RunTest(fun any, cases ...TestCase) {
	RunTest(s.T(), fun, cases...)
}

func assertCase(t *testing.T, res []interface{}, expected TestExpected, caseName string) {
	t.Helper()

	if expected.HaveError() {
		assertWithError(t, &res, expected, caseName)
	} else {
		require.Equalf(t, len(expected.ExpectedReturns), len(res),
			"Testcase with name: %s, different len of expected and gotten return values", caseName)
	}

	for i, expected := range expected.ExpectedReturns {
		assert.EqualValuesf(t, expected, res[i], "Testcase with name: %s", caseName)
	}
}

func assertWithError(t *testing.T, res *[]interface{}, expected TestExpected, caseName string) {
	t.Helper()

	size := len(*res)
	require.NotZerof(t, size, "Testcase with name: %s, return nothing, but wait return error", caseName)
	gottenError, ok := (*res)[size-1].(error)

	if !ok && gottenError != nil {
		require.Failf(t, "Last value not error, but expected error",
			"Testcase with name: %s", caseName)
	}

	if expected.MustErrorExpected().CheckError {
		assert.Error(t, gottenError, "Testcase with name: %s", caseName)
	} else {
		checkForCorrectnessError(t, gottenError, expected, caseName)
	}

	*res = (*res)[:size-1]
}

func checkForCorrectnessError(t *testing.T, gottenError error, expected TestExpected, caseName string) {
	t.Helper()

	if expected.MustErrorExpected().Error == nil {
		assert.NoError(t, gottenError,
			"Testcase with name: %s", caseName)
	} else {
		assert.ErrorIs(t, gottenError, expected.MustErrorExpected().Error,
			"Testcase with name: %s", caseName)
	}
}

func checkForCorrectnessPanicError(t *testing.T, msg any, expected TestExpected, caseName string) {
	t.Helper()

	if expected.HavePanicError() {
		assert.EqualValuesf(t, expected.MustPanicErrorExpected().Msg, msg, "Testcase with name: %s", caseName)

		return
	}

	if err, is := msg.(error); is {
		assert.Failf(t, "Panic error testcase: ", "%s %s", caseName, err)
	}

	assert.Failf(t, "Panic error testcase: ", "%s %v", caseName, msg)
}

func runTestCase(t *testing.T, test TestCase, fun any, ctrl *gomock.Controller) {
	t.Helper()

	defer func(t *testing.T, test TestCase) {
		t.Helper()

		if r := recover(); r != nil {
			checkForCorrectnessPanicError(t, r, test.Expected, test.Name)
		}
	}(t, test)

	args := test.Args

	if test.InitMocks != nil {
		mocks := test.InitMocks(ctrl)
		args = append(mocks, args...)
	}

	res, err := runFunction(args, fun)
	assert.NoError(t, err, "Catch error when bind args to testing function")

	assertCase(t, res, test.Expected, test.Name)
}

// RunTest is a function to run test cases for testing function out of TestCasesSuite
func RunTest(t *testing.T, fun any, cases ...TestCase) {
	t.Helper()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, cs := range cases {
		cs := cs
		t.Run(cs.Name, func(t *testing.T) {
			runTestCase(t, cs, fun, ctrl)
		})
	}
}

// ToTestArgs is a small function to convert a list of args to an array of []interfaces{}
// It can be easier write ToTestArgs(a, b, c) then []interface{}{a, b, c}.
func ToTestArgs(args ...interface{}) []interface{} {
	return args
}

// TTA is a wrapper over a function ToTestArgs to shorten its name
func TTA(args ...interface{}) []interface{} {
	return ToTestArgs(args...)
}

func runFunction(args []interface{}, fun any) ([]interface{}, error) {
	funValue := reflect.ValueOf(fun)
	funType := funValue.Type()

	if funType.Kind() != reflect.Func {
		return nil, fmt.Errorf("testing function is not function type it is %s", funType.String())
	}

	numIn := funType.NumIn()
	if numIn > len(args) {
		return nil, fmt.Errorf("testing function must have minimum %d params. Have %d", numIn, len(args))
	}
	if numIn != len(args) && !funType.IsVariadic() {
		return nil, fmt.Errorf("testing function must have %d params. Have %d", numIn, len(args))
	}

	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		var inType reflect.Type
		if funType.IsVariadic() && i >= numIn-1 {
			inType = funType.In(numIn - 1).Elem()
		} else {
			inType = funType.In(i)
		}

		argValue := reflect.ValueOf(args[i])
		if !argValue.IsValid() {
			return nil, fmt.Errorf("param[%d] of testing function must be %s. Have %s", i, inType, argValue.String())
		}

		argType := argValue.Type()
		if argType.ConvertibleTo(inType) {
			in[i] = argValue.Convert(inType)
		} else {
			return nil, fmt.Errorf("param[%d] of testing function must be %s. Have %s", i, inType, argType)
		}
	}

	funcRes := funValue.Call(in)
	res := make([]interface{}, len(funcRes))

	for i, rs := range funcRes {
		res[i] = rs.Interface()
	}

	return res, nil
}
