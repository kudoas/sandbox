class Ruler
  def initialize(length)
    @length = length
  end

  def to_s
    "=" * @length
  end

  def inspect
    "Ruler(length: #{@length})"
  end
end

ruler = Ruler.new(30)
p ruler.to_s

# a = Integer(gets)
# b = Integer(gets)

num = 10
`head -#{num} ~/.vimrc`

system("uname")
pid = spawn("uname")

# SIGINT: Ctrl + C
# trap :INT do
#   puts "\nIntercept!"
#   exit
# end

trap :EXIT do
  puts "Finalizing..."
end

puts "Running..."
