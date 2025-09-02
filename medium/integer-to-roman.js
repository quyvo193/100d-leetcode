/**
 * @param {number} num
 * @return {string}
 */
var intToRoman = function (num) {
  const map = {
    1: "I",
    5: "V",
    10: "X",
    50: "L",
    100: "C",
    500: "D",
    1000: "M",
  };

  const romanVal = [1000, 500, 100, 50, 10, 5, 1];
  const roman = [];
  let i = 0;

  while (num) {
    let x = num / romanVal[i];

    if (x === 0) {
      i++;
    } else {
      for (let j = 0; j < x; j++) {
        roman.push(map[romanVal[j]]);
      }
      num -= x * romanVal[i];
      i++;
    }
  }

  return roman.join("");
};
