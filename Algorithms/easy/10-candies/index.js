const ratings = [2, 4, 2, 6, 1, 7, 8, 9, 2, 1];

const candies = (n, arr) => {
    let candies = Array(n).fill(1)

     for(let i = 1; i < n; i++) {
        if(arr[i] > arr[i-1]) {
            candies[i] = candies[i-1] + 1;
        }
    }

    for(let i = n-2; i >= 0; i--) {
        if(arr[i] > arr[i+1]) {
            candies[i] = Math.max(candies[i], candies[i+1] + 1)
        }
    }

    console.log(candies)
}

candies(ratings.length, ratings)