# Client への Interface
class Printer
  def initialize(obj)
    @obj = obj
  end

  def print_weak
    @obj.print_weak
  end

  def print_strong
    @obj.print_strong
  end
end

# Targetにはないインターフェイスを持つ (Adaptee)
class OldPrinter
  def initialize(string)
    @string = string.dup
  end

  # カッコに囲って文字列を表示する
  def show_with_paren
    puts "(#{@string})"
  end

  # アスタリスクで囲って文字列を表示する
  def show_with_aster
    puts "*#{@string}*"
  end
end
# Targetが利用できるインターフェイスに変換 (Adapter)
class Adapter
  def initialize(string)
    @old_printer = OldPrinter.new(string)
  end

  def print_weak
    @old_printer.show_with_paren
  end

  def print_strong
    @old_printer.show_with_aster
  end
end

# 利用者(Client)
p = Printer.new(Adapter.new("Hello"))
p.print_weak
#=> (Hello)

p.print_strong
#=> *Hello*
