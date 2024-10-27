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
