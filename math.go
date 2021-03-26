package tools

import (
	"github.com/shopspring/decimal"
	"github.com/zheng-ji/goSnowFlake"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func FloatSub(x, y float64) float64 {
	res, _ := decimal.NewFromFloat(x).Sub(decimal.NewFromFloat(y)).Float64()
	return res
}

func FloatAdd(x, y float64) float64 {
	res, _ := decimal.NewFromFloat(x).Add(decimal.NewFromFloat(y)).Float64()
	return res
}

func FloatMul(x, y float64) float64 {
	res, _ := decimal.NewFromFloat(x).Mul(decimal.NewFromFloat(y)).Float64()
	return res
}

func FloatDiv(x, y float64) float64 {
	res, _ := decimal.NewFromFloat(x).Div(decimal.NewFromFloat(y)).Float64()
	return res
}

const (
	mWord = "abcdefghijklmnopqrstuvwxyz_"
	num   = "1234567890"
)

// 生成指定长度的随机字符串
func RandStr(cnt int) string {
	s := make([]byte, 0)
	// 组合包含大小写字母和数字下划线的集合
	for i := 0; i < cnt; i++ {
		s = append(s, randStrSet[RandIntN(len(randStrSet))])
	}
	return string(s)
}

func RandIntN(n int) int {
	// 设置一次随机数种子,每次整个程序启动时更新
	rand.Seed(time.Now().Unix())
	return rand.Intn(n)
}

var (
	worker     *goSnowFlake.IdWorker
	snowOnce   sync.Once
	randStrSet = mWord + strings.ToUpper(mWord) + num
)

func SnowFlakeId() (int64, error) {
	snowOnce.Do(func() {
		var err error
		wd, _ := strconv.ParseInt(os.Getenv("WorkerId"), 10, 64)
		worker, err = goSnowFlake.NewIdWorker(wd + 1)
		if err != nil {
			log.Fatalf("failed get snowFlake worker")
		}
	})
	return worker.NextId()
}
