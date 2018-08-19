package main

import (
    "fmt"
)

func HeapSort(array []int) {
    m := len(array)
    s := m/2
    for i := s; i > -1; i-- {
        heap(array, i, m-1)
    }
    fmt.Println("....after first heap", array)
    for i := m-1; i > 0; i-- {
        array[i], array[0] = array[0], array[i]
        heap(array, 0, i-1)
        fmt.Println("...sort heap", 0, i - 1, array)
    }
}

func heap(array []int, i, end int){
    l := 2*i+1
    if l > end {
        return
    }
    n := l
    r := 2*i+2
    if r <= end && array[r]>array[l]{
        n = r
    }
    if array[i] > array[n]{
        return
    }
    array[n], array[i] = array[i], array[n]
    heap(array, n, end)
}

func main()  {
    array := []int{
        55, 94,87,1,4,32,11,77,39,42,64,53,70,12,9,
    }
    fmt.Println(array)
    HeapSort(array)
    fmt.Println("final array", array)
}