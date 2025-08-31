/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function (s) {
  let longest = "";

  for (let i = 0; i < s.length; i++) {
    let l = i;
    let r = i;

    while (l >= 0 && r < s.length && s[l] === s[r]) {
      if (r + 1 - l > longest.length) {
        longest = s.substring(l, r + 1);
      }

      l--;
      r++;
    }

    l = i;
    r = i + 1;
    while (l >= 0 && r < s.length && s[l] === s[r]) {
      if (r + 1 - l > longest.length) {
        longest = s.substring(l, r + 1);
      }

      l--;
      r++;
    }
  }

  return longest;
};
