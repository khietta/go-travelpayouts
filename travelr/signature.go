package travelr

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
)

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
