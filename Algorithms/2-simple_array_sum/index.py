import array as arr

numbers = [2, 3, 4]

# Try out with an array
# numbers = arr.array('i', [2, 3, 4])


def simple_array_sum(arr):
    sum = 0

    for num in numbers:
        sum += num

    print(sum)


simple_array_sum(numbers)
