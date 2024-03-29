# coding: euc-jp

# https://rurema.clear-code.com/3.2.0/method/Module/i/ancestors.html
"hellp".class.ancestors

require "ostruct"

def hello(names)
  names.each do |name|
    puts "HELLO, #{name.upcase}"
  end
end

rubies = ["MRI", "jruby", "mruby"]

hello(rubies)

File.open "time.txt" do |file|
  puts file.read
end

line = "oneline"
 puts(line)

# global変数は$
$global_num = 777

# 定数は大文字
ENV = "production"
ENV = "staging"

ab = OpenStruct.new
ab.foo = 25
p ab
