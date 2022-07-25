var minWindow = function (s, t) {
  let need = new Map(),
    window = new Map();
  let left = 0,
    right = 0;

  let valid = 0,
    start = 0;
  let len = Infinity;

  for (let i = 0; i < t.length; i++) {
    let curr = t[i];
    if (need.has(curr)) {
      let v = need.get(curr);
      need.set(curr, ++v);
    } else {
      need.set(curr, 1);
    }
  }

  while (right < s.length) {
    let c = s[right];
    right++;

    if (need.has(c)) {
      if (window.has(c)) {
        let v = window.get(c);
        window.set(c, ++v);
      } else {
        window.set(c, 1);
      }
      if (window.get(c) == need.get(c)) {
        valid++;
      }
    }
    while (valid === need.size) {
      if (right - left <= len) {
        start = left;
        len = right - left;
      }

      let d = s[left];
      left++;

      if (need.has(d)) {
        let v = window.get(d);
        if (window.get(d) === need.get(d)) {
          valid--;
        }
        window.set(d, --v);
      }
    }
  }

  if (len === Infinity) {
    return "";
  } else {
    return s.substr(start, len);
  }
};

console.log(minWindow("ADOBECODEBANC", "ABC"));
