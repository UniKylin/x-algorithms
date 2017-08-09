/**
 * 
 * Write a function that takes a string as input and returns the string reversed.
 * Example:
 *  Given s = "hello", return "olleh". 
 * 
 * 翻转字符串
 */
var reverseString = function (str) {
  var result = '';
  var arr = str.split('');

  for (var i = arr.length - 1; i >= 0; i--) {
    result += arr[i];
  }

  return result;
};

console.log(reverseString('hello node.js'));

/**
 * 第二种：入栈和出栈解决，感觉有点大材小用
 * @param {*} str 
 */
var reverseStringAnt = function (str) {
  var temp = [],
    result = [],
       arr = str.split('');
  
  for (var i = 0; i < arr.length; i++) {
    temp.push(arr[i]);
  } 

  for (var j = 0; j < arr.length; j++) {
    result.push(temp.pop());
  }

  return result.join('');
};

console.log(reverseStringAnt('hello node.js'));

/**
 * 第三种：api
 * 但是这个效率不高
 */
var reverseStringBear = function (str) {
  return str.split('').reverse().join('');
};

console.log(reverseStringBear('hello node.js'));

/**
 * 第四种：截取字符的方式和第一种差不多
 * 但是这种是所有里面效率最高的
 */
var reverseStringCat = function (str) {
  var result = '';

	for (var i = str.length - 1; i >= 0; i--) {
		result += str.charAt(i);
	}

	return result;
};

console.log(reverseStringCat('hello vue'));