# catch/throw でループを任意の場所から抜け出せる
catch :outer do
  catch :middle do
    catch :inner do
      throw :middle
      p "not"
    end
    p "not"
  end
  p "go!!!"
end

# 第二引数が返り値
p catch(:foo) {
  throw :foo, "retruned value"
}

# ぼっち演算子
people = ["alice", "bob", "carol"]
people[3]&.capitalize

# メソッド名が大文字の時は括弧を省略できない！
def Hello
  p "Hello"
end

Hello()

def greet(recipient)
  return false unless recipient

  "Hi, #{recipient.capitalize}"
end

p greet("alice")
p greet(nil)
