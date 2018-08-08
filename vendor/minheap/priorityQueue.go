//Package minheap implements the min heap structure
package minheap

//PriorityQueue implements a min heap
type PriorityQueue []*Item

//Len returns number of elemsnts in the heap
func (pq PriorityQueue) Len() int {
	return len(pq)
}

//Less verifies priority order between two items in the heap
func (pq PriorityQueue) Less(i int, j int) bool {
	//The lower the value, the higher the priority
	return pq[i].Value < pq[j].Value
}

//Swap swaps two items in the heap
func (pq PriorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

//Push inserts element into priority heap
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

//Pop returns minimum element of the heap
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
