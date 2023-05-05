number = [1000000001, 1000000002, 1000000003, 1000000004, 1000000005]

def a_very_big_sum (arr):
    result = 0

    for x in range (len(arr)):
        result += arr[x]

    print(result)    

a_very_big_sum(number)