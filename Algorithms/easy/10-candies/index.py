import functools

ratings = [2, 4, 2, 6, 1, 7, 8, 9, 2, 1]


def candies(n, arr):
    candies = [1 for i in range(n)]

    for i in range(1, len(arr)):
        if arr[i] > arr[i - 1]:
            candies[i] = candies[i - 1] + 1

    for i in reversed(range(n-1)):
        if arr[i] > arr[i + 1]:
            candies[i] = max(candies[i], candies[i + 1] + 1)        
    
    result = functools.reduce(lambda a, b: a + b, candies)

    print(result)

candies(len(ratings), ratings)