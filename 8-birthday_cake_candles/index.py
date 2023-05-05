
# Test case 1
candles = [3, 2, 1, 3]

# Test case 2
# candles = [18, 90, 90, 13, 90, 75, 90, 8, 90, 43]


def birthday_cake_candles(candles):
    max_candle = max(candles)
    count = 0

    for x in candles:
        if (x == max_candle):
            count += 1

    # for x in range(len(candles)):
    #     if (candles[x] == max_candle):
    #         count += 1

    print(count)


birthday_cake_candles(candles)
