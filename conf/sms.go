package conf

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/kataras/iris"
)


var client *dysmsapi.Client


func loadClient(c iris.Configuration) {
	client, _ = dysmsapi.NewClientWithAccessKey("cn-hangzhou", c.Other["AccessKeyID"].(string),
		c.Other["AccessKeySecret"].(string))
}

func SendSMS(phoneNumber string, code int64){
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phoneNumber
	request.SignName = "李璞玉"
	request.TemplateCode = "SMS_139580033"
	request.TemplateParam = fmt.Sprintf("{\"code\": %d}", code)
	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
