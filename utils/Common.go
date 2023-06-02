package utils

import (
	"bytes"
	"compress/gzip"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/shopspring/decimal"
)

// GenerateMd5
// 生成Md5
func GenerateMd5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return fmt.Sprintf("%x", m.Sum(nil))
}

// HexEncode
// hex加密
func HexEncode(str string) string {
	return hex.EncodeToString([]byte(str))
}

// HexDecode
// hex解密
func HexDecode(str string) ([]byte, error) {
	return hex.DecodeString(str)
}

// GenerateSha256
// 生成Sha256
func GenerateSha256(str string) string {
	return fmt.Sprintf("%X", sha256.New().Sum([]byte(str)))
}

// GenerateHmac
// 生成Hmac
func GenerateHmac(message, keys string) string {
	h := hmac.New(sha256.New, []byte(keys))
	h.Write([]byte(message))
	return fmt.Sprintf("%X", h.Sum(nil))
}

// GenerateUrlValuesByMapString
// 根据MapString生成UrlValues
func GenerateUrlValuesByMapString(param map[string]string) url.Values {
	values := make(url.Values, 0)
	for k, v := range param {
		values.Set(k, v)
	}
	return values
}

// RefStructGetFieldPtr
// 反射结构体获取字段指针
func RefStructGetFieldPtr(stc interface{}, field ...interface{}) ([]interface{}, error) {
	stcTypeOf := reflect.TypeOf(stc)
	stcValueOf := reflect.ValueOf(stc)
	fieldPtrSlice := make([]interface{}, 0)
	if stcTypeOf.Elem().Kind() != reflect.Struct {
		log.Println("stcTypeOf.Kind() :", stcTypeOf.Kind())
		return nil, errors.New("parameter type is not Struct")
	}

	if stcTypeOf.Kind() != reflect.Ptr {
		log.Println("stcTypeOf.Kind() :", stcTypeOf.Kind())
		return nil, errors.New("parameter type is not Ptr")
	}

	stcFieldNum := stcValueOf.Elem().NumField()

	for i := 0; i < stcFieldNum; i++ {
		fieldName := stcTypeOf.Elem().Field(i).Name

		for _, v := range field {
			if v == "*" || fieldName == v {
				fieldPointer := reflect.NewAt(stcValueOf.Elem().Field(i).Type(), unsafe.Pointer(stcValueOf.Elem().Field(i).UnsafeAddr())).Pointer()
				fieldKind := stcValueOf.Elem().Field(i).Type().Kind()
				switch fieldKind {
				case reflect.Int:
					fieldContext := (*int)(unsafe.Pointer(fieldPointer))
					fieldPtrSlice = append(fieldPtrSlice, fieldContext)
					break
				case reflect.String:
					fieldContext := (*string)(unsafe.Pointer(fieldPointer))
					fieldPtrSlice = append(fieldPtrSlice, fieldContext)
					break
				case reflect.Float64:
					fieldContext := (*float64)(unsafe.Pointer(fieldPointer))
					fieldPtrSlice = append(fieldPtrSlice, fieldContext)
					break
				}
			}
		}

	}
	return fieldPtrSlice, nil
}

// RefStructChangeMap
// 反射结构体转换成map
func RefStructChangeMap(structPtr interface{}) (map[string]string, error) {

	typeOf := reflect.TypeOf(structPtr)
	valueOf := reflect.ValueOf(structPtr)
	kind := typeOf.Kind()

	if kind != reflect.Ptr {
		return nil, errors.New(" Parameter &Type Not *Ptr ")
	}
	if typeOf.Elem().Kind() != reflect.Struct {
		return nil, errors.New(" Parameter *Type Not Struct ")
	}

	container := make(map[string]string)
	fieldNum := typeOf.Elem().NumField()
	for i := 0; i < fieldNum; i++ {
		fieldName := typeOf.Elem().Field(i).Name
		fieldTag := typeOf.Elem().Field(i).Tag.Get("json")

		if valueOf.Elem().FieldByName(fieldName).String() == "" {
			continue
		}

		if strings.Index(fieldTag, ",") == -1 {
			if typeOf.Elem().Field(i).Type.String() == "[]map[string]string" {

				mapInterface := valueOf.Elem().FieldByName(fieldName).Interface()
				maps := mapInterface.([]map[string]string)
				for k, v := range maps {
					for s, s2 := range v {
						keys := fmt.Sprintf("%s[%d][%s]", fieldTag, k, s)
						container[keys] = s2
					}

				}
				continue
			}
			container[fieldTag] = valueOf.Elem().FieldByName(fieldName).String()
		} else {
			keys := strings.Split(fieldTag, ",")
			container[keys[0]] = valueOf.Elem().FieldByName(fieldName).String()
		}
	}
	//log.Println("kind : ", kind)
	//log.Println("*kind : ", typeOf.Elem().Kind())
	//log.Println("data : ", container)
	return container, nil
}

// TrimRightZeroByFloatStr
// 根据浮点型字符串去除零
func TrimRightZeroByFloatStr(floatStr string) string {

	if strings.Index(floatStr, ".") != -1 {
		floatStr = strings.TrimRight(floatStr, "0")
		floatStr = strings.TrimRight(floatStr, ".")
	}

	return floatStr
}

// InCheckIntOrIntSlice
// 检查int类型是否存在int类型切片中
func InCheckIntOrIntSlice(index int, checkValue []int) bool {
	for _, v := range checkValue {
		if index == v {
			return true
		}
	}

	return false
}

// InCheckStringOrStringSlice
// 检查string类型是否存在string类型切片中
func InCheckStringOrStringSlice(index string, checkValue []string) bool {
	for _, v := range checkValue {
		if index == v {
			return true
		}
	}

	return false
}

// DeleteStringSliceByString
// 根据string删除string类型切片中的元素
func DeleteStringSliceByString(a []string, elem string) []string {
	j := 0
	for _, v := range a {
		if v != elem {
			a[j] = v
			j++
		}
	}
	return a[:j]
}

// MatchRandBoolSlicePop
// 随机弹出Bool类型切片中的值
func MatchRandBoolSlicePop(boolSlice []bool) bool {

	var res bool
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(boolSlice); i++ {
		index := rand.Intn(len(boolSlice))
		res = boolSlice[index]
	}

	return res
}

// MatchRandFloatBySectionAndPec
// 根据定制区间随机生成浮点字符串
func MatchRandFloatBySectionAndPec(min, max string, pec int) (string, error) {

	minParseFloat, err := strconv.ParseFloat(min, 64)
	if err != nil {
		return "", err
	}
	maxParseFloat, err := strconv.ParseFloat(max, 64)
	if err != nil {
		return "", err
	}

	rand.Seed(time.Now().UnixNano())
	num := rand.Float64()*(maxParseFloat-minParseFloat) + minParseFloat

	return strconv.FormatFloat(num, 'f', pec, 64), nil
}

// CheckDecimalLenByString
// 根据字符串检查是否有小数点位数(必须是浮点字符串)
func CheckDecimalLenByString(decimalStr string) int {

	decimals := 0
	indexOf := strings.Index(decimalStr, ".")

	if indexOf != -1 {
		decimalStr = strings.TrimRight(decimalStr, "0")
		decimalStr = strings.TrimRight(decimalStr, ".")
		decimalOf := strings.Index(decimalStr, ".")

		if decimalOf != -1 {
			decimals = len(decimalStr) - decimalOf - 1
		}
	}

	return decimals
}

// BcSub
// 两个任意精度数字的减法
func BcSub(num1, num2 string) (float64, error) {

	num1ParseFloat, err := strconv.ParseFloat(num1, 64)
	if err != nil {
		return 0, err
	}
	num2ParseFloat, err := strconv.ParseFloat(num2, 64)
	if err != nil {
		return 0, err
	}

	num1Decimal := decimal.NewFromFloat(num1ParseFloat)
	num2Decimal := decimal.NewFromFloat(num2ParseFloat)

	res, _ := num1Decimal.Sub(num2Decimal).Float64()

	return res, nil
}

// BcAdd
// 两个任意精度数字的加法计算
func BcAdd(num1, num2 string) (float64, error) {

	num1ParseFloat, err := strconv.ParseFloat(num1, 64)
	if err != nil {
		return 0, err
	}
	num2ParseFloat, err := strconv.ParseFloat(num2, 64)
	if err != nil {
		return 0, err
	}

	num1Decimal := decimal.NewFromFloat(num1ParseFloat)
	num2Decimal := decimal.NewFromFloat(num2ParseFloat)

	res, _ := num1Decimal.Add(num2Decimal).Float64()

	return res, nil
}

// BcMul
// 两个任意精度数字乘法计算
func BcMul(num1, num2 string) (float64, error) {
	num1ParseFloat, err := strconv.ParseFloat(num1, 64)
	if err != nil {
		return 0, err
	}
	num2ParseFloat, err := strconv.ParseFloat(num2, 64)
	if err != nil {
		return 0, err
	}

	num1Decimal := decimal.NewFromFloat(num1ParseFloat)
	num2Decimal := decimal.NewFromFloat(num2ParseFloat)

	res, _ := num1Decimal.Mul(num2Decimal).Float64()

	return res, nil
}

// BcDiv
// 两个任意精度的数字除法计算
func BcDiv(num1, num2 string) (float64, error) {
	num1ParseFloat, err := strconv.ParseFloat(num1, 64)
	if err != nil {
		return 0, err
	}
	num2ParseFloat, err := strconv.ParseFloat(num2, 64)
	if err != nil {
		return 0, err
	}

	num1Decimal := decimal.NewFromFloat(num1ParseFloat)
	num2Decimal := decimal.NewFromFloat(num2ParseFloat)

	res, _ := num1Decimal.Div(num2Decimal).Float64()

	return res, nil
}

// BcComp
// 两个任意精度的数字比较计算
func BcComp(num1, num2 string) (float64, error) {
	num1ParseFloat, err := strconv.ParseFloat(num1, 64)
	if err != nil {
		return 0, err
	}
	num2ParseFloat, err := strconv.ParseFloat(num2, 64)
	if err != nil {
		return 0, err
	}

	num1Decimal := decimal.NewFromFloat(num1ParseFloat)
	num2Decimal := decimal.NewFromFloat(num2ParseFloat)
	res := float64(num1Decimal.Cmp(num2Decimal))
	return res, nil
}

// GZipDecompress
// 解析二进制byte字节
func GZipDecompress(input []byte) (string, error) {
	buf := bytes.NewBuffer(input)
	reader, gzipErr := gzip.NewReader(buf)
	if gzipErr != nil {
		return "", gzipErr
	}
	defer reader.Close()

	result, readErr := ioutil.ReadAll(reader)
	if readErr != nil {
		return "", readErr
	}

	return string(result), nil
}

// GZipCompress
// 转换string变成二进制byte字节
func GZipCompress(input string) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)

	_, err := gz.Write([]byte(input))
	if err != nil {
		return nil, err
	}

	err = gz.Flush()
	if err != nil {
		return nil, err
	}

	err = gz.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
