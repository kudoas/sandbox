# 宇宙船演算子
p 1 <=> 2
p 2 <=> 1
p 1 <=> 1

["Alice", "Bob", "Charlie"].sort do |a, b|
  a.length <=> b.length
end
