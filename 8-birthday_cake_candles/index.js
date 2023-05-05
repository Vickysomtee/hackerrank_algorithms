// Test case 1
var candles = [3, 2, 1, 3];

// Test case 2
// var candles = [18, 90, 90, 13, 90, 75, 90, 8, 90, 43];

function birthdayCakeCandles(candles) {
  // Write your code here
  var max = Math.max(...candles);
  var count = 0;
  for (var i = 0; i < candles.length; i++) {
    if (candles[i] == max) {
      count++;
    }
  }

  // Using for of loop
  //   for (var candle of candles) {
  //     if (candle == max) {
  //       count++;
  //     }
  //   }

  console.log(count);
}

birthdayCakeCandles(candles);
