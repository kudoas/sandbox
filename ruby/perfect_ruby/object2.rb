# str
str = <<-EOS
  one
  two
EOS
puts str

# %qでスケープする
puts %q(It's #{weather})

# symbol: 読み書きのしやすさ => 識別子/キーワード
:ruby

people = %w[Mario Peach Teresa]
p people[0]

%i(red, green, blue)

# hash
colors = { red: "ff0000", green: "00ff00" } # => { :red => "ff0000", ...}
puts colors[:red]

# range
(1..5).include?(5) # true
(1...5).include?(5) # false

("a".."c").each { |s| puts s }

pattern = /[0-9]+/
p pattern === "HAL 9000"
p pattern =~ "HAL 9000" # 4

%r(/usr/bin) # regexp使うときに中で\が不要になって楽
%r</usr/bin> # ()じゃなくてもOK

# Proc
greeter = Proc.new { |name|
  puts "Hello, #{name}"
}

greeter.call "Proc"
by_literal = ->(name) { puts "Hello, #{name}" } 

# 代入
a, *b = [1, 2, 3]
c ||= 2

# exception
begin
  1 / 0
rescue => exception
  puts exception
end

# ruby object2.rb foo
p ARGV # => foo
