/** @format */
/*
A palindromic number reads the same both ways. The largest palindrome
made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two n-digit numbers.
*/
const isPalindrome = number => {
  return (
    number ==
    +number
      .toString()
      .split('')
      .reverse()
      .join('')
  );
};

const largestPalindromeProduct = x => {
  const max = Math.pow(10, x) - 1;
  let maxProduct = 0;

  for (let i = max; i > max / 2; i--) {
    for (let j = max; j > max / 2; j--) {
      if (isPalindrome(i * j)) maxProduct = Math.max(maxProduct, i * j);
    }
  }
  return maxProduct;
};

console.log(largestPalindromeProduct(3));
