function largestPrimeFactor(n) {
  const sqrt = Math.sqrt(n);
  let largestPrimeFactor;

  /*
    since 2 is the only even prime number, 
    we run a separate loop for 2 and then
    run another loop only for odd numbers
  */

  if (n % 2 === 0) {
    largestPrimeFactor = 2;
    // make n smaller, remove all factors that are multiples of 2
    while (n % 2 === 0) {
      n /= 2;
    }
  }

  for (let d = 3, s = sqrt; d <= s; d += 2) {
    while (n % d === 0) {
      largestPrimeFactor = d;
      n /= d;
    }
  }

  // return n if its itself a prime number
  return largestPrimeFactor || n;
}

console.log(largestPrimeFactor(600851475140));
