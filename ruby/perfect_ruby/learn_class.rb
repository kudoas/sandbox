class MyClass
  def hello
    puts "Hello, My Object"
  end
end

my_object = MyClass.new
my_object.hello

class Ruler
  # 代入と取得両方使える
  attr_accessor :length

  # 使用用途が限定的ならこっちを使う
  # attr_reader
  # attr_writer

  # 代入用メソッド
  # def length=(val)
  #   @length = val
  # end

  # 取得用メソッド
  # def length
  #   @length
  # end

  def initialize(length = nil)
    @length = length
  end

  def display_length
    puts length # 参照は省略できる
  end

  def set_length
    self.length = 30 # 代入は省略できない
  end

  # class method
  class << self
    def pair
      [Ruler.new(250), Ruler.new]
    end

    def trio
      [Ruler.new, Ruler.new, Ruler.new]
    end
  end
end

ruler = Ruler.new
puts ruler.length

ruler.length = 10 # length= が呼ばれる
puts ruler.length
ruler.display_length

ruler.set_length
puts ruler.length

my_rule = Ruler.new(100)
puts ruler.length

p Ruler.pair

class MyClass
  # クラス変数：クラスとインスタンスがスコープ
  @@cvar = "Hello, My class variable"

  def cvar_in_method
    p @@cvar
  end

  def self.cvar_in_class_method
    p @@cvar
  end
end

my_object = MyClass.new

my_object.cvar_in_method
MyClass.cvar_in_class_method

class Parent
  def hello
    puts "Hello, Parent class!"
  end
end

class Child < Parent
  def hi
    puts "Hello, Child class!"
  end

  def hello
    super # お案じ名前のメソッドを呼び出せる
    puts "Hello, Child class!"
  end
end

child = Child.new
child.hi
child.hello
Child.superclass # Parent
