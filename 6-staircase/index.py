def staircase(n):

    # Solution 1

    # This is my own solution, the same flow I used in the Javascript and Go files
    # The next solution (Solution 2) is a different method I found online which is quite
    # shorter and readabale

    string = ""
    for x in range(n):
        for y in range(1, n + 1):

            if y < n - x:
                string += " "
            else:
                string += "#"

        print(string)
        string = ""

    # Solution 2

    # for stairs in range(1, n + 1):
    #     print(' ' * (n - stairs) + '#' * stairs)


staircase(6)
