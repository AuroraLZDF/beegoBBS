package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/smtp"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"net/http"
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
 * 获取变量类型
 */
func TypeOf(v interface{}) string {
	return reflect.TypeOf(v).String()
}

/**
 * map to json
 */
func MapToJson(obj map[string]interface{}) string {
	jsonBytes, err := json.Marshal(obj)

	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return ""
	}

	return string(jsonBytes)
}

/**
 * json to map
 */
func JsonToMap(jsonStr string) map[string]interface{} {
	var mapResult map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &mapResult); err != nil {
		fmt.Println("json.Unmarshal failed:", err)
	}

	return mapResult
}

/**
 * struct to json
 */
func StructToJson(obj interface{}) string {
	jsonBytes, err := json.Marshal(obj)

	if err != nil {
		fmt.Println("json.Marshal failed:", err)
	}

	return string(jsonBytes)
}

/**
 * int to string
 */
func IntToString(v int) string {
	return strconv.Itoa(v)
}

/**
 * string to int
 */
func StringToInt(str string) int {
	_int, _ := strconv.Atoi(str)
	return _int
}

/**
 * float64 to int
 */
func FloatToInt(v float64) int {
	return int(v)
}

/**
 * struct to map
 */
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
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
	err := smtp.SendMail(host+":"+port, auth, user, sendTo, msg)

	return err
}

/**
 * 去除空格
 */
func TrimS(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}

/**
 * 声明 errors 类型
 */
func Error(str string) error {
	err := errors.New(str)
	return err
}

/**
 * 格式化日期
 */
func Date(format string) string {
	return time.Now().Format(format)
}

/**
 */
func CurrentPath(request *http.Request) string {
	return request.URL.Path
}
