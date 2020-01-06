const largestPrimeFactor = x => {
  let maxPrimeFactor = 0;

  // parse out the evens
  while (x % 2 == 0) {
    maxPrimeFactor = 2;
    x /= 2;
  }

  // parse out the odds
  for (let i = 3; i <= Math.sqrt(x); i += 2) {
    while (x % i == 0) {
      maxPrimeFactor = i;
      x /= i;
    }
  }

  // x is prime at this point
  if (x > 2) maxPrimeFactor = x;

  return maxPrimeFactor;
};

console.log(largestPrimeFactor(2))            // 2
console.log(largestPrimeFactor(3))            // 3
console.log(largestPrimeFactor(5))            // 5
console.log(largestPrimeFactor(7))            // 7
console.log(largestPrimeFactor(13195))        // 29
console.log(largestPrimeFactor(600851475143)) // 6857
console.log(largestPrimeFactor(1));