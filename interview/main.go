package main

import (
	"os"
	"log"
	"fmt"
	"bufio"
	"sort"
)

type TopicItem struct {
	topic string
	count int
}

const (
	topNum = 10
)

func main()  {
	f, err := os.Open("./access.log")
	if err != nil {
		log.Println(fmt.Sprintf("open file failed, err is %v", err))
		return
	}

	defer f.Close()

	totalLogs := 0
	uniques := make(map[string]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		totalLogs += 1
		text := scanner.Text()
		if _, ok := uniques[text]; !ok {
			uniques[text] = 0
		}
		uniques[text] += 1
	}
	fmt.Println(fmt.Sprintf("total logs=%v, uniques=%v", totalLogs, len(uniques)))

	// 部分排序算法，数据移动量比较大，考虑用根堆
	topArray := []TopicItem{}

	for topic, count := range uniques {
		item := TopicItem{topic: topic, count: count}
		topArray = append(topArray, item)
		sort.Slice(topArray, func(i, j int) bool {
			return topArray[i].count > topArray[j].count
		})
		if len(topArray) > 10 {
			topArray = topArray[0:10]
		}
	}

	for k, item := range topArray {
		fmt.Println("final top ", k, item.topic, item.count)
	}
}
