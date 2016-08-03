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

$.fn.serializeObject = function() {
    var o = {};
    var a = this.serializeArray();
    $.each(a, function() {
        if (o[this.name]) {
            if (!o[this.name].push) {
                o[this.name] = [o[this.name]];
            }
            o[this.name].push(this.value || '');
        } else {
            o[this.name] = this.value || '';
        }
    });
    return o;
};
