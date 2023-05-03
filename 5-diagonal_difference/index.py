arr = [
    [11, 2, 4],
    [4, 5, 6],
    [10, 8, -12],
]


def diagonal_difference(arr):
    diagonal_a = 0
    diagonal_b = 0

    for x in range(len(arr)):
        diagonal_a += arr[x][x]
        diagonal_b += arr[x][len(arr) - x - 1]

    print(abs(diagonal_a - diagonal_b))


diagonal_difference(arr)
