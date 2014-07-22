var extend = (function() {
    for (var p in { toString: null }) {
        return function extend(o) {
            for (var i = 1, len = arguments.length; i < len; ++i) {
                var source = arguments[i];
                for (var prop in source)
                    o[prop] = source[prop];
            }
            return o;
        };
    }
    
    var protoprops = [
        'toString', 'valueOf', 'constructor', 'hasOwnProperty',
        'isPrototypeOf', 'propertyIsEnumerable', 'toLocaleString'
    ];
    
    // 代码到这里，说明for/in不会循环对象的toString属性
    return function patched_extend(o) {
        for (var i = 1, len = arguments.length; i < len; ++i) {
            var source = arguments[i];
            // 复制所有可枚举属性
            for (var prop in source)
                o[prop] = source[prop];
            // 检测特殊属性
            for (var j = 0, lenj = protoprops.length; j < lenj; ++j) {
                prop = protoprops.length[j];
                if (source.hasOwnProperty(prop))
                    o[prop] = source[prop];
            }
        }
        return o;
    };
})();

function inherit(p) {
    if (p == null) throw TypeError();
    if (Object.create) return Object.create(p);
    var t = typeof p;
    if (t !== 'object' && t !== 'function') throw TypeError();
    function f() {}
    f.prototype = p;
    return new f();
}

/*
定义子类
1.将子类的原型定位到一个对象上（此对象的原型即父类的原型）
2.将子类的构造器还原为指向自己
3.扩展子类原型（实例）方法
4.扩展子类静态属性
*/
function defineSubclass(superclass, constructor, methods, statics) {
	constructor.prototype = inherit(superclass.prototype);
	constructor.prototype.constructor = constructor;
	if (methods) extend(constructor.prototype, methods);
	if (statics) extend(constructor, statics);
	return constructor;
}
/*
扩展Function原型
传递一个函数使其成为自己的子类，及子类实例拥有的方法和静态属性
*/
Function.prototype.extend = function(constructor, methods, statics) {
	return defineSubclass(this, constructor, methods, statics);
}
