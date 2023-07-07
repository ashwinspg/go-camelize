package camelize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformKey(t *testing.T) {
	testCases := []struct {
		key            string
		expectedOutput string
	}{
		{
			key:            "p_qq_rr",
			expectedOutput: "pQqRr",
		},
		{
			key:            "i_jj_kkL",
			expectedOutput: "iJjKkL",
		},
		{
			key:            "pqr_stu_sdf",
			expectedOutput: "pqrStuSdf",
		},
		{
			key:            "pqr_stu",
			expectedOutput: "pqrStu",
		},
		{
			key:            "abc.xyz",
			expectedOutput: "abcXyz",
		},
		{
			key:            "foo-bar",
			expectedOutput: "fooBar",
		},
	}

	for _, testCase := range testCases {
		output := transformKey(testCase.key)

		assertObj := assert.New(t)

		assertObj.Equal(testCase.expectedOutput, output)
	}

}
