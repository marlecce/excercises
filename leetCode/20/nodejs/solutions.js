const isValid = function (s) {
  const map = {
    '(': ')',
    '[': ']',
    '{': '}'
  }
  const stack = []
  for (let i = 0; i < s.length; i++) {
    if (map[s[i]]) {
      stack.push(map[s[i]])
    } else if (s[i] !== stack.pop()) {
      return false
    }
  }
  return stack.length === 0
}

const s1 = '()'
console.log(isValid(s1))

const s2 = '(]'
console.log(isValid(s2))
