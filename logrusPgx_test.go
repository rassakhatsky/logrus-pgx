// Package logrus_pgx provides ability to use Logrus with PGX
package logrus_pgx

import (
	"bytes"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"strings"
	"testing"
)

func TestDebug(t *testing.T) {
	var outbuf bytes.Buffer
	testSeq := []string{"info", "debug", "warning", "error"}

	logger := (*PgxLogger)(log.StandardLogger())
	logger.Out = &outbuf
	logger.Formatter = &log.JSONFormatter{}
	logger.Level = log.DebugLevel

	testText := "A group of walrus emerges from the ocean"
	boolKey := "key_bool_"
	intKey := "key_int_"
	strKey := "key_str_"
	strValue := "test_string_"

	logger.Info(testText,
		strKey+testSeq[0], strValue+testSeq[0],
		intKey+testSeq[0], 1,
		boolKey+testSeq[0], true)
	logger.Debug(testText,
		strKey+testSeq[1], strValue+testSeq[1],
		intKey+testSeq[1], 2,
		boolKey+testSeq[1], true)
	logger.Warn(testText,
		strKey+testSeq[2], strValue+testSeq[2],
		intKey+testSeq[2], 3,
		boolKey+testSeq[2], true)
	logger.Error(testText,
		strKey+testSeq[3], strValue+testSeq[3],
		intKey+testSeq[3], 4,
		boolKey+testSeq[3], true)

	messages := strings.Split(outbuf.String(), "\n")
	if len(messages) != 5 {
		t.Errorf("%d messages not expected.", len(messages))
	}

	// Last message is an empty line – have to delete it
	if messages[len(messages)-1] != "" {
		t.Errorf("The last line is wrong: %v.", messages[:len(messages)-1])
	}
	messages = append(messages[:len(messages)-1], messages[len(messages):]...)
	for i, test := range testSeq {
		var body map[string]interface{}
		if json.Unmarshal([]byte(messages[i]), &body) != nil {
			t.Errorf("Wrong format of JSON file: %v.", messages[i])
		}
		if body["msg"] != testText {
			t.Errorf("Unexpected message: %v", body["msg"])
		}
		if body["level"] != test {
			t.Errorf("Unexpected log type: %v", body["level"])
		}

		if body[boolKey+test] != true {
			t.Errorf("Unexpected bool key: %b", body[boolKey+test])
		}
		if body[intKey+test] != float64(i+1) {
			t.Errorf("Unexpected integer key: %d", body[intKey+test])
		}
		if body[strKey+test] != strValue+test {
			t.Errorf("Unexpected string key: %v", body[strKey+test])
		}
	}
}
