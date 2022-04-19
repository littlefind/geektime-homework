package main

import (
	"fmt"
	"sync"
	"time"
)

type Bucket struct { //单个桶,窗口滚动时间间隔
	TotalRequestNum int64 //捅的总请求量
}

type Windows struct { //窗口
	Duration int64 //窗口时长,单位ms
	BucketDuration int64 //窗口滚动间隔,单位ms
	BucketNum int64 //桶的数量
	Buckets []*Bucket //窗口中所有桶的详细信息
	CurrentBucketIndex int64 //当前桶的位置
	CurrentBucketTime time.Time //当前桶的开始时间
	mutex sync.Mutex
	TotalRequestNum int64 //窗口的总请求量
	MaxTotalRequestNum int64 //窗口的请求阈值
}

func NewWindows(windowsDuration int64, BucketDuration int64, MaxTotalRequestNum int64) *Windows {
	BucketNum := windowsDuration / BucketDuration
	fmt.Println("桶的数量 = ", BucketNum)
	windows := &Windows{
		Duration:           windowsDuration,
		BucketDuration:     BucketDuration,
		BucketNum:          BucketNum,
		Buckets:            make([]*Bucket, BucketNum),
		CurrentBucketIndex: -1,//-1表示未设置为桶位置
		CurrentBucketTime: time.Now(),//-1表示未设置为桶位置
		TotalRequestNum:    0,
		MaxTotalRequestNum:             MaxTotalRequestNum,
	}
	for i := int64(0); i < BucketNum; i++ {
		windows.Buckets[i] = &Bucket{
			TotalRequestNum: 0,
		}
	}
	return windows
}

func (w *Windows) add()  {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if w.CurrentBucketIndex == -1 { //当首次进入时，桶的位置是1
		w.CurrentBucketIndex = 0
	}
	now := time.Now()
	//判断当前桶所在的位置(当前的时间 - 上个桶开始时间 > 窗口滚动间隔时间 则桶的位置需要发生变化)
	currentBucketIndex := w.CurrentBucketIndex
	if now.UnixMilli() - w.CurrentBucketTime.UnixMilli() > w.BucketDuration {
		if w.CurrentBucketIndex + 1 >= w.BucketNum {
			currentBucketIndex = 0
		} else {
			currentBucketIndex = w.CurrentBucketIndex + 1
		}
	}
	fmt.Println("当前的桶位置 = ", w.CurrentBucketIndex)

	//判断当前的桶的位置和上次窗口发生请求所在的桶的位置是否发生了变化，如果发生了变化，那么需要清空距离现在最远的1个桶的数据，也就是清空上一轮当前的桶
	if currentBucketIndex != w.CurrentBucketIndex { //说明桶的位置发生了变化
		fmt.Println("当前桶和上个请求的桶不是一个桶，进行相关设置")
		//清空距离当前桶最远的那个桶的数据，并且总访问量减去最远的那个桶的的访问量，也就是清空当前的桶
		w.TotalRequestNum -= w.Buckets[currentBucketIndex].TotalRequestNum
		w.Buckets[currentBucketIndex].TotalRequestNum = 0 // 清空当前桶的数据
		fmt.Println("清空第", currentBucketIndex, "个桶的TotalRequestNum为0")
		w.CurrentBucketIndex = currentBucketIndex //更新当前桶的位置
		fmt.Println("更新桶的位置为 = ", currentBucketIndex)
		w.CurrentBucketTime = now //更新当前桶的开始时间
		fmt.Println("更新桶的开始时间为 = ", now)
	}

	//判断当前窗口，访问总数是否已满，如果已满则返回失败，否则增加当前的访问记录
	if (w.TotalRequestNum + 1) > w.MaxTotalRequestNum {
		fmt.Println("当前访问总数 = ", w.TotalRequestNum, "，再次请求后，将超过最大访问限制总数，拒绝访问")
		return
	}

	//窗口总访问数量+1
	w.TotalRequestNum++
	//当前桶的总访问量+1
	w.Buckets[currentBucketIndex].TotalRequestNum++
	fmt.Println("当前窗口总访问量 = ", w.TotalRequestNum, ", 当前桶的位置 = ", currentBucketIndex, ", 当前桶的总访问量= ", w.Buckets[currentBucketIndex].TotalRequestNum)
}

func main() {
	//设置三个参数
	//1.窗口时长windowsDuration(单位毫秒)
	//2.窗口滚动间隔(也就是单个桶的时长,单位毫秒)bucketDuration
	//3.窗口最大请求量MaxTotalRequestNum
	windows := NewWindows(1000, 100, 50)

	//循环模拟请求
	//1.第一种,固定每10ms请求一次，总共请求10秒
	now := time.Now()
	for i := 1; i <= 100; i++ { //最多模拟1000次请求
		fmt.Println("第", i, "次请求")
		if now.Sub(time.Now()).Seconds() > float64(10) {
			fmt.Println("第11秒的请求跳出循环")
			break
		}
		windows.add()
		time.Sleep(time.Millisecond * 10) //每次请求后，休息10ms
	}

	//2.第二种随机1~200ms请求一次，总共请求10秒
}
