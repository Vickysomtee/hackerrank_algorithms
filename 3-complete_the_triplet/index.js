const alice = [17, 28, 30];
const bob = [99, 16, 8];

//Second test case
// const alice = [6, 7, 8];
// const bob = [9, 7, 4];

function compareTriplets(a, b) {
  let alice = 0;
  let bob = 0;
  for (var i = 0; i < 3; i++) {
    if (a[i] > b[i]) {
      alice += 1;
    } else if (a[i] < b[i]) {
      bob += 1;
    }
  }

  const result = [alice, bob];

  console.log(result);
}

compareTriplets(alice, bob);
