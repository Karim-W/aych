//
//  utils.go
//  aych
//
//  Created by karim-w on 10/07/2025.
//

package aych

import (
	"encoding/base64"
	"strings"
)

func string_builder(strs ...string) string {
	builder := strings.Builder{}
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}

func base64_encode(decoded string) string {
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(decoded)))
	base64.StdEncoding.Encode(encoded, []byte(decoded))
	return string(encoded)
}
