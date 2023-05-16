

# Test case 1
time = "07:05:45PM"

# Test case 2
# time =  "12:45:54PM"


def time_conversion(s):
    meridiemPm = {
        "01": "13",
        "02": "14",
        "03": "15",
        "04": "16",
        "05": "17",
        "06": "18",
        "07": "19",
        "08": "20",
        "09": "21",
        "10": "22",
        "11": "23",
        "12": "12",
    }
    meridiemAm = {
        12: "00",
    }

    meridiem = s[len(s) - 2:len(s)]

    militaryHour = ""

    if (meridiem == 'PM'):
        militaryHour = meridiemPm[s[0:2]]
    elif (meridiem == 'AM' and s[0:2] == '12'):
        militaryHour = meridiemAm[s[0:2]]
    else:
        militaryHour = s[0:2]

    time = militaryHour + s[2:len(s) - 2]
    print(time)


time_conversion(time)
