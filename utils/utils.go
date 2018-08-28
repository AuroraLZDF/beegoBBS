package utils

import (
	"github.com/astaxie/beego"
	"math/rand"
	"fmt"
	"time"
	"crypto/md5"
	"encoding/hex"
	"os"
)

func Configs() interface{} {
	var Config = map[string]string{
		"lang":  beego.AppConfig.String("local"),
		"title": "goLang 学习教程",
		"email": "18862324237@163.com",
	}

	return Config
}

// 检测登录状态
func CheckCk() bool {
	/*var ctx *context.Context
	_, ok := ctx.Input.Session("uid").(int)

	if !ok && ctx.Request.RequestURI != "/login" {
		//ctx.Redirect(302, "/login")
		return false
	}*/

	return false
	return true
}

func Csrf_token() string {
	return GetRandomString(40)
}

func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	result := make([]byte, length)
	for i := range result {
		result[i] = str[rand.Intn(len(str))]
	}
	fmt.Println("1111", result, string(result))
	return string(result)
}

func ShowErr(err interface{}) {
	fmt.Println("ERROR | ", time.Now().Format("2006-01-02 15:04:05"), " | ", err)
	os.Exit(500)
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串
	cipherStr := h.Sum(nil)
	fmt.Println(cipherStr)
	result := fmt.Sprintf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果

	return result
}
