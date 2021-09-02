package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type handler struct {
	key   []byte
	stats map[string]uint64
}

func (h handler) health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h handler) token(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	h.stats["requests"] += 1

	body, _ := io.ReadAll(r.Body)
	out := createMAC(body, h.key)
	fmt.Fprintf(w, "%x", out)

	w.WriteHeader(201)
}

func (h handler) metrics(w http.ResponseWriter, r *http.Request) {

	enc := json.NewEncoder(w)
	enc.Encode(h.stats)
	w.WriteHeader(201)
}

func createMAC(message, key []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}
