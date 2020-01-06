/** @format */
const fiboEvenSum = x => {
  let sum = 0;
  const fibs = [1, 2];

  for (let i = 0; i < x - 2; i++) {
    fibs.push(fibs[i] + fibs[i + 1]);
  }
  fibs.forEach(x => (x % 2 ? null : (sum += x)));
  return sum;
};

console.log(fiboEvenSum(10)); // 44
console.log(fiboEvenSum(18)); // 3382
console.log(fiboEvenSum(23)); // 60696
console.log(fiboEvenSum(43)); // 350704366
