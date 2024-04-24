const threeSum = (nums: number[]): number[][] => {
  const triplets: number[][] = []

  for (let i = 0; i < nums.length; i++) {
    for (let j = i + 1; j < nums.length; j++) {
      for (let k = j + 1; k < nums.length; k++) {
        if (nums.length < k) {
          break
        }
        if (nums[i] + nums[j] + nums[k] == 0) {
          const newArray = [nums[i], nums[j], nums[k]]
          const isInTheArray = triplets.some(array => array.sort((a, b) => a + b).toString() == newArray.sort((a, b) => a + b).toString())
          if (isInTheArray) {
            continue
          }
          triplets.push(newArray)
        }
      }
    }
  }

  return triplets
};



console.log(threeSum([-1, 0, 1, 2, -1, -4]))