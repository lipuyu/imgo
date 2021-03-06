package service

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"os"
	"stouch_server/conf"
	"strings"
)

var bucket *oss.Bucket

func init() {
	// 创建OSSClient实例。
	c := conf.Config
	client, err := oss.New(c.Other["EndPoint"].(string), c.Other["AccessKeyID"].(string), c.Other["AccessKeySecret"].(string))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err = client.Bucket("lipuyu")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 上传字符串。
	err = bucket.PutObject("test.txt", strings.NewReader("hello world!"))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func GetOrSave(name string, reader io.Reader) bool {
	isExist, err := bucket.IsObjectExist(name)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if !isExist {
		if err := bucket.PutObject(name, reader); err != nil {
			conf.Logger.Error("Error:", err)
		}
	}
	return isExist
}
