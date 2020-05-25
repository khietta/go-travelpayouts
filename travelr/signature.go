package travelr

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"time"
)

// convertStructToMap ...
func convertStructToMap(st interface{}) map[string]interface{} {
	if reflect.TypeOf(st).Kind() != reflect.Struct {
		return nil
	}

	rp := make(map[string]interface{}, 0)

	tf := reflect.TypeOf(st)
	vf := reflect.ValueOf(st)

	for i := 0; i < vf.NumField(); i++ {
		if !vf.Field(i).IsNil() {
			key := tf.Field(i).Tag.Get("json")
			if key != "" {
				if rs, ok := vf.Field(i).Elem().Interface().(Passengers); ok {
					rp[key] = convertStructToMap(rs)
				} else if rs, ok := vf.Field(i).Elem().Interface().([]Segment); ok {
					en := make([]interface{}, 0)
					for _, v := range rs {
						en = append(en, convertStructToMap(v))
					}
					rp[key] = en
				} else {
					switch dt := vf.Field(i).Elem().Interface().(type) {
					case time.Time:
						rp[key] = dt.Format(timeFormat)
					default:
						rp[key] = dt
					}
				}
			}
		}
	}

	return rp
}

// parse signature from query params
func makeSignatureString(params interface{}) string {
	var sigStr string

	if rs, ok := params.([]interface{}); ok {
		for _, param := range rs {
			sigStr += makeSignatureString(param)
		}
	} else if rs, ok := params.(map[string]interface{}); ok {
		keys := make([]string, 0, len(rs))
		for k := range rs {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, param := range keys {
			sigStr += makeSignatureString(rs[param])
		}
	} else {
		sigStr = fmt.Sprintf(":%v", params)
	}

	return sigStr
}

// parse signature with token
func MakeSignatureToken(token string, params map[string]interface{}) (string, error) {
	if params == nil {
		return "", invalidParams
	}

	sig := makeSignatureString(params)

	hasher := md5.New()
	hasher.Write([]byte(token + sig))

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
