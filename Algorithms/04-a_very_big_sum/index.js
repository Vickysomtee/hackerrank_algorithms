var number = [1000000001, 1000000002, 1000000003, 1000000004, 1000000005];

var aVeryBigSum = (arr) => {
  var result = 0;

  for (var i = 0; i < arr.length; i++) {
    result += arr[i];
  }

  console.log(result);
};

aVeryBigSum(number)
