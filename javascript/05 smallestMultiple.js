/** @format */

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to n?

const sieve = max => {
  const numbers = new Array(max + 1);
  for (let i = 2; i <= max; i++) numbers[i] = true;
  for (let i = 2; i * i <= max; i++) {
    if (numbers[i]) {
      for (let j = i * 2; j <= max; j += i) {
        numbers[j] = false;
      }
    }
  }

  const primes = [];
  numbers.forEach((e, i) => (e ? primes.push(i) : null));
  return primes;
};

const smallestMult = n => {
  const primes = sieve(n);
  let product = 1;

  primes.forEach(e => {
    let temp = 1;
    while (temp <= n / e) {
      temp *= e;
    }
    product *= temp;
    temp = 1;
  });

  return product;
};

console.log(smallestMult(20)); // 60
