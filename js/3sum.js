/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function (nums) {
  const result = [];
  nums.sort((a, b) => a - b);

  for (let i = 0; i < nums.length; i++) {
    let l = i + 1;
    let r = nums.length - 1;

    if (i > 0 && nums[i] === nums[i - 1]) {
      continue;
    }

    while (nums[i] <= 0 && l < r) {
      if (nums[i] + nums[l] + nums[r] > 0) {
        r--;
      } else if (nums[i] + nums[l] + nums[r] < 0) {
        l++;
      } else {
        result.push([nums[i], nums[l], nums[r]]);
        l++;
        while (nums[l] === nums[l - 1] && l < r) {
          l++;
        }
      }
    }
  }
  return result;
};

console.log(
  threeSum([2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10])
);
