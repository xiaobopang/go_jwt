package libs

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numLetterBytes = "0123456789"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func Md5(str string) string {
	if str == "" {
		return ""
	}
	init := md5.New()
	init.Write([]byte(str))
	return fmt.Sprintf("%x", init.Sum(nil))
}

func Sha256(s string) string {
	if s == "" {
		return ""
	}
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Base64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

// 字符串转换整型
func IpString2Int(ipstring string) int64 {
	ipSegs := strings.Split(ipstring, ".")
	ipInt := 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return int64(ipInt)
}

// 整型转换成字符串
func IpInt2String(ipInt int) string {
	ipSegs := make([]string, 4)
	var len int = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < len; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[len-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < len; i++ {
		buffer.WriteString(ipSegs[i])
		if i < len-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}

func randString(n int, LetterBytes string) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(LetterBytes) {
			b[i] = LetterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func RandString(n int) string {
	return randString(n, letterBytes)
}

func RandNumString(n int) string {
	return randString(n, numLetterBytes)
}

var Rander = rand.New(src)

// 随机数生成
// @Param	min 	int	最小值
// @Param 	max		int	最大值
// @return  int
func RandInt(min int, max int) int {
	num := Rander.Intn((max - min)) + min
	return num
}

// 获取服务器IP
func GetLocalIp() string {
	addrSlice, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Get local IP addr failed!")
		return "127.0.0.1"
	}
	for _, addr := range addrSlice {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}
