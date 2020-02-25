var findMedianSortedArrays = function(nums1, nums2) {
  var merged = [...nums1, ...nums2];
  if (merged.length === 1) {
    return merged[0];
  }
  var sorted = merged.sort((a, b) => a - b);
  var len = Math.floor(sorted.length / 2);
  if (sorted.length % 2 === 0) {
    return (sorted[len - 1] + sorted[len]) / 2;
  }
  return sorted[len];
};

console.log(
  findMedianSortedArrays([23, 53, 2, 4, 3], [2, 3, 2, 3, 2, 423, 52, 34])
);
