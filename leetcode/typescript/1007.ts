function minDominoRotations(tops: number[], bottoms: number[]): number {
  const countRotations = (target: number) => {
    let rotationsTop = 0;
    let rotationsBottom = 0;

    for (let i = 0; i < tops.length; i++) {
      if (tops[i] !== target && bottoms[i] !== target) return -1;

      if (tops[i] !== target) rotationsTop++;
      if (bottoms[i] !== target) rotationsBottom++;
    }

    return Math.min(rotationsTop, rotationsBottom);
  };

  const rotationsForTop = countRotations(tops[0]);
  if (rotationsForTop !== -1) return rotationsForTop;

  return countRotations(bottoms[0]);
}

function minDominoRotationsLargeRange(
  tops: number[],
  bottoms: number[]
): number {
  const uniqueValues = new Set([tops[0], bottoms[0]]);

  const countRotations = (target: number): number => {
    let rotationsTop = 0;
    let rotationsBottom = 0;

    for (let i = 0; i < tops.length; i++) {
      if (tops[i] !== target && bottoms[i] !== target) return -1;
      if (tops[i] !== target) rotationsTop++;
      if (bottoms[i] !== target) rotationsBottom++;
    }

    return Math.min(rotationsTop, rotationsBottom);
  };

  let minRotations = -1;
  for (const target of uniqueValues) {
    const rotations = countRotations(target);
    if (rotations !== -1) {
      minRotations =
        minRotations === -1 ? rotations : Math.min(minRotations, rotations);
    }
  }

  return minRotations;
}

(() => {
  const testCases = [
    {
      name: "Regular dominoes",
      func: minDominoRotations,
      tests: [
        { tops: [2, 1, 2, 4, 2, 2], bottoms: [5, 2, 6, 2, 3, 2], expected: 2 },
        { tops: [3, 5, 1, 2, 3], bottoms: [3, 6, 3, 3, 4], expected: -1 },
      ],
    },
    {
      name: "Large dominoes",
      func: minDominoRotationsLargeRange,
      tests: [
        { tops: [2, 1, 2, 4, 2, 2], bottoms: [5, 2, 6, 2, 3, 2], expected: 2 },
        { tops: [3, 5, 1, 2, 3], bottoms: [3, 6, 3, 3, 4], expected: -1 },
      ],
    },
  ];

  testCases.forEach((testCase) => {
    console.log(`Testing ${testCase.name}:`);
    testCase.tests.forEach((test, index) => {
      const result = testCase.func(test.tops, test.bottoms);
      const passed = result === test.expected;
      console.log(`Test ${index + 1}: ${passed ? "PASSED" : "FAILED"}`);
      console.log(
        `  Input: tops = [${test.tops}], bottoms = [${test.bottoms}]`
      );
      console.log(`  Expected: ${test.expected}, Got: ${result}`);
      if (!passed) {
        console.log(`  ❌ Test failed!`);
      }
      console.log();
    });
    console.log();
  });
})();
