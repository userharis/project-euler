function sumOfEvenFibs(n) {
  let sum = 2;

  if (n < 2) return 0;

  for (
    let f1 = 0, f2 = 2, f3 = 8;
    f3 <= n;
    f1 = f2, f2 = f3, f3 = 4 * f2 + f1
  ) {
    sum += f3;
  }

  return sum;
}

console.log(sumOfEvenFibs(4000000));
