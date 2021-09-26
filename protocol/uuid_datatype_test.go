package protocol

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testUUIDJSON = `{"uid":"3ed1e2e9-d7b7-40c0-81d8-00a3150d9da7"}`

func Test_EncodingDecoding(t *testing.T) {
	var testStruct = struct {
		UID UUID `json:"uid"`
	}{}

	err := json.Unmarshal([]byte(testUUIDJSON), &testStruct)
	if !assert.NoError(t, err, "Unmarshal UUID") {
		return
	}

	res, err := json.Marshal(&testStruct)
	if !assert.NoError(t, err, "Marshal UUID") {
		return
	}

	if !bytes.Equal(res, []byte(testUUIDJSON)) {
		t.Error("Invalid data marshal-unmarshal")
	}
}
