// Author: sheppard(ysf1026@gmail.com) 2014-01-06

package proto

import (
	"encoding/json"
)

func Encode(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(b)
}
