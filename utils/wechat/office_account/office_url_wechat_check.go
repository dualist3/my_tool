package office_account

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

func WechatOfficeUrlCheck(token, signature, timestamp, nonce string) bool {
	// 将 timestamp、nonce、token 进行字典序排序
	tmpArr := []string{token, timestamp, nonce}
	sort.Strings(tmpArr)

	// 将排序后的参数拼接成一个字符串
	tmpStr := strings.Join(tmpArr, "")

	// 对拼接后的字符串进行 sha1 加密
	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(tmpStr))
	hashed := fmt.Sprintf("%x", sha1Hash.Sum(nil))

	// 比较加密后的字符串与 signature 是否一致
	if hashed == signature {
		return true
	} else {
		return false
	}
}
