class FizzBuzz
  def convert(num)
    return "Fizz" if num % 3 == 0
    return "Buzz" if num % 5 == 0
    return num.to_s
  end
end
