/**
 * Forward declaration of guess API.
 * @param {number} num   your guess
 * @return 	            -1 if num is lower than the guess number
 *			             1 if num is higher than the guess number
 *                       otherwise return 0
 * var guess = function(num) {}
 */

/**
 * @param {number} n
 * @return {number}
 */
var guessNumber = function (n) {
  let low = 1;
  let high = n;
  while (low < high) {
    let mid = Math.floor((high - low) / 2) + low;
    let ret = guess(mid);
    if (ret === 0) {
      return mid;
    } else if (ret === -1) {
      high = mid - 1;
    } else {
      low = mid + 1;
    }
  }
  return low;
};
