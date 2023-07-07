package camelize_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ashwinspg/go-camelize"
)

func TestTransformJSONKeys(t *testing.T) {
	testCases := []struct {
		data           string
		expectedOutput string
	}{
		{
			data: `{
				"fee_fie_foe": "fum",
				"beep_boop": [
					{ "abc.xyz": "mno" },
					{ "foo-bar": "baz" }
				]
			}`,
			expectedOutput: `{
				"feeFieFoe":"fum",
				"beepBoop":[
					{ "abcXyz": "mno" },
					{ "fooBar": "baz"}
				]
			}`,
		},
		{
			data: `{
				"a_b": "cc_cc",
				"d_e": ["d"],
				"e_f": {
					"g_h": {
						"i_j": "k",
						"l_m": {},
						"n_o": true,
						"p_q_r": null
					}
				}
			}`,
			expectedOutput: `{
				"aB": "cc_cc",
				"dE": ["d"],
				"eF": {
					"gH": {
						"iJ": "k",
						"lM": {},
						"nO": true,
						"pQR": null
					}
				}
			}`,
		},
		{
			data: `{
				"a_b": [
					{
						"c_d": "cc_dd"
					},
					{
						"e_f": 123,
						"g_h": false
					}
				],
				"i_jj_kkL": "m_n_o"
			}`,
			expectedOutput: `{
				"aB": [
					{
						"cD": "cc_dd"
					},
					{
						"eF": 123,
						"gH":false
					}
				],
				"iJjKkL": "m_n_o"
			}`,
		},
	}

	for _, testCase := range testCases {
		output, err := camelize.TransformJSONKeys([]byte(testCase.data))

		assertObj := assert.New(t)

		assertObj.NoError(err)

		var expectedOutputMap map[string]any
		err = json.Unmarshal([]byte(testCase.expectedOutput), &expectedOutputMap)
		assertObj.NoError(err)

		var outputMap map[string]any
		err = json.Unmarshal(output, &outputMap)
		assertObj.NoError(err)

		assertObj.True(reflect.DeepEqual(expectedOutputMap, outputMap))
	}
}
