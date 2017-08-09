/**
    Given an array of integers, return indices of the two numbers such that they add up to a specific target.
    You may assume that each input would have exactly one solution.
    Example:
    Given nums = [2, 7, 11, 15], target = 9,

    Because nums[0] + nums[1] = 2 + 7 = 9,
    return [0, 1].

    1. 在数组中找打两个数字之和等于给定的数字
    2. 不能重复用数据中的同一个数字进行相加
    3. 每个输入都有一种解决方案，所以不用考虑有多种的情况
 */
var twoSumAnt = function (arr, target) {
  for (var i = 0; i < arr.length; i++) {
    for (var j = i + 1; j < arr.length; j++) {
      if (arr[i] + arr[j] == target) {
        return [i, j];
      }
    }
  }
};

var twoSumBear = function (arr, target) {
  var result = [];

  for (var i = 0; i < arr.length; i++) {
    var current = arr[i];

    if (result[target - current] >= 0) {
      return [result[target - current], i];
    } else {
      result[current] = i;
    }
  }
};

module.exports = {
  ant: twoSumAnt,
  bear: twoSumBear
}