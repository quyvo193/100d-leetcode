/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function (s, t) {
  if (s.length !== t.length) {
    return false;
  }

  const arr = Array.from({ length: 26 }).fill(0);

  for (let c of s) {
    arr[c.charCodeAt() - "a".charCodeAt()]++;
  }

  for (let c of t) {
    const pos = c.charCodeAt() - "a".charCodeAt();

    arr[pos]--;
    if (arr[pos] < 0) {
      return false;
    }
  }

  return true;
};

console.log(isAnagram("anagram", "nagaram"));
