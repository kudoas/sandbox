# https://rurema.clear-code.com/3.2.0/method/Module/i/ancestors.html
'hellp'.class.ancestors

def hello(names)
  names.each do |name|
    puts "HELLO, #{name.upcase}"
  end
end

rubies = %w[MRI jruby mruby]

hello(rubies)
