const threeSum = (nums: number[]): number[][] => {
  const triplets: number[][] = [];
  const n = nums.length;

  nums.sort((a, b) => a - b);

  for (let i = 0; i < n - 2; i++) {
    if (i > 0 && nums[i] === nums[i - 1]) continue;

    let left = i + 1;
    let right = n - 1;

    while (left < right) {
      const sum = nums[i] + nums[left] + nums[right];

      if (sum === 0) {
        triplets.push([nums[i], nums[left], nums[right]]);

        while (left < right && nums[left] === nums[left + 1]) left++;
        while (left < right && nums[right] === nums[right - 1]) right--;

        left++;
        right--;
      } else if (sum < 0) {
        left++;
      } else {
        right--;
      }
    }
  }

  return triplets;
};

console.log(threeSum([-5, 1, -10, 2, -7, -13, -3, -8, 2, -15, 9, -3, -15, 13, -6, -10, 5, 6, 11, 1, 13, -12, 14, 6, 11, 4, 13, -14, 0, 11, 1, 10, -11, 6, -11, -10, 8, 2, -3, -13, -6, -9, -9, -6, 11, -8, -9, 1, 13, 9, 9, 3, 13, 0, -6, 1, -10, -15, 3, 5, 3, 11, -8, 0, 2, -11, 5, -13, 6, 9, -11, 7, 8, -13, 8, 4, -6, 14, 13, -15, 1, 7, -5, -1, -7, 5, 7, -2, -3, -13, 10, 7, 13, 9, -8, -8, 13, 12, -6, 4, 7, -10, -12, -8, -8, 11, 11, -6, 3, 9, -14, -11, 2, -4, -5, 10, 8, -13, -7, 12, -10, 10]));
