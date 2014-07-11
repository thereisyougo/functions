function aaa(obj) {
	var sum = 0;
	var val = obj.value;
	var vals = [];
	for (var i = 0, len = val.length; i < len; ++i) {
		var curChar = val.charAt(i);
		if (/[^\x00-\xff]/.test(curChar))
			sum += 2;
		else
			sum += 1;
		vals.push(curChar);
		if (sum >= obj.length)
			break;
	}
	if (sum > obj.length)
		vals.pop();
		
	obj.value = vals.join('');
}
