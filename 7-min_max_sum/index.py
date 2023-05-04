
# Test case 1
numbers = [7, 69, 2, 221, 8974]

# Test case 2
# numbers = [1, 2, 3, 4, 5]


def min_max_sum(arr):

    min = 0
    max = 0

    arr.sort()

    for x in range(len(arr)):
        if (x < len(arr) - 1):
            min += arr[x]
        if (x > 0):
            max += arr[x]

    print(min, max)        


min_max_sum(numbers)
