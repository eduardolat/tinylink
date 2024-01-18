package validatorutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name string `validate:"required"`
	Age  int    `validate:"gte=0"`
}

func TestPrettyValidate(t *testing.T) {
	// Create a struct with invalid values
	s := TestStruct{Name: "", Age: -1}

	// Call PrettyValidate with the invalid struct
	err := PrettyValidate(s)

	// Assert that an error was returned
	assert.Error(t, err)

	// Assert that the error message is as expected
	expectedErrMsg := "Name: required, Age: gte"
	assert.Equal(t, expectedErrMsg, err.Error())

	// Create a struct with valid values
	s = TestStruct{Name: "Test", Age: 20}

	// Call PrettyValidate with the valid struct
	err = PrettyValidate(s)

	// Assert that no error was returned
	assert.NoError(t, err)
}
