package utils
import (
	"crypto/md5"
	"encoding/hex"
	_"reflect"
	_"fmt"
)

func Md5(str string)  string {	
	h := md5.New()
	h.Write([]byte(str))
	md5Str1 := h.Sum(nil)
	return hex.EncodeToString(md5Str1)
}





