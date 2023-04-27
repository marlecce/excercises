const merge = function (nums1, m, nums2, n) {
  let i = m - 1
  let j = n - 1
  let k = m + n - 1

  while (i >= 0 && j >= 0) {
    if (nums1[i] >= nums2[j]) {
      nums1[k] = nums1[i]
      i--
    } else {
      nums1[k] = nums2[j]
      j--
    }
    k--
  }

  while (j >= 0) {
    nums1[k] = nums2[j]
    j--
    k--
  }

  return nums1
}

const nums1 = [1, 2, 3, 0, 0, 0]
const m1 = 3
const nums2 = [2, 5, 6]
const n2 = 3
console.log(merge(nums1, m1, nums2, n2))

const nums3 = [1]
const m3 = 1
const nums4 = []
const n4 = 0
console.log(merge(nums3, m3, nums4, n4))

const nums5 = [0]
const m5 = 0
const nums6 = [1]
const n6 = 1
console.log(merge(nums5, m5, nums6, n6))
