const number = [2, 3, 4];

const simpleArraySum = (arr) => {
  let result = 0;

  for (const num of number) {
    result += num;
  }

  //   Solution 2
  //   for (var i = 0; i < arr.length; i++) {
  //     result += arr[i];
  //   }

  console.log(result);
};

simpleArraySum(number);
