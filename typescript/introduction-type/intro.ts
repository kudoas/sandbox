// Please referd to this issue: https://qiita.com/uhyo/items/e2fdef2d3236b9bfe74a

// premitive
// hoge: 型注釈
const a: number = 2;
const b: string = "foo";

// --strinctNullChecksがオンにしない場合はnull, undefinedは他の型の値として扱える
const c: null = null;
// const d: string = c;

// literal
const e: "foo" = "foo";
// stringの一部
const ee: string = e;

// 'bar'しか許容しない型
// const ee: "bar" = "foo";

// leteralと型推論: constのみ
const aa = "foo";
// const bb: 'bar' = aa error!

// letの場合はpremitive
let aaa = "foo";
const bb: string = aaa;
// let aaa: "foo" = "foo"; 型注釈すればOK

// object型
interface Person {
  name: string;
  age: number;
}

const me: Person = {
  name: "kudoa",
  age: 22,
};
interface Person2 {
  name: string;
}
// Preson2はPreson型変数を含んでいるのでOK！
const you: Person2 = me;
// オブジェクトリテラルの場合は余計なプロパティを持つオブジェクトは弾かれる
// const her: Person2 = { name: "Miyu", age: 32 };

// 配列型
const arr: number[] = [0, 1, 2, 3, 4];
// arr.push('hoge'); error
arr.push(5);

// 関数型
// const f: (foo: string) => number = func;
const funcA = (arg: string) => {
  return Number(arg);
};

// const fucn: 型　= () => {}
const f1: (foo: string) => string = (foo) => {
  return foo + "hoge";
};
const f2: (foo: string, bar: number) => void = f1;

// 可変長引数
const f3 = (foo: string, ...bar: number[]) => bar;
// f3("hoge", 1, 2, 3, 3); 1,2,3,3

// void型
const z: void = undefined;
// const zz: undefined = z; error

// class
class Foo {
  method: (foo: string) => void = (foo) => {
    console.log(foo);
  };

  // こっちでもかける！
  // method = (foo: string): void => {
  //   console.log(foo);
  // };
}

// Foo型を定義: 構造的型付け
const obj: Foo = new Foo();

interface MyFoo {
  method: (foo: string) => void;
}

const obj2: MyFoo = new Foo();

// Generics: 型変数を使用できる
interface Hoo<S, T> {
  foo: S;
  bar: T;
}

const obj1: Hoo<number, string> = {
  foo: 12,
  bar: "hage",
};
// class
class FHoo<T> {
  constructor(obj: T) {}
}
const goo = new FHoo<string>("foo");
// function
const ff: <T>(obj: T) => T = (arg) => {
  return arg;
};
// <T>(obj: T)=> voidという型変数の残った型になる！
const func = ff;
func<number>(3);
// 型推論が効く
const indentify: <T>(value: T) => T = (value) => {
  return value;
};

const value = indentify(3);
// const str: string = value error

// tuple型: JSにはないから配列を代わりに使用する
const tup: [string, number] = ["foo", 5];
const str: string = tup[0];

// ただしタプルでもただの配列なので普通に操作できる点に注意
const makePair1 = (x: string, y: number): [string, number] => {
  return [x, y];
};
const makePair2: (x: string, y: number) => [string, number] = (x, y) => {
  return [x, y];
};

// 可変長引数のタプルも定義可能
type NumAndStrings = [number, ...string[]];
const a1: NumAndStrings = [3, "hgoe", "hoge"];

// タプル型と可変長引数
type Args = [string, number, boolean];
const Func = (...args: Args) => args[1];
const v = Func("foo", 3, true);

type Args1 = [string, ...number[]];
const fn = (f: string, ...args: Args) => args[0];

// タプル型と可変長引数とジェネリクス
function bind2<T, U extends any[], R>(
  func: (arg1: T, ...rest: U) => R,
  value: T
): (...args: U) => R {
  return (...args: U) => func(value, ...args);
}

const ff2: <T>(obj: T) => T = (arg) => {
  return arg;
};

const add = (x: number, y: number) => x + y;
const add1 = bind2(add, 1);

// union型
let value3: string | number = "foo";
value3 = 100;
value3 = "bar";
// value3 = true

interface I1 {
  foo: string;
  bar: number;
}
interface I2 {
  foo: number;
  baz: number;
}
// type型を定義して名前を付けられる！
type I12 = I1 | I2;
// ただしこのままだとbarプロパティは参照できない
const oj: I12 = {
  foo: "hoge",
  bar: 12,
  baz: 21,
};

const ffunc = (value: string | number): number => {
  if ("string" == typeof value) {
    // valueはstring型なのでlengthプロパティを参照できる！
    return value.length;
  } else {
    // number型
    return value;
  }
};

// null check
const stringOrNull = (value: string | null): number => {
  // if (value != null) {
  //   return value.length;
  // } else {
  //   return 0;
  // }
  return value != null ? value.length : 0;
};

interface Some<T> {
  type: "Some";
  value: T;
}
interface None {
  type: "None";
}
type Option<T> = Some<T> | None;

const map: <T, U>(obj: Option<T>, f: (obj: T) => U) => Option<U> = (obj, f) => {
  if (obj.type == "Some") {
    return {
      type: "Some",
      value: f(obj.value),
    };
  } else {
    return {
      type: "None",
    };
  }
};

// 頻出
const switchMap: <T, U>(obj: Option<T>, f: (obj: T) => U) => Option<U> = (obj, f) => {
  switch (obj.type) {
    case "Some":
      return {
        type: "Some",
        value: f(obj.value),
      };
    case "None":
      return {
        type: "None",
      };
  }
};
