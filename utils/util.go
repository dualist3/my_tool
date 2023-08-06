package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func PackInt64ToString(Packs []int64) string {
	var TempString = ""
	for _, v := range Packs {
		stringInt64 := strconv.FormatInt(v, 10)
		if TempString == "" {
			TempString = TempString + stringInt64
		} else {
			TempString = TempString + "," + stringInt64
		}
	}
	return TempString
}

func PackStringToString(Packs []string) string {
	var TempString = ""
	for _, v := range Packs {
		stringInt64 := v
		if TempString == "" {
			TempString = TempString + stringInt64
		} else {
			TempString = TempString + "," + stringInt64
		}
	}
	return TempString
}

func UnpackStringToInt64Slice(s string) (int64Slice []int64) {
	Ids := strings.Split(s, ",")

	for _, v := range Ids {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
		int64Slice = append(int64Slice, id)
	}
	return int64Slice
}

func UnpackStringToStringSlice(s string) (int64Slice []string) {
	if s == "" {
		return nil
	}
	Ids := strings.Split(s, ",")

	return Ids
}

func GenerateID() string {
	var prefix string = "N"
	rand.Seed(time.Now().UnixNano())
	var number int = rand.Intn(1000000)
	return fmt.Sprintf("%s%d", prefix, number)
}

// GetIdCardAge 根据身份证号获取年龄
func GetIdCardAge(idCard string) (int, error) {
	layout := "20060102"
	birthDate, err := time.Parse(layout, idCard[6:14])
	if err != nil {
		return 0, err
	}

	now := time.Now()
	age := now.Year() - birthDate.Year()

	if now.Month() < birthDate.Month() || (now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		age--
	}

	return age, nil
}

// IsIDCardValid 判断身份证号是否合法
func IsIDCardValid(idCard string) bool {
	// 匹配身份证号规则（18位）
	pattern := `^[1-9]\d{5}(19|20)\d{2}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])\d{3}[\dXx]$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(idCard)
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func IsTwelveYearsOldAt(id string, timestamp int64) bool {
	// 将时间戳转换为 time.Time 对象
	t := time.Unix(timestamp, 0)

	// 解析身份证号中的出生日期
	birthDate, err := time.Parse("20060102", id[6:14])
	if err != nil {
		panic(err)
	}

	// 计算出生日期和目标时间之间的时间差
	age := int(t.Sub(birthDate).Hours() / 24 / 365)

	// 返回年龄是否大于等于 12
	return age >= 12
}
func GetMidnightTime() int64 {
	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return midnight.Unix()
}
func GetMonthStartAndEnd(timestamp int64) (int64, int64) {
	t := time.Unix(timestamp, 0).UTC()
	year, month, _ := t.Date()

	// 获取当月开始时间
	monthStart := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Unix()

	// 获取下个月开始时间，然后减去1秒，即为当月结束时间
	nextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	monthEnd := nextMonth.Add(-time.Second).Unix()

	return monthStart, monthEnd
}
func MaxConsecutiveDays(timestamps []int64) int {
	if len(timestamps) == 0 {
		return 0
	}

	maxDays := 1
	currentDays := 1

	for i := 1; i < len(timestamps); i++ {
		if timestamps[i] == timestamps[i-1]+86400 { // 检查下一个时间戳是否比前一个时间戳多一天（86400秒）
			currentDays++
			if currentDays > maxDays {
				maxDays = currentDays
			}
		} else {
			currentDays = 1 // 重置连续天数为1
		}
	}

	return maxDays
}
