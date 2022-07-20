// find the largest palindrome product from the
// product of two three digit numbers

function largestPalindromeProduct() {
  let result = 0;
  for (let i = 999; i > 99; i--) {
    for (let j = 999; j > 99; j--) {
      const product = i * j;
      const s = product.toString();
      const r = s.split("").reverse().join("");
      if (r === s && product > result) {
        result = product;
      }
    }
  }
  return result;
}

console.log(largestPalindromeProduct());
