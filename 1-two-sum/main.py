def twoSum(nums, target):
  d = {}
  for i, n in enumerate(nums): 
    if n in d:
      return [d[n], i]
    d[target-n] = i

print(twoSum([1,2,5,7,6], 13))