
package util
import (
 	"crypto/rand"
    "unsafe"

)
// TableName:   Customer/
var alphabet = []byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

 
func Generate() string {
	size:=20
    b := make([]byte, size)
    rand.Read(b)
    for i := 0; i < size; i++ {
        b[i] = alphabet[b[i] / 5]
    }
    return *(*string)(unsafe.Pointer(&b))
}

