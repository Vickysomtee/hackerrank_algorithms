const arr = [
  [11, 2, 4],
  [4, 5, 6],
  [10, 8, -12],
];

function diagonalDifference(arr) {
  let diagonalA = 0;
  let diagonalB = 0;

  for (var i = 0; i < arr.length; i++) {
    diagonalA += arr[i][i];
    diagonalB += arr[i][arr.length - i - 1];
  }
  console.log(Math.abs(diagonalA - diagonalB));
}

diagonalDifference(arr);
