/**
 * @param {number[]} heights
 * @return {number}
 */
var heightChecker = function(heights) {
  function sortNumber(a, b) {
    return a - b;
  }
  var h = [...heights];
  heights.sort(sortNumber);
  var c = 0;
  for (var i = 0, l = heights.length; i < l; i++) {
    if (h[i] !== heights[i]) {
      c++;
    }
  }
  return c;
};
