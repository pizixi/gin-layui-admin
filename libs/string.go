/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-08 00:24:25
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-17 10:12:06
***********************************************/

package libs

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

var emailPattern = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func SizeFormat(size float64) string {
	units := []string{"Byte", "KB", "MB", "GB", "TB"}
	n := 0
	for size > 1024 {
		size /= 1024
		n += 1
	}

	return fmt.Sprintf("%.2f %s", size, units[n])
}

func IsEmail(b []byte) bool {
	return emailPattern.Match(b)
}

func Password(len int, pwdO string) (pwd string, salt string) {
	salt = GetRandomString(4)
	defaultPwd := "admin"
	if pwdO != "" {
		defaultPwd = pwdO
	}
	pwd = Md5([]byte(defaultPwd + salt))
	return pwd, salt
}

// 生成32位MD5
// func MD5(text string) string{
//    ctx := md5.New()
//    ctx.Write([]byte(text))
//    return hex.EncodeToString(ctx.Sum(nil))
// }

//生成随机字符串
func GetRandomString(lens int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func FormatTime(v int64) string {
	return time.Unix(v, 0).Format("2006-01-02 15:04:05")
}

func ParseTime(tmstr string) int64 {
	// 注意：time.Parse 在linux下，使用的UTC时区
	// 参考：https://www.jianshu.com/p/92d9344425a7
	tm, err := time.ParseInLocation("2006-01-02 15:04:05", tmstr, time.Local)
	if err == nil {
		return tm.Unix()
	}
	return 0
}
