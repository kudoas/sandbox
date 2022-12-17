# yield: 受け取ったブロックを実行する
def block_sample
  puts "stand up"
  yield if block_given?
  puts "sit down"
end

block_sample do
  puts "walk"
end

block_sample

file = File.open("without_block.txt", "w")
file.puts "without block"
file.close

File.open "with_block.txt", "w" do |file1|
  file1.puts "with block"
end

def display_value
  puts yield
end

display_value do
  1111
end # => 1111

display_value do
  next 42
end

display_value do
  break 42 # display_valueを中断しているため何も返さない
end

def with_current_time
  yield Time.now
end

with_current_time do |now|
  puts now.year
end

# 引数がない場合は何も起こらない
with_current_time do
  puts "Hi"
end

# 引数が多くてもエラーにはならないがnil
with_current_time do |now, something|
  puts now
  puts something # nil
end

def default_argument_for_block
  yield
end

# default value
default_argument_for_block do |val = "Hi"|
  puts val
end

def flexible_argument_for_block
  yield 1, 2, 3
end

flexible_argument_for_block do |*params|
  puts params.inspect
end

def dummy_block_sample(&block)
  puts "stand up"

  block.call if block

  puts "sit down"
end

dummy_block_sample do
  puts "walk"
end
