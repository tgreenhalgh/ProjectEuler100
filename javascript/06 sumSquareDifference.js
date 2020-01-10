/** @format */
// Find the difference between the sum of the squares of the first n natural numbers and the square of the sum.

const sumSquareDifference = n => {
  let sum = 0;
  let sqSum = 0;

  for (let i = 1; i <= n; i++) {
    sum += i;
    sqSum += i * i;
  }

  return sum * sum - sqSum;
};

console.log(sumSquareDifference(100));
