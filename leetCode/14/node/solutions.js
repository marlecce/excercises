const longestCommonPrefix = function (strs) {
  if (strs.length === 0) return ''
  if (strs.length === 1) return strs[0]

  let prefix = strs[0]

  for (let i = 1; i < strs.length; i++) {
    while (strs[i].indexOf(prefix) !== 0) {
      prefix = prefix.slice(0, -1)
      if (prefix === '') {
        return ''
      }
    }
  }

  return prefix
}

const s1 = ['flower', 'flow', 'flight']
console.log(longestCommonPrefix(s1))

const s2 = ['dog', 'racecar', 'car']
console.log(longestCommonPrefix(s2))
