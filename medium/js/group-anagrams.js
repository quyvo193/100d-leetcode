/**
 * @param {string[]} strs
 * @return {string[][]}
 */
var groupAnagrams = function (strs) {
  const anagramGroup = new Map();

  strs.forEach((s) => {
    const count = Array.from({ length: 26 }).fill(0);

    for (let i = 0; i < s.length; i++) {
      count[s.charCodeAt(i) - "a".charCodeAt()]++;
    }

    const key = count.join("-");

    anagramGroup.has(key)
      ? anagramGroup.get(key).push(s)
      : anagramGroup.set(key, [s]);
  });

  return Array.from(anagramGroup.values());
};

console.log(groupAnagrams(["bdddddddddd", "bbbbbbbbbbc"]));
