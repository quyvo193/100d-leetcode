/**
 * @param {number[]} nums
 * @return {number[]}
 */
var productExceptSelf = function (nums) {
  let indexZero = [];

  nums.forEach((num, i) => {
    if (num === 0) indexZero.push(i);
  });

  if (indexZero.length > 1) {
    return Array.from({ length: nums.length }).fill(0);
  }

  if (indexZero.length === 1) {
    const result = Array.from({ length: nums.length }).fill(0);

    result[indexZero[0]] = nums.reduce((p, num, i) => {
      if (i === indexZero[0]) {
        return p;
      }
      return p * num;
    }, 1);

    return result;
  }

  const productOfAll = nums.reduce((p, num) => p * num, 1);

  return nums.map((num) => productOfAll / num);
};

console.log(productExceptSelf([-1, 1, 0, -3, 3]));
