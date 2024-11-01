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

(() => {
  const testCases = [
    {
      name: "One Edit Distance",
      func: isOneEditDistance,
      tests: [
        { s: "ab", t: "acb", expected: true },
        { s: "cab", t: "ad", expected: false },
        { s: "1203", t: "1213", expected: true },
        { s: "", t: "a", expected: true },
        { s: "a", t: "a", expected: false },
        { s: "abcde", t: "abXde", expected: true },
        { s: "abcde", t: "abcde", expected: false },
        { s: "abcde", t: "abcd", expected: true },
        { s: "abcde", t: "abcdef", expected: true },
      ],
    },
  ];

  testCases.forEach((testCase) => {
    console.log(`Testing ${testCase.name}:`);
    testCase.tests.forEach((test, index) => {
      const result = testCase.func(test.s, test.t);
      const passed = result === test.expected;
      console.log(`Test ${index + 1}: ${passed ? "PASSED" : "FAILED"}`);
      console.log(`  Input: s = "${test.s}", t = "${test.t}"`);
      console.log(`  Expected: ${test.expected}, Got: ${result}`);
      if (!passed) {
        console.log(`  ❌ Test failed!`);
      }
      console.log();
    });
    console.log();
  });
})();
