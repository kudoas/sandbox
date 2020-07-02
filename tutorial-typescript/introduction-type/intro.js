// Please referd to this issue: https://qiita.com/uhyo/items/e2fdef2d3236b9bfe74a
var __spreadArrays = (this && this.__spreadArrays) || function () {
    for (var s = 0, i = 0, il = arguments.length; i < il; i++) s += arguments[i].length;
    for (var r = Array(s), k = 0, i = 0; i < il; i++)
        for (var a = arguments[i], j = 0, jl = a.length; j < jl; j++, k++)
            r[k] = a[j];
    return r;
};
// premitive
// hoge: 型注釈
var a = 2;
var b = "foo";
// --strinctNullChecksがオンにしない場合はnull, undefinedは他の型の値として扱える
var c = null;
// const d: string = c;
// literal
var e = "foo";
// stringの一部
var ee = e;
// 'bar'しか許容しない型
// const ee: "bar" = "foo";
// leteralと型推論: constのみ
var aa = "foo";
// const bb: 'bar' = aa error!
// letの場合はpremitive
var aaa = "foo";
var bb = aaa;
var me = {
    name: "kudoa",
    age: 22
};
// Preson2はPreson型変数を含んでいるのでOK！
var you = me;
// オブジェクトリテラルの場合は余計なプロパティを持つオブジェクトは弾かれる
// const her: Person2 = { name: "Miyu", age: 32 };
// 配列型
var arr = [0, 1, 2, 3, 4];
// arr.push('hoge'); error
arr.push(5);
// 関数型
// const f: (foo: string) => number = func;
var funcA = function (arg) {
    return Number(arg);
};
// const fucn: 型　= () => {}
var f1 = function (foo) {
    return foo + "hoge";
};
var f2 = f1;
// 可変長引数
var f3 = function (foo) {
    var bar = [];
    for (var _i = 1; _i < arguments.length; _i++) {
        bar[_i - 1] = arguments[_i];
    }
    return bar;
};
// f3("hoge", 1, 2, 3, 3); 1,2,3,3
// void型
var z = undefined;
// const zz: undefined = z; error
// class
var Foo = /** @class */ (function () {
    function Foo() {
        this.method = function (foo) {
            console.log(foo);
        };
        // こっちでもかける！
        // method = (foo: string): void => {
        //   console.log(foo);
        // };
    }
    return Foo;
}());
// Foo型を定義: 構造的型付け
var obj = new Foo();
var obj2 = new Foo();
var obj1 = {
    foo: 12,
    bar: "hage"
};
// class
var FHoo = /** @class */ (function () {
    function FHoo(obj) {
    }
    return FHoo;
}());
var goo = new FHoo("foo");
// function
var ff = function (arg) {
    return arg;
};
// <T>(obj: T)=> voidという型変数の残った型になる！
var func = ff;
func(3);
// 型推論が効く
var indentify = function (value) {
    return value;
};
var value = indentify(3);
// const str: string = value error
// tuple型: JSにはないから配列を代わりに使用する
var tup = ["foo", 5];
var str = tup[0];
// ただしタプルでもただの配列なので普通に操作できる点に注意
var makePair1 = function (x, y) {
    return [x, y];
};
var makePair2 = function (x, y) {
    return [x, y];
};
var a1 = [3, "hgoe", "hoge"];
var Func = function () {
    var args = [];
    for (var _i = 0; _i < arguments.length; _i++) {
        args[_i] = arguments[_i];
    }
    return args[1];
};
var v = Func("foo", 3, true);
var fn = function (f) {
    var args = [];
    for (var _i = 1; _i < arguments.length; _i++) {
        args[_i - 1] = arguments[_i];
    }
    return args[0];
};
// タプル型と可変長引数とジェネリクス
function bind2(func, value) {
    return function () {
        var args = [];
        for (var _i = 0; _i < arguments.length; _i++) {
            args[_i] = arguments[_i];
        }
        return func.apply(void 0, __spreadArrays([value], args));
    };
}
var ff2 = function (arg) {
    return arg;
};
var add = function (x, y) { return x + y; };
var add1 = bind2(add, 1);
// union型
var value3 = "foo";
value3 = 100;
value3 = "bar";
// ただしこのままだとbarプロパティは参照できない
var oj = {
    foo: "hoge",
    bar: 12,
    baz: 21
};
var ffunc = function (value) {
    if ("string" == typeof value) {
        // valueはstring型なのでlengthプロパティを参照できる！
        return value.length;
    }
    else {
        // number型
        return value;
    }
};
// null check
var stringOrNull = function (value) {
    // if (value != null) {
    //   return value.length;
    // } else {
    //   return 0;
    // }
    return value != null ? value.length : 0;
};
var map = function (obj, f) {
    if (obj.type == "Some") {
        return {
            type: "Some",
            value: f(obj.value)
        };
    }
    else {
        return {
            type: "None"
        };
    }
};
