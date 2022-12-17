class MyObject
  def ==(other)
    self.class == other.class
  end
end

MyObject.new == MyObject.new

# 同じオブジェクトかはequal?を使う
a = [1, 2, 3]
b = [1, 2, 3]
a.equal?(b) # => false

n = 0
p "0!!!" if n.zero?

stone = "ruby" 

# caseは===で比較するのでregexpでも書ける
case stone
when "ruby" # stone == "ruby"
  puts "7"
when "perl", "white"
  puts "6"
else
  "unknown"
end

languages = ["Perl", "Python", "Ruby"]
i = 0
while i < languages.length # untilは逆
  p languages[i]
  i += 1
end

# sleep 1 while processing?
# begin
#   process1
#   process2
# end while needed?

# while, until, forはブロックにならない
for name in ["Alice", "Bob", "Carol"]
  p name
end
p name # => Carol


# 直接実行された時だけ実行する
return unless $PROGRAM_NAME == __FILE__

puts "このファイルは直接実行されたぜ！"

__END__

ここには何を書いていても問題ないぜ〜！
俺がついてるぜ〜！
