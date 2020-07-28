package controller

import (
	h "../helpers"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
)

type Hash struct {
	Phrase       string `json:"phrase"`
	HashedPhrase int64
}
type DetailedHash struct {
	Phrase       string `json:"phrase"`
	HexedPhrase  string
	HexedCuttedPhrase string
	HashedPhrase int64
}

var MainMethod = func(w http.ResponseWriter, r *http.Request) {
	var hash Hash
	var tmpVal map[string]interface{} // проместим входящую переменную в тип мап(ключ: значение), нужный нам ключ мы знаем
	json.NewDecoder(r.Body).Decode(&tmpVal)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	hash.Phrase = fmt.Sprintf("%v", tmpVal["phrase"]) //если на вход нам дали число(float64) - парсим его в строку
	if hash.Phrase == "<nil>" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Header().Add("Content-Type", "application/json")
		encoder.Encode(w)
	} else {
		sha256bytes := sha256.Sum256([]byte(hash.Phrase)) //получаем хэш в виде массива байт
		hash.HashedPhrase = h.ByteArrayToInt(sha256bytes[:8])
		encoder.Encode(&hash)
	}
}
var DetailedMethod = func(w http.ResponseWriter, r *http.Request) {
	var detailedHash DetailedHash
	var tmpVal map[string]interface{} // проместим входящую переменную в тип мап(ключ: значение), нужный нам ключ мы знаем
	json.NewDecoder(r.Body).Decode(&tmpVal)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	detailedHash.Phrase = fmt.Sprintf("%v", tmpVal["phrase"]) //если на вход нам дали число(float64) - парсим его в строку
	if detailedHash.Phrase == "<nil>" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Header().Add("Content-Type", "application/json")
		encoder.Encode(w)
	} else {
		sha256bytes := sha256.Sum256([]byte(detailedHash.Phrase)) //получаем хэш в виде массива байт
		detailedHash.HexedPhrase = h.B2hex(sha256bytes[:])       //хексим слайс массива
		detailedHash.HexedCuttedPhrase = h.B2hex(sha256bytes[:8])
		detailedHash.HashedPhrase = h.ByteArrayToInt(sha256bytes[:8])
		encoder.Encode(&detailedHash)
	}
}