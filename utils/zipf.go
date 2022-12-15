package utils

import (
	"math"
	"math/rand"
)

// 公式 f(x) = (1/x^theta) / sum[n=1:n=N](1/n^theta)
// 公式中的theta, theta小则曲线平滑、分布更均匀、冲突少
// N为zipf分布的元素数量
func MakeZipfArray(theta float64, N int) []float64 {
	var prob float64
	var sum float64
	zipf := []float64{}
	for x:=1; x<=N; x++ {
		ntheta := 1/math.Pow(float64(x), theta)
		sum += ntheta
	}
	for x:=1; x<=N; x++ {
		xtheta := 1/math.Pow(float64(x), theta)
		prob += xtheta/sum
		zipf = append(zipf, prob)
	}
	return zipf
}

func GetZipfRand(count int, zipf []float64, seed int64) []int {
	//rand.Seed(seed)
	idx := 1
	zipfRand := []int{}
	for x:=1; x<=count; x++ {
		randProb := rand.Float64()
		for idx=0; idx<len(zipf); idx++ {
			if randProb <= zipf[idx] {
				zipfRand = append(zipfRand, idx+1)
				break
			} else {
				continue
			}
		}
		if idx == len(zipf) {
			zipfRand  = append(zipfRand, len(zipf))
		}
	}
	return zipfRand
}


