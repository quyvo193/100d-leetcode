/**
 * @param {number[]} nums
 * @return {number[]}
 */
var productExceptSelf = function (nums) {
  const n = nums.length;
  const result = Array.from({ length: n }).fill(1);

  for (let i = 1; i < n; i++) {
    result[i] = result[i - 1] * nums[i - 1];
  }

  let prevSuffix = 1;
  for (let i = n - 2; i >= 0; i--) {
    prevSuffix *= nums[i + 1];
    result[i] *= prevSuffix;
  }

  return result;
};

console.log(productExceptSelf([-1, 1, 0, -3, 3]));
