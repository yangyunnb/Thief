package common

import (
	"bytes"
	"encoding/json"
)

// JSONDecode json解析
func JSONDecode(input []byte, target interface{}) error {
	jsonDecoder := json.NewDecoder(bytes.NewBuffer(input))
	jsonDecoder.UseNumber()
	err := jsonDecoder.Decode(target)
	if err != nil {
		return err
	}
	return nil
}

func JSONEncode(v interface{}) string {
	param := bytes.NewBuffer(nil)
	j, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	if _, err := param.Write(j); err != nil {
		return ""
	}
	return param.String()
}

func LoadFormMapOrStruct(dst interface{}, src ...interface{}) error {
	for _, data := range src {
		bytes, err := json.Marshal(data)
		if err != nil {
			return err
		}
		err = json.Unmarshal(bytes, dst)
		if err != nil {
			return err
		}
	}

	return nil
}
