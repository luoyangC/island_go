package utils

import (
	"encoding/json"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"

	"island/cache"
)

func SendSmsCode(mobile string, t int) string {
	code := MakeCode()

	accessKeyId := os.Getenv("ALI_ACCESS_KEYID")
	accessSecret := os.Getenv("ALI_ACCESS_SECRET")
	client, _ := dysmsapi.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessSecret)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = "island"
	request.TemplateCode = "SMS_178695047"

	request.PhoneNumbers = mobile
	request.TemplateParam = MakeParams(code)
	response, _ := client.SendSms(request)
	if response.Code == "OK" {
		CacheCode(mobile, t, code)
	}
	return response.Message
}

func MakeCode() string {
	var letterRunes = []rune("1234567890")
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, 6)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func MakeParams(code string) string {
	p, _:= json.Marshal(map[string]string{
		"code": code,
	})
	return string(p)
}

func CacheCode(mobile string, t int, code string) {
	val := strings.Join([]string{strconv.Itoa(t), code}, ",")
	if err := cache.Client.Set(mobile, val, time.Minute).Err(); err != nil {
		panic(err)
	}
}