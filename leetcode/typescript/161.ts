function isOneEditDistance(s: string, t: string): boolean {
  if (Math.abs(s.length - t.length) > 1) return false;

  if (s.length > t.length) [s, t] = [t, s];

  let i = 0,
    j = 0;
  let diffCount = 0;

  while (i < s.length && j < t.length) {
    if (s[i] !== t[j]) {
      diffCount++;
      if (diffCount > 1) return false;

      if (s.length === t.length) i++;
    } else {
      i++;
    }
    j++;
  }

  if (diffCount === 0 && s.length !== t.length) diffCount++;
  return diffCount === 1;
}

console.log(isOneEditDistance("ab", "acb")); // true
console.log(isOneEditDistance("cab", "ad")); // false
console.log(isOneEditDistance("1203", "1213")); // true
console.log(isOneEditDistance("", "a")); // true
console.log(isOneEditDistance("a", "a")); // false
console.log(isOneEditDistance("abcde", "abXde")); // true
console.log(isOneEditDistance("abcde", "abcde")); // false
console.log(isOneEditDistance("abcde", "abcd")); // true
console.log(isOneEditDistance("abcde", "abcdef")); // true
