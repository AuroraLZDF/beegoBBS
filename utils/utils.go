package utils

import (
	"github.com/astaxie/beego"
	"math/rand"
	"fmt"
	"time"
	"crypto/md5"
	"encoding/hex"
	"os"
	"encoding/base64"
	"encoding/json"
	"net/smtp"
	"strings"
)

var cfg = beego.AppConfig

func Configs() interface{} {
	var Config = map[string]string{
		"lang":  beego.AppConfig.String("local"),
		"title": "goLang 学习教程",
		"email": "18862324237@163.com",
	}

	return Config
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
	result := fmt.Sprintf("%s", hex.EncodeToString(cipherStr)) // 输出加密结果

	return result
}

/**
 * str : json string
 */
func AuthCode(str string, flag string) string {
	key := beego.AppConfig.String("app_secret")

	if flag == "DECODE" {
		decoded, _ := base64.StdEncoding.DecodeString(str)
		decodeStr := string(decoded)

		return decodeStr
	}

	res := JsonToMap(str)
	res["sign"] = Md5(str + key)
	_str := MapToJson(res)

	strBytes := []byte(_str)
	encodedStr := base64.StdEncoding.EncodeToString(strBytes)

	return encodedStr
}

/**
 * map 转 json
 */
func MapToJson(m map[string]interface{}) string {
	js, err := json.Marshal(m)

	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return ""
	}

	return string(js)
}

/**
 * json 转 map
 */
func JsonToMap(jsonStr string) map[string]interface{} {
	var mapResult map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &mapResult); err != nil {
		fmt.Println("json.Unmarshal failed:", err)
	}

	return mapResult
}

/**
 * 发送邮件
 */
func SendMail(to, subject, body, mailType string) error {
	user := cfg.String("mail_username")
	password := cfg.String("mail_password")
	host := cfg.String("mail_host")
	port := cfg.String("mail_port")

	auth := smtp.PlainAuth("", user, password, host)

	var contentType string
	if mailType == "html" {
		contentType = "Content-Type: text/html" + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host + ":" + port, auth, user, sendTo, msg)

	return err
}

// 去除空格
func TrimS(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}
