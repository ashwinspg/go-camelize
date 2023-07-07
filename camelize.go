package camelize

import (
	"encoding/json"
	"regexp"
	"strings"
)

func transformKey(key string) string {
	var re = regexp.MustCompile(`[_.-](\w|$)`)

	transformedByte := re.ReplaceAllFunc(
		[]byte(key),
		func(s []byte) []byte {
			return []byte(strings.ToUpper(string(s[1])))
		},
	)

	return string(transformedByte)
}

func transformKeys(dataMap any) (any, error) {
	var err error

	switch dataMapVal := dataMap.(type) {
	case map[string]any:
		for k, v := range dataMapVal {
			transformedKey := transformKey(k)
			delete(dataMapVal, k)

			if dataMapVal[transformedKey], err = transformKeys(v); err != nil {
				return nil, err
			}
		}

		return dataMapVal, nil
	case []any:
		transformedList := make([]any, len(dataMapVal))

		for i, v := range dataMapVal {
			if transformedList[i], err = transformKeys(v); err != nil {
				return nil, err
			}
		}

		return transformedList, nil
	}

	return dataMap, nil
}

func TransformJSONKeys(data []byte) ([]byte, error) {
	m := make(map[string]any)

	if err := json.Unmarshal([]byte(data), &m); err != nil {
		// Not a JSON object
		return nil, err
	}

	out, err := transformKeys(m)

	if err != nil {
		return data, err
	}

	return json.Marshal(out)
}
