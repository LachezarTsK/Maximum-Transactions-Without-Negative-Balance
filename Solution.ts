
// const {PriorityQueue} = require('@datastructures-js/priority-queue');
/*
 PriorityQueue is internally included in the solution file on leetcode.
 When running the code on leetcode it should stay commented out. 
 It is mentioned here just for information about the external library 
 that is applied for this data structure.
 */

function maxTransactions(transactions: number[]): number {
    let valueOfTransactions = 0;
    let numberOfTransactions = 0;
    const minHeapNegativeValuesIncludedInTransactions = new PriorityQueue<number>((x, y) => x - y);

    for (let value of transactions) {

        if (value >= 0) {
            ++numberOfTransactions;
            valueOfTransactions += value;
            continue;
        }

        if (valueOfTransactions + value >= 0) {
            minHeapNegativeValuesIncludedInTransactions.enqueue(value);
            ++numberOfTransactions;
            valueOfTransactions += value;
            continue;
        }

        if (!minHeapNegativeValuesIncludedInTransactions.isEmpty() && minHeapNegativeValuesIncludedInTransactions.front() < value) {
            valueOfTransactions += value - minHeapNegativeValuesIncludedInTransactions.dequeue();
            minHeapNegativeValuesIncludedInTransactions.enqueue(value);
        }
    }

    return numberOfTransactions;
};
