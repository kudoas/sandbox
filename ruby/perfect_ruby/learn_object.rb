# 全てのオブジェクトはObjectのサブクラス
class MyClass
end

MyClass.superclass

o = Object.new
o.class
o.is_a?(Object)
o.object_id
o.nil?
o.frozen?
o.tap { |v| puts v }

class Ruler
  attr_accessor :length

  def initialize(length)
    self.length = length
  end

  def ==(other)
    length == other.length
  end
end

r1 = Ruler.new(10)
r2 = Ruler.new(10)
p r1 == r2
r2.length = 20
p r1 == r2

o.freeze

# objectのコピー
original = Object.new
original.freeze

copy_dup = original.dup
copy_dup.frozen? # false

copy_clone = original.clone
copy_clone.frozen? # true

# Object#dupやObject#cloneはshallow copyなので、配列の要素まではコピーされない
value = "foo"
array = [value]
array_dup = array.dup
array_clone = array.clone

value.object_id

p array.object_id
p array_dup[0].object_id
p array_clone[0].object_id
