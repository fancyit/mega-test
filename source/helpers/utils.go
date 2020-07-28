package helpers

import (
	"encoding/hex"
	"strconv"
	"strings"
)

// массив байт в 16-ю строку
func B2hex(b []byte) string {
	// возвращаем шеснадцатеричную строку, убрав в ней префикс 0x, если есть
	//dst := make([]byte, hex.EncodedLen(len(b)))
	//hex.Encode(dst, b)
	//return fmt.Sprintf("%s", dst)
	res := strings.Replace(hex.EncodeToString(b[:]), "0x", "", -1)
	return res
}

// 16-я строка в uint64
func Hex2int(hexStr string) uint64 {
	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(hexStr, 16, 64)
	return uint64(result)
}
