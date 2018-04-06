package main

import (
	"os"
	"log"
	"fmt"
	"bufio"
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
	topArray := []*TopicItem{
		&TopicItem{}, &TopicItem{},&TopicItem{},&TopicItem{},&TopicItem{},
		&TopicItem{}, &TopicItem{},&TopicItem{},&TopicItem{},&TopicItem{},
	}
	for topic, count := range uniques {
		item := &TopicItem{topic: topic, count: count}
		needInsert := false
		insertPos := 0

		if len(topArray) == 0 {
			topArray = append(topArray, item)
			continue
		}

		for k, item := range topArray {
			if count >= item.count {
				needInsert = true
				insertPos = k
				break
			}
		}
		if !needInsert {
			continue
		}

		tmpArray := []*TopicItem{}
		if insertPos == 0 {
			tmpArray = append(tmpArray, item)
			tmpArray = append(tmpArray, topArray[0:]...)
		} else {
			tmpArray = append(tmpArray, topArray[0: insertPos]...)
			tmpArray = append(tmpArray, item)
			tmpArray = append(tmpArray, topArray[insertPos:]...)
		}
		end := len(tmpArray)
		if end > 10 {
			end = 10
		}
		topArray = tmpArray[0:end]
	}

	for k, item := range topArray {
		fmt.Println("final top ", k, item.topic, item.count)
	}

	//heap
}
