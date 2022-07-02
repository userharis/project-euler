// returns sum of all multiples of 3 and 5 less than <number>
function sumOfMultiples(number) {
  let sum = 0;
  for (let i = 0; i < number; i += 3) {
    sum += i;
  }
  for (let i = 0; i < number; i += 5) {
    // continue if i is multiple of 3
    if (i % 3 === 0) continue;
    sum += i;
  }
  return sum;
}

console.log(sumOfMultiples(1000));
