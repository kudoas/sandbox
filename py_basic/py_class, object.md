#　classとobject指向

grobalは汚さない方がよい！汚さないで記述する方法がobject指向である。

## classとは

ずっと使ってきたmethodは色々なclassの中に定義されている. それをobjectとして引き出して使う。

object指向では非常に重要である(後に後述)

```python
f = "djkhfaoilesjflajsofd"
# capitalize()は下記のクラスの中に定義されている
print(f.capitalize())

"""
# classを一部抜粋
class str(object):
    def captalize(self)
"""
```

##　classとobject指向

親クラスとしてobjectを渡してある(諸説あり)

```python
class Person(object):
    # 引数のselfはクラスに値を保持させるために必要
    def say_something(self, name):
        self.name = name
        if self.name == "girl":
            print("Hello")
            
"""
直接関数を宣言するよりもクラスを定義して
「人間が何か言った」という形で出力した方が現実社会に即している
(例は女の子なら挨拶する男のしたたかさを表している)
-> オブジェクト指向
"""     
person = Person()
person.say_something("girl")
>> Hello
```

## 特殊メソッド

特殊な条件で呼び出されるメソッド、いっぱいあるからドキュメント参照

###　\__init__()

クラスが呼び出されたときにで最初の呼び出されるmethod. 

ここで初期設定(ずっと使う変数の設定、フォントサイズや色etc)を行えばよい。

```python
class Person(object):
    def __init__(self):
        print("Hello")
    
    def say_something(self, name):
        self.name = name
        if self.name == "girl":
            print("fXXk you!!!")
# 呼び出されるだけで発動する
person = Person()
>> Hello
```

###　\__str__()

オブジェクトを文字列扱いすると呼び出せる

```python
class Sample(object):
    def __str__():
        return "WORD!!!"

sample = Sample()
# 文字列扱い
print(sample)
>>> WORD!!!
```



##　basic (self, constructor, destructor)

``` python
class Person(object):
    
    def __init__(self, name):
        # Personのobjectにnameの値を保持させたい時にselfを使う
        self.name = name

    def say_something(self):
        # selfでnameを保持しているから他のmethodを使う時でも呼び出せる
        # self.hogehogeの形
        print("I am {}. hello".format(self.name))
        # このmethodから別のmethodを呼び出すこともできる
        self.run(10)
	
    def run(self, num):
        print("run " * num)
    # 最後に読まれるmethod(デストラクタ)
    def __del__(self):
        print("goodbye")

# methodを呼び出すとき
# classを呼び出す
person = Person("Daichi")
# 中の関数を呼び出す
person.say_something()

# ここでmethodを終了するとこの時点でgoodbyeが表示される
# del person

print("##########")

>> # 出力
I am Daichi. hello
run run run run run run run run run run
##########
goodbye

# コードは下のようにも書けるが、classで定義した方が人間に分かりやすい
"""
def person(name):
    if name == "A":
        print("hello")
"""
```

## classの継承

コードを綺麗にできる, Object指向で重要

```python
class Car(object):
    def run(self):
        print("run")
        
# Car()のmethodを引き継げるし、()の中を見れば由来が分かるのでコードがきれいに書ける
class ToyotaCar(Car):
    # passは特に何もしないので、引き継いだCarが発動
    pass

# Car()を引き継ぎ、かつauto_run()もあるので両方のmethodを呼び出せる
class TeslaCar(Car):
    def auto_tun(self):
        print("auto run")

car = Car()
car.run()
>> run

toyota_car = ToyotaCar()
# 引き継いでいるのでCarが発動
toyota_car.run()
>> run

tesla_car = TeslaCar()
tesla_car.run()
>> run

tesla_car.auto_tun()
>> auto run
```

### 多重継承

複数のクラスを継承できる(ただし順番に注意)

```python
class Person(object):
    def talk(self):
        print("talk")

    def run(self):
        print("person run")


class Car(object):
    def run(self):
        print("car run")

# run関数がCar()とPerson()にあるので後に来た方に上書きされる
class PersonCarRobot(Person, Car):
    def fly(self):
        print("fly")

person_car_robot = PersonCarRobot()
person_car_robot.talk()
person_car_robot.run()
person_car_robot.fly()

>> talk
>> car run
>> fly
```

## methodのオーバーライド

```python
# 継承しても再度定義し直せば上書きできる
class Car(object):
    def run(self):
        print("run")
        
class ToyotaCar(Car):
    def run(self):
        print("fast")
       
toyotacar = ToyotaCar()
toyotacar.run()
>> fast
```

## classから呼び出せるのはmethodだけではない

```python
class Car(object):
    def __init__(self, model=None):
        self.model = model
        
toyota_car = ToyotaCar("Lexus")
# methodだけじゃなくてクラス変数も呼び出せるのでメモ
print(toyota_car.model)
>> Lexus
```

##　superによる親methodの呼び出し

継承したクラスで親クラスと同じ関数を定義する場合

```python
class Car(object):
    def __init__(self, model=None):
        self.model = model
    
class TeslaCar(Car):
    # 下は親クラスにもあるので上書きされてしまう(書き直す必要がある)
    def __init__(self, model="Model S", enable_auto_run=False):
        self.model = model
        self.enable_auto_run = enable_auto_run
```

親クラスの処理が増えた時に2回書くのは面倒

```python
class Car(object):
    def __init__(self, model=None):
        self.model = model
    
class TeslaCar(Car):
    def __init__(self, model="Model S", enable_auto_run=False):
        # super()は親クラスという意味。親クラスがもつメソッドを呼び出し、更に新たに変数を定義できる。
        super().__init__(model)
        self.enable_auto_run = enable_auto_run
```

### 例.  将棋

```python
class Shogi(object):
    def __init__(self, model="9*9", player="HABU"):
        self.model = model
        print("Shogi __init__:", model, player)


class Chess(Shogi):
    def __init__(self, model="8*8", player="GAURI"):
        # super().__init__(model, player)にするとShogiから呼び出されるもののmodel, playerが上書きできる
        super().__init__()
        self.player = player
        print("Chess __init__:", model, player)


chess = Chess()
>>>
Shogi __init__: 9*9 HABU
Chess __init__: 8*8 GAURI
```



##　propertyを使った属性の設定

```python
class TeslaCar(Car):
    def __init__(self, model="Model S",
                 enable_auto_run=False,
                 passwd="1234"
                 ):
        self.model = model
        super().__init__(model)
        # __enable_auto_runだとclass内以外からは書き換えられない
        self._enable_auto_run = enable_auto_run
        self.passwd = passwd

    @property
    def enable_auto_run(self):
        return self._enable_auto_run

    @enable_auto_run.setter
    def enable_auto_run(self, is_enable):
        if self.passwd == 1234:
            self._enable_auto_run = is_enable
        else:
            raise ValueError

# passwardが正しければ書き換えられるプログラム
# 間違えているとValueError
tesla_car = TeslaCar("Model S", passwd=1234)
tesla_car.enable_auto_run = True
print(tesla_car.enable_auto_run)

>> True
```

## classをデータ構造体として扱う際の注意事項

classはデータ構造体として扱えるのは便利(あとから変数を入れられる)

```python
class T(object):
    pass
t = T()
# T()に変数が入っていく
t.name = "Mike"
t.age = "21"
print(t.name, t.age)
>> Mike 21
```

classの中に元々定義されたものがあると、書き換えた際に書き換えられてしまう

```python
class TeslaCar(Car):
    def __init__(self, model="Model S",
                 enable_auto_run=False,
                 passwd="1234"
                 ):
        self.model = model
        super().__init__(model)
        self._enable_auto_run = enable_auto_run
        self.passwd = passwd

    @property
    def enable_auto_run(self):
        return self._enable_auto_run

    @enable_auto_run.setter
    def enable_auto_run(self, is_enable):
        if self.passwd == 1234:
            self._enable_auto_run = is_enable
        else:
            raise ValueError

tesla_car = TeslaCar("Model S")
# 書き換えると再定義されて出力できてしまうので注意
tesla_car.__enable_auto_run = "##########"
print(tesla_car.__enable_auto_run)
>>> ##########
```

##　クラスメソッドとスタティックメソッド

```python
class Person(object):

    kind = "human"

    def __init__(self):
        self.x = 100
# クラスのオブジェクトを指している
a = Person()
print(a)
>>>　<__main__.Person object at 0x00000272F8712240>

# クラス自身を指している
b = Person
print(b)
>>>　<class '__main__.Person'>

# クラス内のオブジェクトを呼び出す
print(a.x) # 100
print(b.x) # オブジェクトではないのでerror
 
# しかしどちらもkindにはアクセスできる(オブジェクト内のグローバル変数的なイメージ)
print(a.kind)
print(b.kind)
>>>
human
human
```

### オブジェクトを生成する前にアクセスする方法

###　①クラスメソッド

```python
class Person(object):

    kind = "human"

    def __init__(self):
        self.x = 100
    
    # classのまま呼び出せる
    @classmethod
    def what_is_your_kind(cls):
        return cls.kind
    
b = Person
print(b.what_is_your_kind())
>>> human
# これはあんまり関係ないけどオブジェクト内のグローバル変数的なイメージでOK
print(Person.kind)
>>> human
```

###　②スタティックメソッド

使用頻度低め。外に定義しても使えるが、クラスとの関連性は担保しておきたい時に使用する。

```python
class Person(object):
    
    kind = "human"
    
    # オブジェクトを作らなくてもダイレクトにクラスから呼び出せる
    # selfもいらない
    @staticmethod
    def about(year):
        print('human about {}'.format(year))
        
Person.about(1999)
>>> human about 1999
```

