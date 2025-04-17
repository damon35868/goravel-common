package utils

import (
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/goravel/framework/contracts/validation"
	"github.com/samber/lo"
)

func MarshalStr(data any) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}

func MD5(s string) string {
	hash := md5.New()
	_, err := hash.Write([]byte(s))
	if err != nil {
		panic(err)
	}
	sum := hash.Sum(nil)
	return fmt.Sprintf("%x\n", sum)
}

func GetTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetFormatTime(time time.Time) string {
	return time.Format("20060102")
}

func GenerateOrderNo() string {
	randomNum := rand.Intn(1000)
	date := GetFormatTime(time.Now())
	return fmt.Sprintf("%s%d%03d", date, GetTimestamp(), randomNum)
}

func FormatNullString(data string, isLikes ...bool) sql.NullString {
	val, isLike := data, true
	if len(isLikes) > 0 {
		isLike = isLikes[0]
	}
	if isLike && data != "" {
		val = "%" + val + "%"
	}
	content := sql.NullString{
		String: val,
		Valid:  data != "",
	}

	return content
}

func TakeRandNum(len int64) string {
	var list []string
	nums := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; int64(i) < len; i++ {
		nums := lo.Shuffle(nums)
		list = append(list, strconv.FormatInt(nums[0], 10))
	}
	return strings.Join(list, "")
}

func FormatRequest(key string, data validation.Data) error {
	if val, exist := data.Get(key); exist {

		switch g := val.(type) {
		case int:
			return data.Set(key, int64(g))
		case float64:
			return data.Set(key, int64(g))
		case []any:
			for _, v := range g {
				switch val := v.(type) {
				case int:
					return data.Set(key, int64(val))
				case float64:
					return data.Set(key, int64(val))
				}
			}
		}
	}
	return nil
}
