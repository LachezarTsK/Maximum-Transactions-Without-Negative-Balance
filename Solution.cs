
using System;
using System.Collections.Generic;

public class Solution
{
    public int MaxTransactions(int[] transactions)
    {
        long valueOfTransactions = 0;
        int numberOfTransactions = 0;
        PriorityQueue<int, int> minHeapNegativeValuesIncludedInTransactions = new PriorityQueue<int, int>();

        foreach (int value in transactions)
        {

            if (value >= 0)
            {
                ++numberOfTransactions;
                valueOfTransactions += value;
                continue;
            }

            if (valueOfTransactions + value >= 0)
            {
                minHeapNegativeValuesIncludedInTransactions.Enqueue(value, value);
                ++numberOfTransactions;
                valueOfTransactions += value;
                continue;
            }

            if (minHeapNegativeValuesIncludedInTransactions.Count > 0 && minHeapNegativeValuesIncludedInTransactions.Peek() < value)
            {
                valueOfTransactions += value - minHeapNegativeValuesIncludedInTransactions.Dequeue();
                minHeapNegativeValuesIncludedInTransactions.Enqueue(value, value);
            }
        }

        return numberOfTransactions;
    }
}
