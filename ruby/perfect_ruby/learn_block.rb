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

p File.open "with_block.txt", "w" do |file1|
  file1.puts "with block"
end

def display_value
  puts yield
end

# blockの返り値、最後に評価した式の値を返す
# => 1111
display_value do
  1111
end

# => 42
display_value do
  next 42
end

# display_valueを中断しているため何も返さない
display_value do
  break 42
end

# yieldの引数はブロックの引数として渡される
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

# 引数が多くてもエラーにはならないが、nilになる
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

# wrong number of arguments (given 1, expected 0) (ArgumentError)
# dummy_block_sample "hi"

people = ["Alice", "Bob", "Chalie"]
block = proc { |name| puts name } # Proc.new

people.each(&block)

p1 = proc { |val| val.upcase }
p2 = :upcase.to_proc

p1.call("h1")
p2.call("h1")

people.map { |person| person.upcase }
people.map(&:upcase)

def write_with_lock
  File.open "time.txt", "w" do |f|
    f.flock File::LOCK_EX

    yield f

    f.flock File::LOCK_UN
  end
end

write_with_lock do |f|
  f.puts Time.now
end

# method1 arg, method2 {...} blockはmethod2に渡される
# method1(arg, method2) {...} blockはmethod1に渡される
# method1 arg, method2 do
# ...
# end

# blockローカル変数
["Alice", "Bob", "Chalie"].each do |_person|
  someone = "block local"
end

other_people = []
["Alice", "Bob", "Chalie"].each do |person|
  other_people << person
end
# personは受け取れない！
