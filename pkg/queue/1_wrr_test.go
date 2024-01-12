package queue_test

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func Example_Queue() {
	queue := new(PriorityQueue)
	heap.Init(queue)

	heap.Push(queue, &Item{value: "A", priority: 3})
	heap.Push(queue, &Item{value: "B", priority: 6})
	heap.Push(queue, &Item{value: "C", priority: 1})
	heap.Push(queue, &Item{value: "D", priority: 2})

	len := queue.Len()
	fmt.Println("优先队列长度：", len)

	item := (*queue)[0]
	fmt.Println("top 提取值: ", item.value)
	fmt.Println("top 提取优先级: ", item.priority)

	fmt.Println("遍历")
	for queue.Len() > 0 {
		fmt.Println(heap.Pop(queue).(*Item).value)
	}

	//Output:
	//优先队列长度： 4
}
