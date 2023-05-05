function staircase(n) {
  var string = "";
  for (var i = 0; i < n; i++) {
    for (var j = 1; j < n + 1; j++) {
      if (j < n - i) {
        string += " ";
      } else {
        string += "#";
      }
    }
    console.log(string);
    string = "";
  }
}

staircase(6);
