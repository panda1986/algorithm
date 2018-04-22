package main

import "fmt"

func mergeSort(a []int) []int {
    if len(a) == 1{
        return a
    }
    half := len(a) / 2
    l0 := a[0: half]
    l1 := a[half:]
    fmt.Println(fmt.Sprintf("l0:%v, l1:%v", l0, l1))

    x := mergeSort(l0)
    y := mergeSort(l1)
    fmt.Println(x, y)
    return merge(x, y)
}

func merge(a0, a1 []int) (a []int) {
    fmt.Println(fmt.Sprintf("merge a0:%v, a1:%v", a0, a1))
    for _, v0 := range a0 {
        for _, v1 := range a1 {
            if v0 < v1 {
                a = append(a, v0)
            } else {
                a = append(a, v1)
            }
        }
    }
    fmt.Println("after merge:%v", a)
    return
}

func main()  {
    arr := []int{14, 33, 27, 10, 35, 19, 42, 44}
    res := mergeSort(arr)
    fmt.Println("res:%v", res)
}
