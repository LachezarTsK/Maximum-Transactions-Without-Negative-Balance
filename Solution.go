
package main
import "container/heap"

func maxTransactions(transactions []int) int {
    var valueOfTransactions int64 = 0
    numberOfTransactions := 0
    minHeapNegativeValuesIncludedInTransactions := PriorityQueue{}

    for _, value := range transactions {

        if value >= 0 {
            numberOfTransactions++
            valueOfTransactions += int64(value)
            continue
        }

        if valueOfTransactions+int64(value) >= 0 {
            heap.Push(&minHeapNegativeValuesIncludedInTransactions, value)
            numberOfTransactions++
            valueOfTransactions += int64(value)
            continue
        }

        if minHeapNegativeValuesIncludedInTransactions.Len() > 0 && minHeapNegativeValuesIncludedInTransactions.Peek().(int) < value {
            valueOfTransactions += int64(value) - int64(heap.Pop(&minHeapNegativeValuesIncludedInTransactions).(int))
            heap.Push(&minHeapNegativeValuesIncludedInTransactions, value)
        }
    }

    return numberOfTransactions
}

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
    return len(pq)
}

func (pq PriorityQueue) Less(first int, second int) bool {
    return pq[first] < pq[second]
}

func (pq PriorityQueue) Swap(first int, second int) {
    pq[first], pq[second] = pq[second], pq[first]
}

func (pq *PriorityQueue) Push(object any) {
    *pq = append(*pq, object.(int))
}

func (pq *PriorityQueue) Pop() any {
    value := (*pq)[pq.Len() - 1]
    *pq = (*pq)[0 : pq.Len() - 1]
    return value
}

func (pq *PriorityQueue) Peek() any {
    return (*pq)[0]
}
