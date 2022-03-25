package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
)

//HttpPostForm 获取post表单数据
func HttpPostForm(gourl string, form url.Values) ([]byte, error) {
	resp, err := http.PostForm(gourl, form)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//RandInt 得到一个随机数
func RandInt(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if max < 1 {
		return r.Int()
	} else {
		return r.Intn(max)
	}
}

//GetLocalIP 获取本地ip
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

//GetDate 获取当前时间
func GetDate() string {
	return time.Now().Format("2006-01-02")
}

// RemoteIp 获取远程客户端的 IP
func RemoteIp(req *http.Request) string {
	addr := httpx.GetRemoteAddr(req)
	//可能会被仿冒，我们取最后一个IP，也就是原始IP
	strs := strings.Split(addr, ",")
	return strings.TrimSpace(strs[len(strs)-1])
}

// RemotePlatform 获取远程客户端系统：ipad,iphone,mac,windows,android,linux,ubuntu
func RemotePlatform(r *http.Request) (string, error) {
	u := r.Header.Get("User-Agent")
	if len(u) == 0 {
		u = r.Header.Get("user-agent")
		if len(u) == 0 {
			return u, errors.New("未知设备的请求！")
		}
	}
	//ipad
	padR := regexp.MustCompile(`(?i)ipad`)
	if padR.MatchString(u) {
		return "ipad", nil
	}
	//iphone
	iphoneR := regexp.MustCompile(`(?i)iphone`)
	if iphoneR.MatchString(u) {
		return "iphone", nil
	}
	//mac
	macR := regexp.MustCompile(`(?i)mac`)
	if macR.MatchString(u) {
		return "mac", nil
	}
	//windows
	windowsR := regexp.MustCompile(`(?i)windows`)
	if windowsR.MatchString(u) {
		return "windows", nil
	}
	//android
	androidR := regexp.MustCompile(`(?i)android`)
	if androidR.MatchString(u) {
		return "android", nil
	}
	//linux
	linuxR := regexp.MustCompile(`(?i)linux`)
	if linuxR.MatchString(u) {
		return "linux", nil
	}
	//ubuntu
	ubuntuR := regexp.MustCompile(`(?i)ubuntu`)
	if ubuntuR.MatchString(u) {
		return "ubuntu", nil
	}
	return u, errors.New("未知设备的请求")
}

func GetJwtUid(ctx context.Context) (int64, error) {
	jwtUserId := ctx.Value("userId")
	uid, err := json.Number(fmt.Sprintf("%v", jwtUserId)).Int64()
	if err != nil {
		return 0, errors.New("获取不到你的ID")
	}
	return uid, nil
}

func GetJwtRoleId(ctx context.Context) (int64, error) {
	JwtRoleId := ctx.Value("userRoleId")
	uid, err := json.Number(fmt.Sprintf("%v", JwtRoleId)).Int64()
	if err != nil {
		return 0, errors.New("获取不到你角色ID")
	}
	return uid, nil
}

func GetJwtCompanyId(ctx context.Context) (int64, error) {
	JwtUserCompanyId := ctx.Value("userCompanyId")
	uid, err := json.Number(fmt.Sprintf("%v", JwtUserCompanyId)).Int64()
	if err != nil {
		return 0, errors.New("获取不到你所属公司ID")
	}
	return uid, nil
}

//IsSlice 判断数据是否为切片
func IsSlice(data interface{}) bool {
	if data == nil {
		return false
	}
	name := reflect.TypeOf(data).Kind().String()
	return name == "slice"
}

func SliceLen(data interface{}) int {
	if !IsSlice(data) {
		return -1
	}
	return reflect.ValueOf(data).Len()
}

//expToNumber 科学计算数字转int64
//例： ExpNum = 1.641020629e+09
func ExpToNumber(ExpNum string) (int64, error) {
	if res, err := strconv.ParseInt(ExpNum, 10, 64); err == nil {
		return res, nil
	}
	var newNum float64
	_, err := fmt.Sscanf(ExpNum, "%e", &newNum)
	if err != nil {
		return 0, fmt.Errorf("fmt.Sscanf error, ExpiresAt:%s, err:%v", ExpNum, err)
	}
	num := fmt.Sprintf("%.f", newNum)
	res, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("strconv.ParseInt error, num:%s, err:%v", num, err)
	}
	return res, nil
}

//GetReqFormValue 获取请求form字段值
func GetReqFormValue(r *http.Request, Key string) string {
	//r.Form 类型为 map[string][]string
	if r.Form.Has(Key) {
		return strings.TrimSpace(r.Form.Get(Key))
	}
	return ""
}
