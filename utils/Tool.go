package utils
import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"math/big"
	"net"
	_"reflect"
	_"fmt"
)

func Md5(str string)  string {	
	h := md5.New()
	h.Write([]byte(str))
	md5Str1 := h.Sum(nil)
	return hex.EncodeToString(md5Str1)
}


func Ip2Long(ipStr string) (int64,error){
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0,errors.New("Illegal IP")
	}
	ip = ip.To4()
	ret := big.NewInt(0)
	ret.SetBytes(ip)
	return ret.Int64(),nil
} 

func Long2IP (intIp int64) (string,error){
	var bytes [4]byte
	bytes[0] = byte(intIp & 0xFF)
	bytes[1] = byte((intIp >> 8) & 0xFF)
	bytes[2] = byte((intIp >> 16) & 0xFF)
    bytes[3] = byte((intIp >> 24) & 0xFF)
	ip := net.IPv4(bytes[3],bytes[2],bytes[1],bytes[0])
	return ip.String(),nil
}



