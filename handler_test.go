package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestToken(t *testing.T) {
	body := "123"

	h := handler{
		stats: make(map[string]uint64),
		key:   []byte("some-baked-in-secret"),
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", strings.NewReader(body))

	h.token(rec, req)

	mac := createMAC([]byte(body), h.key)
	actual, _ := hex.DecodeString(rec.Body.String())

	if !hmac.Equal(actual, mac) {
		t.Errorf("failed to validate hmac")
	}

}

func validMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
