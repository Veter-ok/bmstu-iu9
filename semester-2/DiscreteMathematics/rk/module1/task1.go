package main

func MergeSort(items int, compare func(i, j int) int, indices chan int) {
    defer close(indices)
    mergeSortRec(0, items - 1, compare, indices)
}

func mergeSortRec(low, high int, compare func(i, j int) int, indices chan int) {
    if low == high {
        indices <- low
        return
    }
    mid := (low + high) / 2;
    source1 := make(chan int)
    source2 := make(chan int)
    go func() {
        mergeSortRec(low, mid, compare, source1)
        close(source1)
    }()
    go func() {
        mergeSortRec(mid+1, high, compare, source2)
        close(source2)
    }()
    merge(source1, source2, indices, compare);
}

func merge(source1, source2, dist chan int, compare func(i, j int) int){
    i, ok1 := <-source1
    j, ok2 := <-source2
    for ok1 && ok2 {
        if compare(i, j) <= 0 {
            dist <- i
            i, ok1 = <-source1
        } else {
            dist <- j
            j, ok2 = <-source2
        }
    }
    for ok1 {
        dist <- i
        i, ok1 = <-source1
    }

    for ok2 {
        dist <- j
        j, ok2 = <-source2
    }
}