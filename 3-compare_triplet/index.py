alice = [17, 28, 30]
bob = [99, 16, 8]

# Second test case
# alice = [6, 7, 8]
# bob = [9, 7, 4]


def compare_triplet(a, b):
    alice = 0
    bob = 0

    for i in range(3):
        if a[i] > b[i]:
            alice += 1
        elif a[i] < b[i]:
            bob += 1

    result = [alice, bob]

    print(result);

compare_triplet(alice, bob)        
