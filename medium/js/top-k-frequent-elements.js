/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
var topKFrequent = function (nums, k) {
  const freqMap = new Map();
  const bucket = [];
  const result = [];

  for (let n of nums) {
    freqMap.set((freqMap.get(n) || 0) + 1);
  }

  for (let [num, freq] of freqMap) {
    bucket[freq] = (bucket[freq] || new Set()).add(num);
  }

  console.log(freqMap.values());
  console.log(bucket.map((item) => item));
};

topKFrequent([1, 1, 1, 2, 2, 3], 2);
