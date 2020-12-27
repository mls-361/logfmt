/*
------------------------------------------------------------------------------------------------------------------------
####### logfmt ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package logfmt

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/mls-361/buffer"
)

const (
	_oddList = "--odd-list--"
	_badKey  = "--bad-key--"
)

func cleanKey(r rune) rune {
	if r <= ' ' || r == '=' || r == '"' || r == utf8.RuneError {
		return '?'
	}

	return r
}

// Encode AFAIRE.
func Encode(buf *buffer.Buffer, kv map[string]interface{}) {
	addSpace := false

	for key, value := range kv {
		if addSpace {
			buf.AppendByte(' ')
		} else {
			addSpace = true
		}

		// key
		buf.AppendString(strings.Map(cleanKey, key))
		// =
		buf.AppendByte('=')
		// value
		buf.AppendString(fmt.Sprintf("%#v", value))
	}
}

// EncodeList AFAIRE.
func EncodeList(buf *buffer.Buffer, kv ...interface{}) {
	if len(kv)%2 == 1 {
		buf.AppendString(_oddList)
		return
	}

	addSpace := false

	for i := 0; i < len(kv); i += 2 {
		key, value := kv[i], kv[i+1]

		if addSpace {
			buf.AppendByte(' ')
		} else {
			addSpace = true
		}

		// key
		k, ok := key.(string)
		if ok {
			buf.AppendString(strings.Map(cleanKey, k))
		} else {
			buf.AppendString(_badKey)
		}
		// =
		buf.AppendByte('=')
		// value
		buf.AppendString(fmt.Sprintf("%#v", value))
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
