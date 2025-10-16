
class Solution {

    fun maxTransactions(transactions: IntArray): Int {
        var valueOfTransactions: Long = 0
        var numberOfTransactions = 0
        val minHeapNegativeValuesIncludedInTransactions = PriorityQueue<Int>()

        for (value in transactions) {

            if (value >= 0) {
                ++numberOfTransactions
                valueOfTransactions += value
                continue
            }

            if (valueOfTransactions + value >= 0) {
                minHeapNegativeValuesIncludedInTransactions.add(value)
                ++numberOfTransactions
                valueOfTransactions += value
                continue
            }

            if (!minHeapNegativeValuesIncludedInTransactions.isEmpty() && minHeapNegativeValuesIncludedInTransactions.peek() < value) {
                valueOfTransactions += value - minHeapNegativeValuesIncludedInTransactions.poll()
                minHeapNegativeValuesIncludedInTransactions.add(value)
            }
        }

        return numberOfTransactions
    }
}
