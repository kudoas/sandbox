# https://docs.ruby-lang.org/ja/latest/doc/spec=2fcall.html#super
class Foo
  def foo(arg = nil)
    p arg
  end
end

class Bar < Foo
  def foo(arg)
    super(5)       # 5 を引数にして呼び出す
    super(arg)     # 5 を引数にして呼び出す
    super          # 5 を引数にして呼び出す super(arg) の略記法
    arg = 1
    super          # 1 を引数にして呼び出す super(arg) の略記法
    super()        # 引数なしで呼び出す
  end
end

Bar.new.foo 5

class Calculator
  def twice(n = nil)
    n * 2
  end
end

class Moblie < Calculator
  def twice
    n = super(10) # calculator.twice(10)の返り血を返す
    n * 1.5
  end
end

p Moblie.new.twice
