# 宇宙船演算子
p 1 <=> 2
p 2 <=> 1
p 1 <=> 1

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
