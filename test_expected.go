package ts

// TestPanicErrorExpected is a struct contained info about the expected panic message of tested function
type TestPanicErrorExpected struct {
	Msg interface{} // expected message
}

// TestErrorExpected is a struct contained info about expected returns of tested function
type TestErrorExpected struct {
	CheckError bool  // indication that it is only necessary to check for an error, without checking for a match with the expected one
	Error      error // expected error
}

// TestExpected is a struct contained info about expected returns of tested function
type TestExpected struct {
	PanicError      *TestPanicErrorExpected // expected panic message of tested function
	Error           *TestErrorExpected      // expected error of tested function
	ExpectedReturns []interface{}           // expected returns of tested function
}

// HaveError is a function checking if an error is expected
func (te *TestExpected) HaveError() bool {
	return te.Error != nil
}

// HavePanicError is a function checking if a panic is expected
func (te *TestExpected) HavePanicError() bool {
	return te.PanicError != nil
}

// MustErrorExpected is a function returning expected error. If it does not exist, function panic with error
func (te *TestExpected) MustErrorExpected() TestErrorExpected {
	if te.Error == nil {
		panic("Expected error, but not")
	}

	return *te.Error
}

// MustPanicErrorExpected is a function returning expected panic error. If it does not exist, function panic with error
func (te *TestExpected) MustPanicErrorExpected() TestPanicErrorExpected {
	if te.PanicError == nil {
		panic("Expected panic error, but not")
	}

	return *te.PanicError
}

// ToTestValuesExpected is a function to construct test exception with only waiting for the specified arguments
// and no one error as the result of the function under test
func ToTestValuesExpected(expedites ...interface{}) TestExpected {
	return TestExpected{
		PanicError:      nil,
		Error:           nil,
		ExpectedReturns: expedites,
	}
}

// TTVE is a wrapper over a function ToTestValuesExpected to shorten its name
func TTVE(expedites ...interface{}) TestExpected {
	return ToTestValuesExpected(expedites...)
}

// ToTestValuesExpectedWithNilError is a function to construct test exception with only waiting for the specified arguments
// // and nil error as the result of the function under test
func ToTestValuesExpectedWithNilError(expedites ...interface{}) TestExpected {
	return TestExpected{
		PanicError: nil,
		Error: &TestErrorExpected{
			Error: nil,
		},
		ExpectedReturns: expedites,
	}
}

// TTVEWNE is a wrapper over a function ToTestValuesExpectedWithNilError to shorten its name
func TTVEWNE(expedites ...interface{}) TestExpected {
	return ToTestValuesExpectedWithNilError(expedites...)
}

// ToTestErrorExpected is a function to construct test exception with only waiting for the specified error
// as the result of the function under test
func ToTestErrorExpected(err error) TestExpected {
	return TestExpected{
		PanicError: nil,
		Error: &TestErrorExpected{
			Error: err,
		},
		ExpectedReturns: nil,
	}
}

// TTEE is a wrapper over a function ToTestErrorExpected to shorten its name
func TTEE(err error) TestExpected {
	return ToTestErrorExpected(err)
}

// ToTestCheckErrorExpected is a function to construct test exception with only waiting for the not specified error
// as the result of the function under test
func ToTestCheckErrorExpected() TestExpected {
	return TestExpected{
		PanicError: nil,
		Error: &TestErrorExpected{
			CheckError: true,
		},
		ExpectedReturns: nil,
	}
}

// TTCEE is a wrapper over a function ToTestCheckErrorExpected to shorten its name
func TTCEE() TestExpected {
	return ToTestCheckErrorExpected()
}

// ToTestPanicErrorExpected is a function to construct test exception with only waiting for the panic message
// as the result of the function under test
func ToTestPanicErrorExpected(msg interface{}) TestExpected {
	return TestExpected{
		PanicError: &TestPanicErrorExpected{
			Msg: msg,
		},
		Error:           nil,
		ExpectedReturns: nil,
	}
}

// TTPEE is a wrapper over a function ToTestPanicErrorExpected to shorten its name
func TTPEE(msg interface{}) TestExpected {
	return ToTestPanicErrorExpected(msg)
}

// ToTestExpected is a function to construct test exception all variant of waiting results of the function under test
func ToTestExpected(checkError bool, err error, withPanic bool, panicMsg interface{},
	expedites ...interface{},
) TestExpected {
	if withPanic {
		return ToTestPanicErrorExpected(panicMsg)
	}

	if !checkError && err != nil {
		return ToTestValuesExpected(expedites...)
	}

	if checkError {
		return ToTestCheckErrorExpected()
	}

	return ToTestErrorExpected(err)
}

// TTE is a wrapper over a function ToTestExpected to shorten its name
func TTE(checkError bool, err error, withPanic bool, panicMsg string,
	expedites ...interface{},
) TestExpected {
	return ToTestExpected(checkError, err, withPanic, panicMsg, expedites...)
}
