package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"strings"
)

func GetUUID() string {
	uuid1, _ := uuid.NewUUID()
	salt := uuid1.String()
	return strings.Replace(salt, "-", "", -1)
}

func GetMD5(file io.Reader) string {
	var md5str1 string
	if buffer, err := ioutil.ReadAll(file); err == nil {
		md := md5.Sum(buffer)
		md5str1 = fmt.Sprintf("%x", md)
	} else {
		fmt.Println(buffer, err)
	}
	return md5str1
}