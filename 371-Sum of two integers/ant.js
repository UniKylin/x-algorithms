/**
 *  LeetCode 371. Sum of Two Integers
 *  Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.
 */

function getSum(ant, bear) {

  if (bear === 0) {
    return ant;
  }

  if (ant === 0) {
    return bear;
  }

  while (bear != 0) {
    var flag = ant & bear;    // 进位值
    ant = ant ^ bear;         // 相加
    bear = flag << 1;         // 进位
  }
  
  return ant;
}

module.exports = getSum;