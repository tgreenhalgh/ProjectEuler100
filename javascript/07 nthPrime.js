/** @format */
const isPrime = n => {
  if (n < 2) return false;
  for (let i = 2; i * i <= n; i++) if (n % i == 0) return false;
  return true;
};

const nthPrime = n => {
  let count = 1;
  let number = 1;
  while (count <= n) {
    if (isPrime(++number)) count++;
  }
  return number;
};

console.log(nthPrime(6)); // 13
console.log(nthPrime(10)); // 29
console.log(nthPrime(100)); // 541
console.log(nthPrime(1000)); // 7919
console.log(nthPrime(10001)); // 104743
