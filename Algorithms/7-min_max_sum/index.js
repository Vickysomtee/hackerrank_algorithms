// Text case 1
var arr = [7, 69, 2, 221, 8974];

// Test case 2
// var arr = [1, 2, 3, 4, 5];

function minMaxSum(arr) {
  arr.sort((a, b) => a - b);

  min = 0;
  max = 0;

  for (var i = 0; i < arr.length; i++) {
    if (i < arr.length - 1) {
      min += arr[i];
    }

    if (i > 0) {
      max += arr[i];
    }
  }

  // Solution 2 - Got this online
  
  //   const max = arr
  //     .sort((a, b) => b - a)
  //     .slice(0, 4)
  //     .reduce((a, b) => a + b, 0);

  //   const min = arr
  //     .sort((a, b) => b - a)
  //     .reverse()
  //     .slice(0, 4)
  //     .reduce((a, b) => a + b, 0);

  console.log(min, max);
}

minMaxSum(arr);
