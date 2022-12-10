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
