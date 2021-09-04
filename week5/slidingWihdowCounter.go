package week5

import (
	"sync"
	"time"
)

var (
	timeSpan int64 = 2  //采样间隔
	winSize  int64 = 10 //滑动窗口大小
)

type Number struct {
	Buckets  map[int64]*numberBucket
	lastTime int64   //上次计算时间
	Result   float64 //窗口内的最大值
	Mutex    *sync.Mutex
}

type numberBucket struct {
	Value float64
}

func NewNumber() *Number {
	return &Number{
		Buckets: make(map[int64]*numberBucket),
		Mutex:   new(sync.Mutex),
	}
}

//Add 数据采样
func (r *Number) Add(i float64) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	now := time.Now().Unix()
	r.Buckets[now] = &numberBucket{i}
	return
}

//Max 获取滑动窗口内最大值
func (r *Number) Max() float64 {
	if (time.Now().Unix() - r.lastTime) < timeSpan {
		return r.Result
	}

	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	r.removeOldBuckets()
	r.Result = 0
	for _, value := range r.Buckets {
		r.Result = max(r.Result, value.Value)
	}

	r.lastTime = time.Now().Unix()
	return r.Result
}

//更新滑动窗口内数据
func (r *Number) removeOldBuckets() {
	leftLimit := time.Now().Unix() - winSize
	for timeStamp, _ := range r.Buckets {
		if timeStamp < leftLimit {
			delete(r.Buckets, timeStamp)
		}
	}

	return
}

func max(i, j float64) float64 {
	if i > j {
		return i
	}
	return j
}
