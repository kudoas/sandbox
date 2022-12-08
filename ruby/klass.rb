class MyClass
  def hello
    puts 'Hello, My Object'
  end
end

my_object = MyClass.new
my_object.hello


class Ruler
  attr_accessor :length # length=とlength

  # インスタンス変数の代入するときに使うメソッド
  # def length=(val)
  #   @length = val
  # end

  # 取得用のメソッド
  # def length
  #   @length
  # end
end

ruler = Ruler.new
ruler.length = 10 # length= が呼ばれる
puts ruler.length
