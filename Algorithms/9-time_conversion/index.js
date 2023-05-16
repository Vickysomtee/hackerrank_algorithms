// Test case 1
const time = "07:05:45PM";

// Test case 2
// const time =  "12:45:54PM"

function timeConversion(s) {
  let meridiemPm = {
    "01": 13,
    "02": 14,
    "03": 15,
    "04": 16,
    "05": 17,
    "06": 18,
    "07": 19,
    "08": 20,
    "09": 21,
    "10": 22,
    "11": 23,
    "12": 12,
  };

  let meridiemAm = {
    12: "00",
  };

  let militaryHour = "";

  const meridiem = s.substring(s.length - 2, s.length);

  if (meridiem === "PM") {
    militaryHour = meridiemPm[s.substring(0, 2)];
  } else if (meridiem === "AM" && s.substring(0, 2)) {
    militaryHour = meridiemAm[s.substring(0, 2)];
  } else {
    militaryHour = s.substring(0, 2);
  }

  const militaryTime = militaryHour + s.substring(2, s.length - 2);

  console.log(militaryTime);
}

timeConversion(time);
