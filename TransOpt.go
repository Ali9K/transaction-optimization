package main

import (
	"container/heap"
	"fmt"
)

type item struct {
	bank   int
	amount int
}

type PriorityQueue []*item

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].amount > pq[j].amount
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*item))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {

	var n, m int
	fmt.Scan(&n, &m)
	balance := make([]int, n+1)

	var u, v, w int
	for i := 1; i <= m; i++ {
		fmt.Scan(&u, &v, &w)
		balance[u] -= w
		balance[v] += w
	}

	debtors := &PriorityQueue{}
	creditors := &PriorityQueue{}

	heap.Init(debtors)
	heap.Init(creditors)

	for i := 1; i <= n; i++ {
		if balance[i] < 0 {
			heap.Push(debtors, &item{bank: i, amount: -balance[i]})
		} else if balance[i] > 0 {
			heap.Push(creditors, &item{bank: i, amount: balance[i]})
		}
	}

}
