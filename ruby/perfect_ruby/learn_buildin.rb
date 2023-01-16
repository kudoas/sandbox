# 宇宙船演算子
p 1 <=> 2
p 2 <=> 1
# p 1 <=> 1

["Alice", "Bob", "Charlie"].sort do |a, b|
  a.length <=> b.length
end

3.step 10 do |num|
  puts num
end

2.next
2.succ

(65..70).map { |n| p n.chr }
Integer("2312")
# Integer(nil) => TypeError

1.upto 3 do |num|
  puts num
end

arr = 20_230_102.digits
p arr
p 7.digits(2)

c = Complex(2, 3)
c.real
c.imaginary
222.to_c # => (222+0i)

"".empty?
"てにをは".bytesize
p "Alice Bod".include?("Bod")
p "Highlight".start_with?("High")

# 文字列はミュータブル
str = "Pine"
str << "apple"
str # => "Pineapple"

message = "The Answer to life, the unicerse, and everything: 42"
message.slice(4)
message.slice(4, 6)
message.slice(4..9)
message[4]

# \t\r\n\f\vと空白
spliter = " hi \t"
spliter.strip
spliter.rstrip

"hi \n\n".chomp # => hi \n
"Users".chop # => User
"wooooooooooooooo".squeeze # => wo
"aabbbcccdddd".squeeze("abc") # => abcdddd

"Abc".swapcase # => aBC
"title".capitalize # => Title

spliter.strip! # breaking change
spliter.strip! # => nil

"abcde".reverse # => "edcba"

arr_str = "Alice, Bob, Chalie"
arr_str.split(", ")
arr_str.split(/,\s+/)
arr_str.split(/,\s+/, 2) # => ["Alice", "Bob, Chalie"]

"bbb".each_char { |s| p "#{s}a" }
"Alice\nBob".each_line { |line| p line }

"てにをは".encoding # => #<Encoding:UTF-8>
"てにをは".encode(Encoding::EUC_JP).encoding # => #<Encoding:EUC-JP>

# => incompatible character encodings: UTF-8 and EUC-JP (Encoding::CompatibilityError)
# "てにをは" + "てにをは".encode(Encoding::EUC_JP)

# ASCII互換だと連結や比較ができる

# 末尾の1バイトが削れてしまった文字列
broken = "尾崎翠".bytes[0..-2].pack("c*").force_encoding(Encoding::UTF_8)
p broken
p broken.scrub
p broken.scrub("*")

limted_str = String.new(capacity: 10_000)
2_000.times do
  limted_str << "hello"
end

# /[0-9]/ === "ruby"
# /[0-9]/ === "ruby4" # => true

/[0-9]/ =~ "ruby"
/[0-9]/ =~ "ruby4"

str = "ruby5"
matched = /[0-9]/.match(str)
p matched[0]
p matched.pre_match
p matched.captures

Regexp.last_match

/ruby/.match("ruby")
Regexp.last_match # => ruby

/\d/.match?("ruby5")
"ruby5".match?(/\d/)
:ruby5.match?(/\d/)
Regexp.last_match # => ruby5

names = "Yamazaki Tanizaki"
names.scan(/\w+zaki/) { |s| puts s.upcase }
names.scan(/(\w+)zaki/) { |s| puts s[0].upcase }

part = Regexp.escape("(incomplete)") # => "\\(incomplete\\)"

# 文字クラス
# \w = [0-9A-Za-z]
# \d = [0-9]

# 量指定子
# + 直前パターン1回以上の繰り返し

# xxx-xxx-xxxx
phone_number_pattern = /\A\d{3}-\d{4}-\d{4}\z/
# phone_number_pattern === "080-2222-3333" # => true
phone_number_pattern.match?("080-2222-3333")
phone_number_pattern.match?("phone: 080-2222-3333 (mobile)") # => false

lines = "1234\nabcd"
/\A\d+\z/.match?(lines) # => false 文字列全体なのでマッチしない
/^\d+$/.match?(lines) # => true 行頭と行末がマッチする
