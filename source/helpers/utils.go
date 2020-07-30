package helpers

import (
	"encoding/hex"
	"strconv"
	"strings"
	"unsafe"
)

// bytes array to hex string
func B2hex(b []byte) string {
	// here I was playing around with different ways os converting bytes array to a hex string
	//dst := make([]byte, hex.EncodedLen(len(b)))
	//hex.Encode(dst, b)
	//return fmt.Sprintf("%s", dst)
	//if we got the string with 0x prefix we will cut it
	res := strings.Replace(hex.EncodeToString(b[:]), "0x", "", -1)
	return res
}

// converter to an int64 from the hex string
func Hex2int(hexStr string) uint64 {
	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(hexStr, 16, 64)
	return uint64(result)

}

//ByteArrayToInt Byte by byte copy of given array to an int64 variable
func ByteArrayToInt(arr []byte) int64 {
	val := int64(0)
	for i := 0; i < 8; i++ { //int64 has 8 bytes
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}
	return val
}
