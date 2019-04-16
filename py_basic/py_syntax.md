# Basic of  "python"

コードは**誰にでも分かりやすくシンプルに書く**ことが大切

> テスト形式で自分で思い出す(答えを知りたかったら出力してみればOK！

##　学習上の注意点

1．学習のための学習にならないようにする(手段の目的化)

2．自分で思い出す作業をしてからどうしても分からない場合に調べる

3．使い方に詳しくなるより、どう使うかの視点を重視する(基礎より応用)

##　output

```python
import math
num = 10
print(num, type(num))
print("hi", "Mike", sep=",", end="\n") #hi,Mike end="\n"(改行)
print("hi", "Mike", sep=",", end="\n") #hi,Mike 
result = math.sqrt(25)
print(result)
print(help(math)) #使い方を確認できる

print('I don\'t know')  # \を入れて、読み込みは次の行からだと伝える
print("say \"I don't know.\"")
print("Hello.\nHow are you.")  # \nで改行する
print("C:aaaa\aaaa\aaaa")  # \が消えちゃう
print(r"C:aaaa\aaaa\aaaa")  # raw dataだと教えてあげればそのまま出力
print("""
line1
line2
line3
""")  # 縦でいい感じに出力されるが、改行されるので嫌なら"""に\
print("######")
print("""\
    line1
    line2
    line3\
    """)
print("######") #改行が消えて見やすく書ける(80文字以上続く場合は改行すべき)
s = "aaaaaaaaaaaaaaaaa"\
    "bbbbbbbbbbbbbbbbb"
print(s)
d = ("aaaaaaaaaaaaaaaa"
     "bbbbbbbbbbbbbbbb")
print(d) # "$$$$$""%%%%%%" の結合はコードを綺麗に書くために重要
# 文字列の掛け算も可能
print("="*10)
> ==========
```

##　method

```python
s = "My name is Yui. Hi, Yui."
is_start = s.startswith("My")　
print(is_start)
# True

is_start = s.startswith("Yui")　
print(is_start)
# False

# どこにある？
print(s.find("Yui"))
print(s.rfind("Yui"))
# 数は？
print(s.count("Yui"))
print(s.capitalize())
print(s.title())
print(s.upper())
print(s.lower())
print(s.replace("Yui","Dobugaeru"))
```

##　format

これはめちゃめちゃ便利

```python
>>> 'a is {}'.format('a')
'a is a'
>>> 'a is {}{}{}'.format('a','b','c')
'a is abc'
>>> 'My name is {name} {family}'.format(name = "Yui", family = "Kobayashi")
'My name is Yui Kobayashi'
>>> "a is {0}{2}{3}{1}".format(1,2,3,4)
'a is 1342'
last_name = "rachel"
first_name = "cook"
f"{last_name}{first_name}"
>> rechel cook
```

## list型

```python
>>> l = [1,2,3,4,5,6,7,8,9]
>>> l
[1, 2, 3, 4, 5, 6, 7, 8, 9]
>>> type(l)
<class 'list'>
>>> l[::3]　# 3つ飛ばし
[1, 4, 7]
>>> l[::-1]
[9, 8, 7, 6, 5, 4, 3, 2, 1]
>>> a = ["a","b","c"]
>>> b = [1,2,3,4,5]
>>> x = [a,b] # 二次元リスト(ネスト)
>>> x
[['a', 'b', 'c'], [1, 2, 3, 4, 5]]
>>> x[0][1]
'b'

>>> s = ["a","b","c","d","e","f"]
>>> s
['a', 'b', 'c', 'd', 'e', 'f']
>>> s[1:4] = ["B","C","D"]
>>> s
['a', 'B', 'C', 'D', 'e', 'f']
>>> s[1:4] = []
>>> s
['a', 'e', 'f']
>>> s[:] = []　# 全消去
>>> s
>>> n = [1,2,3,4,5,6,7,8,9]
>>> n.append(10)　# 後ろ
>>> n
[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
>>> n.insert(0,0)　# 追加位置を指定
>>> n
[0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
>>> n.pop()　# 取り出し
10
>>> n.pop(3)　# 場所指定
3
>>> del n[4]
>>> n
[0, 1, 2, 4, 6, 7, 8, 9]
>>> del n　# 全消去
>>> n = [1,2,2,2,3]
>>> n.remove(2)
>>> n.remove(2)
>>> n.remove(2)
>>> n
[1, 3]
>>> n.remove(2)　# エラーハンドリングに役立つ
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: list.remove(x): x not in list　
>>> a = [1,2,3,4,5]
>>> b = [6,7,8,9]
>>> x = a + b
>>> x
[1, 2, 3, 4, 5, 6, 7, 8, 9]
>>> a.append(b)　# リストごと入っちゃう(要素のみ入れたいときは+=, extend())
>>> a
[1, 2, 3, 4, 5, [6, 7, 8, 9]]
>>> a[5]
[6, 7, 8, 9]
>>> del a[5]
>>> a
[1, 2, 3, 4, 5]
>>> b
[6, 7, 8, 9]
>>> a += b
>>> a
[1, 2, 3, 4, 5, 6, 7, 8, 9]
>>> b
[6, 7, 8, 9]
>>> a.extend(b)
>>> a
[1, 2, 3, 4, 5, 6, 7, 8, 9, 6, 7, 8, 9]

r = [1,2,3,4,5,1,2,3]
print(r.index(3,3)) # 3番目から3を探す(2番目の3は無視)
print(r.index(3))
print(r.count(3))　# rの3の数を数える
if 5 in r:
    print('OK')
r.sort() # [1, 1, 2, 2, 3, 3, 4, 5]
print(r)
r.sort(reverse=True) # [5, 4, 3, 3, 2, 2, 1, 1]
print(r)
r.reverse() # [1, 1, 2, 2, 3, 3, 4, 5]
print(r)

s = "My name is Yui Kobayashi."
to_split = s.split(" ") # ['My', 'name', 'is', 'Yui', 'Kobayashi.']
print(to_split)
x = " 000000 ".join(to_split) # My 000000 name 000000 is 000000 Yui 000000 Kobayashi.
print(x)
print(help(list)) # listのmethodを確認できる

i = [1,2,3,4,5]
j = i
print(i)
print(j)

j[0] = 100
print(j) # [100,2,3,4,5]
print(i) # [100,2,3,4,5] (値渡しがおきて100がiにも入ってしまう→バグに繋がる)

x = [1,2,3,4,5]
y = x.copy() # 値渡しを避ける方法. y = x[:]でも出来るが分かりずらい
y[0] = 100
print(y) # [100,2,3,4,5]
print(x) # [1,2,3,4,5]
```

### list  copyでの注意点(値渡しと参照渡し)

```python
i = [1,2,3,4,5]
j = i
print(i)
print(j)
j[0] = 100
print(j) # [100,2,3,4,5]
print(i) # [100,2,3,4,5] 参照渡し(数値ではなくアドレスをコピー)がおきて100がiにも入ってしまう
# バグの原因になるので注意が必要

x = [1,2,3,4,5]
y = x.copy() # 参照渡しを避ける方法. y = x[:]でも出来るが分かりずらい
y[0] = 100
print(y) # [100,2,3,4,5]
print(x) # [1,2,3,4,5]

# 値渡し
x = 20
y = x　# adressがコピーされない
y = 5
print(x)
print(y)
print(id(x)) # 140726986859952 
print(id(y)) # 140726986859472 adressが違う

#参照渡し
X = ["a","b"]
Y = X # adressがコピーされる
Y[0] = "p"
print(X)
print(Y)
print(id(X)) # 1786278601416
print(id(Y)) # 1786278601416 adressが同じ
```

###　list内包記

```python
t = (1,2,3,4,5)
l =[]
for i in t:
    if i % 2 == 0:
        l.append(i)
    # else:
    #     l.append("#")
print(l)

# 1行で書ける
l = [i for i in t if i % 2 == 0]
print(l)

t1 = [1,2,3,4,5]
t2 = [6,7,8,9]
r = []
for i in t1:
    for j in t2:
        r.append(i * j)
print(r)
# これも1行で書けるがforが2回以上になると分かりづらいので非推奨
r = [i * j for i in t1 for j in t2]
print(r)
```



### list  使いどころ

> 例. 方舟の定員

```python
seat=[]
min = 0
max = 3
min <= len(seat) <= max #定員は３人まで

seat.append("D")
min <= len(seat) <= max #True
seat.append("K")
min <= len(seat) <= max #True
seat.append("I")
min <= len(seat) <= max #True
seat.append("G")
min <= len(seat) <= max #False
print(seat) #['D', 'K', 'I', 'G']

seat.pop(0) #"D"を消す
min <= len(seat) <= max #True
print(seat) #['K', 'I', 'G']
```

### listの中身を結合できる join

リストの中身を結合することが出来る便利なメソッド

```python
l = ['a', 'b', 'c', 'd']
join_l = ''.join(l)
print(join_l)
>> abcd
```

/ただし、int型が混ざるときは注意が必要である

```python
l = ['a', 'b', 'c', 2]
# map型に変換する必要がある
map_l = map(l)
join_l = ''.join(map_l)
print(join_l)
>> abc2
```

因みに分断したい時はsplitかlistを使う

```python
# 一文字ずつ区切る
list_l = list(join_l)
print(list_l)
>> ['a', 'b', 'c', '2']

# 特定のパーツを目印に区切る
split_l = join_l.split('b')
print(split_l)
>> ['a', 'c2']
```

## tuple型

> データの追加、変更が出来ない
>
> データの中身を変更されたくないときに使用

```python
t = (1,2,3,2,2,1,2,)　#()は省略可
print(t[0])
print(t.count(2))
print(t.index(2,2))

t = ([1,2,3],[2,3])
print(type(t))
t[1][0] = 8 # t[1] = 8はエラー
t = 1,2,3,4,5 # tuple
t = 1, # tuple(バグの原因になる)→t + 100が出来ない
new_tuple = (2,3,4)+(5,6,7) # tuple同士の足し算は出来る
#new_tuple = (1) + (2,3,4) # (1)はint型なので足し算が出来ない
```

###　tuple unpacking

``` python
num_tuple = (10, 20)
x, y = num_tuple # x, y = 10, 20 tupleの中身を展開する(変数宣言)
min, max = 0, 100 

a = 100
b = 200
a, b = b, a
print(a, b)  # tupleを使うと変数の中身の入れ替えが分かりやすくなる
```

###　tuple 使いどころ

> 例. ３つの選択肢から２つを選ぶアプリ

```python
choose_from_two = ("A", "B", "C")
answer = []
answer.append("A")
answer.append("C")
print(answer) # ["A", "C"]

choose_from_two = ["a", "b","c"]
answer =[]
choose_from_two.append("a") # 間違ってchoose_from_twoに入れてしまう
choose_from_two.append("c")
print(choose_from_two) # ['a', 'b', 'c', 'a', 'c']

# もしchoose_from_twoがtupleならappendの時点でエラーが出るのでバグだと分かる
```

## dictionary型

```python
d = {"x": 10, "y": 20}
d["x"] = 100　#{"x":100, "y":20}
dict(a=10, b=20)
dict([("a", 10), ("b", 20)])

d = {"X": 40, "Y": 200}
print(d.keys())
print(d.values())
d2 = {"X": 1000, "Z": 45}
d.update(d2)
print(d)
print(d.get("Z"))  # get()はない場合はNoneと返ってくる(d[]でない場合はエラー)
print(d.pop("X"))  # key "X"の値を取り出す
print(d) # get()と違い無くなっている
del d["Y"]
print(d)
del d # 存在が消える→中身だけ消したいときは d.clear を使う

d = {"a":100, "b":250}
"a" in d # True
"z" in d # False

# method
d = {"x": 100, "y": 200} 
print(d.items()) # dict_items([('x', 100), ('y', 200)])
for k, v in d.items():
    print(k, ":", v)
```

### dictionary copy

```python
a = {"x":20}
b = a
b["x"] = 100
print(b) # {'x': 100}
print(a) # {'x': 100}(参照渡し)

a ={"x":20}
b = a.copy() # 参照渡しを分ける記述
b["x"] = 100
print(b) # {'x': 100}
print(a) # {'x': 20}
```

### dictionary 使いどころ

> 商品と値段(listとの違い)

```python
fruits = {
    "grape":1000,
    "apple":250,
    "meron":3000
}
print(fruits["grape"]) # 章に飛んでページを探すイメージ(軽)

l = [
    ["a", 190],
    ["b", 320],
    ["c", 560]
]
print(l[2]) # listの場合は最初からページを捲って探すイメージ(重)
```

###　辞書内包記


> {キー: 値 for 任意の変数名 in イテラブルオブジェクト}


```python
w = ["drink1", "drink2", "drink3"]
f = ["tea", "milk", "coffee"]

d = {}
for x, y in zip(w, f):
    d[x] = y

# 1行で書ける
d = {x: y for x, y in zip(w, f)}
```



## 集合型

```python
>>> a = {2,3,4,4,5,6,6,6,7}
>>> a
{2, 3, 4, 5, 6, 7}　# 重複は1つ
>>> b = {1,2,3,5,5,8,9}
>>> b
{1, 2, 3, 5, 8, 9}　
>>> a - b
{4, 6, 7} # aからbにあるものを消す
>>> a & b # aかつb
{2, 3, 5}
>>> b - a
{8, 1, 9}
>>> a | b #aまたはb
{1, 2, 3, 4, 5, 6, 7, 8, 9}
>>> a ^ b #abの両方にないもの
{1, 4, 6, 7, 8, 9}

#集合内包記(リストと同じなのでリストを見て)
s = set()
for i in range(10):
    s.add(i)
# 1行
s = {i i in range(10)}
```

###　集合 method

```python
>>> s = {1,2,3,4,5} # 場所の概念がないので、s[0]はエラー
>>> s.add(6)
>>> s
{1, 2, 3, 4, 5, 6}
>>> s.remove(6)
>>> s
{1, 2, 3, 4, 5}
>>> s.clear()
>>> s 
set() # dictionary型と区別する手目
>>> a = {}
>>> type(a)
<class 'dict'>
>>> help(set)
```

###　集合 使いどころ

> 共通点探し、種類を知りたいとき

```python
my_friend = {"A", "B", "C"}
D_friend = {"B", "A", "E"}
print(my_friend&D_friend) # 共通の友達が分かる

f = {"apple", "banana","apple", "meron"}
print(set(f)) # {'banana', 'apple', 'meron'}
```

## coment

> comentは上に書くというPython上の暗黙のルール

```python
# applw value
"some_value" = 300
"""
複数行なら
こう書くのよね！！！
"""
```

## if

```python
i = 1
if i < 0:
    print("negative")
elif i == 1:
    print("congraguation")
elif i == 2:
    print("CONGRAGUATION")
else:
    print("positive")

a = 10
b = 20
if a > 0:
    if b > 0:
        print("OK") # 見づらい！
#上のスマート版
if a > 0 and b > 0: #論理演算子
    print("OK")
```

###　値が入っていない判定(かなり使える) 

```python
# False, 0, 0.0, {}, [], (), set()　True, それ以外
is_ok = []
# if len(is_ok) > 0:とやる必要がない
if is_ok:
    print("OK")
else:
    print("NO") # NO
```

###　Noneかどうかの判定

```python
is_empty = None
# print(help(is_empty)) # NoneType(何もない)

if is_empty is None: #変数に入っているのがNoneかどうかなら"is"を使う
    print("NO!!!")
print(1 == True)  # 0以外はTrueですよね
print(1 is True)  # 1とTrueはタイプが異なる
```



##　論理演算子

```python
a == b
a != b
a > 0 and b > 0
a > 0 or b > 0 #よりスマートに書けるようになる
```

###　in, not 使いどころ

> 　if文やコードを簡略化する際に用いる

```python
y = [1, 2, 3]
x = 1
if x in y:
    print("in")
if not 100 in y:
    print("not in")

is_ok = True
if is_ok:
    print("Hello")
# わざわざ"is_ok == True"と書く必要がない
if not is_ok:
    print("Goodnight")

a = 26
if not a == 25: #分かりづらい
    print("OK")
#次のコードの方が良い！
if a != 25:　# This is "simple".
    print("OK")
```

##　if, break, continue

```python
count = 0
print("start")
while count < 5:
    print(count)
    count += 1
print("finish")

count = 0
while True:
    if count >= 5:
        # 完全にループを抜ける
        break

    if count == 2:
        count += 1
        # 次のprint()へ行かずにループを続ける
        continue

    print(count)
    count += 1
print("finish")
```

##　for, break, continue

```python
some_list = [1, 2, 3, 4, 5]
# whileで書くと長い上に分かりづらい
i = 0
while i < len(some_list):
    print(some_list[i])
    i += 1
# forで書くとめちゃめちゃスッキリ
for i in some_list: # listの要素を順番にiに入れる
    print(i)
# 文字列も順番に入れる事が出来る
for i in "abdefg":
    print(i)
# continue, breakも同様に使える
for s in ["My", "name", "is", "Mike."]:
    if s == "is":
        continue
    if s == "is":
        break
    print(s)　# (OP) My, name, Mike.
```

## for else

```python
for fruits in ["banana", "apple", "orange"]:
    if fruits == "apple":
        print("stop eating")
        break
    print(fruits)
else:
    print("I ate all.")
```

##　関数

```python
# input関数
while True:
    word = input("Enter:")
    if word == "ok":
        break
    print("next")

#数を受け取りたいとき
while True:
    word = input("Enter:")
    # inputはstr型で受け取ることに注意
    num = int(word)
    if num == 100
        break
    print("next")

# range関数
for i in range(2,10,2): # 2~9で2個飛ばし
    print(i,"Hello")

# underscore:その数はforループの中では使われない事を示している
for _ in range(5):
    print("HELLO")

# enumerate関数
for i, fruits in enumerate(["banana", "apple", "orange"]):
    print(i, "fruit")
    
# zip関数
days = ["Mon","Tue", "Wed"]
drinks = ["coffee", "tea", "water"]
fruits = ["banana", "apple", "meron"]
for day, drink, fruit in zip(days, drinks, fruits):
    print(day, drink, fruit)
```

### 自作関数

```python
# 関数定義と実行(呼び出し方)
def say_something():
    print("hi")

f = say_something  # function
f() # "hi" と実行される

# 返り値(printを使う)
def say_hello():
    s = "Hello"
    return s

result = say_hello()
print(result)

# 引数
def what_is_this(color):
    if color == "red":
        return "tomato"
    elif color == "green":
        return "green pepper"
    elif color == "yellow":
        return "lemon"

result = what_is_this("red")
print(result)　# tomato

# 引数の型を宣言することは出来る
def add_num(a: int, b: int) -> int:
    return a + b

r = add_num(10, 20)
print(r)

# 注意点
def add_num(a: int, b: int) -> int:
    return a + b 

r = add_num("a", "b")
print(r) #errorで返ってこないので注意

# 位置引数、キーワード引数、デフォルト引数
def menu(entree, drink, dessert):
    print(entree, drink, dessert)

menu("beef", "tea", "ice")
# 指定してやれば引数に代入するものの順序が関係なくなる
menu(drink = "tea", entree = "beef", dessert = "ice")

def menu(entree = "beef", drink = "tea", dessert = "ice"):
    print("entree = ", entree)
    print("drink = ", drink)
    print("dessert = ", dessert)

menu() 
#書き換えも可能
menu(entree= "chicken")
```

### デフォルト引数でリストを渡すときの注意点

```python
# defalut引数にリストを渡すと参照渡しでバグに繋がる
def test_func(x, l=[]):
    l.append(x)
    return l

y = [1,2,3]
r = test_func(100, y)
print(r)

r = test_func(100) #[100]
r = test_func(100)#[100, 100]

#解決法
def test_func(x, l=None):
    if l is None:
        l = []
    l.append(x)
    return l

r = test_func(100) #[100]
r = test_func(100) #[100]

# *hogeで位置引数をtuple化
def say_something(word,*args):
    print("word =", word)
    # tuple化したものをforループで順番に出力してやる
    for arg in args:
        print(arg)

say_something("LOVE","hello","Mike","Nancy")

# advance
t =("Bob", "Yui")
# tupleをunpackingしてもう一度tuple化している
say_something("Hi!", *t)

# まとめて使うこともできる
def menu(food, *args, **kwargs):
    print(food)
    print(args)
    print(kwargs)
# tuple, dictでまとまる
menu("banana", "orange", "meron", entree="beef", drink="coffee")
>>
banana
('orange', 'meron')
{'entree': 'beef', 'drink': 'coffee'}
```

###　ローカル変数とグローバル変数

関数内で定義されたローカル変数とその外で定義されたグローバル変数がある

```python
num = 100
def Test():
    num = 1
    print(nun)
     
print(num) # 100
Test() # 1
print(num) # 100

# 関数内でgrobal変数を書き換えたい場合
def Test():
    global num
    num = 1
    print()
```



```python
# grobal変数
animal = "cat"
def f():
    # ローカル変数ならローカル変数のまま出力しようとするのでanimalがローカルで定義される前に出力しようとするとエラーが出る
    # print(animal)
    animal = "zebra"
    print("local:", locals())
f()
print("global:", animal)

def f():
    """####################"""
    print(f.__doc__) # "################"
    print(f.__name__) # "f"
f()
print("grobal:", __name__) # __main__
```



##　Docstrings

```python
# Docstrings
def example_func(param1,param2):
    """Example function with types documented in the docstrings.

    Args:
        param1 (int): The first parameter
        param2 (str): The second parameter
    
    Retures:
        bool: The return value. True for success, False otherwise.
    """
    print(param1)
    print(param2)
    return True

# 書いた説明書きが読める
example_func.__doc__ # help(example_func)でも可
```

##　closure

関数内関数(inner関数)をouter関数の返り値として実行する

### 用途１：クロージャー(用途１．好きなタイミングで実行出来る)

```python
def outer(a, b):
    def inner():
        return a + b
    # funciton自体を返り値として呼び出す
    return inner

# returnでinnerのfunctionを返しているのでまだ実行されていない
f = outer(1, 2)

print("#########") # 何かやりたい作業が間にあることも

# fを実行して初めて返り値を示す
r = f()
print(r)
```

### eg.  円の面積

```python
def circle_area_func(pi):
    def circle_area(radius):
        return pi * radius * radius
    # circle_areaというfuncionで返してやる
    return circle_area

# 引数を切り替えて実行する
ca1 = circle_area_func(3.14)
ca2 = circle_area_func(3.1415)

print(ca1(10))
print(ca2(10))
```

##　decorater

関数自体(関数の実行ではない)を引数にとってその関数処理の前後に別の処理や機能を付け加えることが出来る

一度記載すれば@のfunctionで何度も実行できるのがメリット

###　基本

```python
# 関数を引数にする
def outer(func):
    def inner(*args, **kwargs):
        print("FIRST")
        # ここで引数で受け取った関数を実行する
        result = func(*args, **kwargs)
        print("LAST")
        return result
    return inner

# add_num関数をdecorateします
def add_num(a, b):
    return a + b

decorated = outer(add_num)
decorated(23, 34)
print(decorated())
>>
FIRST
LAST
35
```



```python
# decorater1
def print_more(func):
    def wrapper(*args, **kwargs):
        print("func:", func.__name__)
        print("args:", args)
        print("kwargs:", kwargs)
        result = func(*args, **kwargs)
        print("result:", result)
        return result
    return wrapper

# decorater2
def print_info(func):
    def wrapper(*args, **kwargs):
        print("start")
        result = func(*args, **kwargs)
        print("end")
        return result
    return wrapper
```

### 呼び出し方

```python
@print_info
@print_more
def add_num(a, b):
    return a + b

r = add_num(1, 2)
```

出力結果

``` 
start
func: add_num
args: (1, 2)
kwargs: {}
result: 3
end
```

呼び出す順番で出力結果が変わることに注意

```python
@print_more
@print_info
def sub_num(c, d):
    return c - d

r = sub_num(4, 2)
print(r)
```

出力結果

```
func: wrapper
args: (4, 2)
kwargs: {}
start
end
result: 2
2
```

## lambda

あんまり使っているコードを見たところないけど、見てわかる程度には分かっておく。

``` python
l = ["Mon", "tue", "Wed", "Thu", "fri", "sat", "Sun"]
def change_words(words, func):
    for word in words:
        print(func(word))

def sample_func(word):
    return word.capitalize()

change_words(l, sample_func)
# lambdaを使うとシンプルに書ける
change_words(l, lambda word: word.capitalize())
change_words(l, lambda word: word.lower())
```

## ganerater

``` python
def greeting():
    yield "Good morning"
    #間に重い処理が入った際にも一回止まってくれる
    # for _ in range(10000):
    #     print("CHAGE")
    yield "Good afternoon"
    yield "Good night"

g = greeting()
# forループと違い、好きなタイミングで出力できる
print(next(g))
print("#############")
print(next(g))
print("#############")
print(next(g))


def counter(num= 10):
    for _ in range(num):
        yield "run"
    
c = counter()
print(next(c))
print(next(c))
print(next(c))
print(next(c))
print(next(c))
print("########")
print(next(c))
print(next(c))
print(next(c))
print(next(c))
print("########")
print(next(c))
# 11個目はerrorが返る(Error handling)
print(next(c))
```

### 内包記

``` python
def g():
    for i in range(10):
        yield i

g = g()
# 1行で書ける
g = (i for i in range(10)) # typeはgenerater
g = tuple(i for i in range(10)) #tuple
```

## 名前空間とスコープ

> local(def内)とglobalの使い分けが大事

``` python
animal = "cat"
def f():
    # ローカル変数ならローカル変数のまま出力しようとするのでanimalがローカルで定義される前に出力しようとするとエラーが出る
    # print(animal)
    animal = "zebra"
    print("local:", locals())
f()
print("global:", animal)

def f():
    """####################"""
    print(f.__doc__) # "################"
    print(f.__name__) # "f"
f()
print("grobal:", __name__) # __main__
```

##　例外処理

```python
# 例外処理
l = [1, 2, 3]
i = 5　

try:
    l[i]
# errorの内容が返せる
except IndexError as ex:
    print("Don't worry: {}".format(ex))
except NameError as ex:
    print(ex)
# pythonではその他全部を引っかけるような記述は好ましくないとされている
except Exception as ex:
    print("other: {}".format(ex))
# tryが成功した時だけ出力される
else:
    print("done")
# 最後に必ず実行される(その前にエラーが起きても)
finally:
    print("clean up")

print("####")
```

###　独自のエラー宣言

> 　独自のエラーを宣言して他のプログラマーに教えてあげる

``` python
# 独自のエラーを宣言
class UppercaseError(Exception):
    pass

def check():
    words = ["APPLE", "banana", "orange"]
    for word in words:
        if word.isupper():
            raise UppercaseError(word)

try:
    check()
except UppercaseError as ex:
    print("This is my falut. Go next.")
```

