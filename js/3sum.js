/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function (nums) {
  const result = [];

  nums = Array.from(new Set(nums));
  if (nums.length < 3) {
    nums = [0, 0, 0];
  }

  for (let i = 0; i < nums.length; i++) {
    let x = i;
    let y = i + 1;
    let z = i + 2;

    while (y < nums.length - 1) {
      while (z < nums.length) {
        if (nums[x] + nums[y] + nums[z] === 0) {
          result.push([nums[x], nums[y], nums[z]]);
        }
        z++;
      }

      y++;
      z = y + 1;
    }

    return result;
  }
};
