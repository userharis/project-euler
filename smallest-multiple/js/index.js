function smallestMultiple(min, max) {
  let multiple = 1;
  for (; min <= max; min++) {
    multiple = lcm(multiple, min);
  }
  return multiple;
}

// least common multiple of two numbers
function lcm(a, b) {
  return (a * b) / gcd(a, b);
}

// greatest common divisor
function gcd(a, b) {
  const r = a % b;
  return r === 0 ? b : gcd(b, r);
}

console.log(smallestMultiple(1, 20));
