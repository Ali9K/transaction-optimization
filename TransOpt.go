package main

type item struct {
	bank   int
	amount int
}

type PriorityQueue []*item

func (pq PriorityQueue) len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].amount > pq[j].amount
}
