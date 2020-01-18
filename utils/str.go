package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"github.com/michaelzx/paladin-go/logger"
)

var REPEAT = true

var PREFIX = false

var Str = new(str)

type str struct {
}

func (s *str) UpperFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func (s *str) LowerFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func (s *str) Md5(psw string) string {
	newMd5 := md5.New()
	newMd5.Write([]byte(psw))
	return hex.EncodeToString(newMd5.Sum(nil))
}

// Substring 截取字符串
func (s *str) Substring(str string, startIdx int, subLen int) string {
	if str == "" {
		return ""
	}
	r := []rune(str)
	rLen := len(r)
	rMaxIdx := rLen - 1
	switch {
	case startIdx < 0:
		startIdx = 0
	case startIdx > rMaxIdx:
		return ""
	}
	endIdx := startIdx + subLen - 1
	if endIdx > rMaxIdx {
		endIdx = rMaxIdx
	}

	return string(r[startIdx : endIdx+1])
}

// StringSubFromStart 字符串截取，指定的开头到结尾
func (s *str) SubStringFromStart(str string, startIdx int) string {
	return s.Substring(str, startIdx, len(str)-startIdx)
}

// StringSubToEnd 字符串截取，从开头到指定的结尾
func (s *str) SubstringToEnd(str string, subLen int) string {
	return s.Substring(str, 0, subLen)
}
func (s *str) FromInt32(n int32) string {
	i := int64(n)
	return s.FromInt64(i)
}
func (s *str) ToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func (s *str) FromInt64(n int64) string {
	buf := [11]byte{}
	pos := len(buf)
	signed := n < 0
	if signed {
		n = -n
	}
	for {
		pos--
		buf[pos], n = '0'+byte(n%10), n/10
		if n == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

func (s *str) StarMask(str string) string {
	strLen := len(str)
	var starStartIdx int
	var starEndIdx int
	switch strLen {
	case 1:
		return str
	case 2:
		return string(str[0]) + "*"
	case 3:
		return string(str[0]) + "*" + string(str[2])
	case 11:
		starStartIdx = 3
		starEndIdx = 7
	default:
		starStartIdx = strLen / 3
		starEndIdx = strLen - starStartIdx
	}
	starLen := starEndIdx - starStartIdx
	var buffer bytes.Buffer
	buffer.WriteString(str[0:starStartIdx])
	for i := 0; i < starLen; i++ {
		buffer.WriteByte('*')
	}
	buffer.WriteString(str[starEndIdx:])
	return buffer.String()
}

// html 标签属性值替换（地址相关的：src、href、link）
/**
 * html 要操作的 html 内容，字符串
 * property 要操作的属性
 */
func (s *str) ReplaceLinkValue(html, property, value string, flag bool) string {
	propertyRegexp := regexp.MustCompile(property + `="(.*?)"`)
	match := propertyRegexp.FindAllString(html, -1)
	logger.Debug("match result: ", match)
	if match != nil {
		for _, v := range match {
			if !strings.HasPrefix(v, "http://") && !strings.HasPrefix(v, "https://") {
				contents := strings.Split(v, `"`)
				logger.Debug("contents: ", contents)
				repeat := value
				if !flag {
					repeat = value + contents[1]
				}
				html = strings.ReplaceAll(html, v, property+`="`+repeat+`"`)
			}
		}

	}
	return html
}
