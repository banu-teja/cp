function stringShift(s: string, shift: number[][]): string {
  let netShift = 0;
  for (const [a, b] of shift) {
    netShift += a === 0 ? -b : b;
  }
  netShift %= s.length;
  if (netShift < 0) {
    netShift += s.length;
  }
  return s.slice(-netShift) + s.slice(0, -netShift);
}

console.log(
  stringShift("abc", [
    [0, 1],
    [1, 2],
  ])
);

(() => {
  const testCases = [
    {
      name: "Regular dominoes",
      func: stringShift,
      tests: [
        {
          s: "abc",
          shift: [
            [0, 1],
            [1, 2],
          ],
          expected: "cab",
        },
        {
          s: "abcdefg",
          shift: [
            [1, 1],
            [1, 1],
            [0, 2],
            [1, 3],
          ],
          expected: "efgabcd",
        },
      ],
    },
  ];

  testCases.forEach((testCase) => {
    console.log(`Testing ${testCase.name}:`);
    testCase.tests.forEach((test, index) => {
      const result = testCase.func(test.s, test.shift);
      const passed = result === test.expected;
      console.log(`Test ${index + 1}: ${passed ? "PASSED" : "FAILED"}`);
      console.log(`  Input: tops = [${test.s}], bottoms = [${test.shift}]`);
      console.log(`  Expected: ${test.expected}, Got: ${result}`);
      if (!passed) {
        console.log(`  ❌ Test failed!`);
      }
      console.log();
    });
    console.log();
  });
})();
