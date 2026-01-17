require "minitest/autorun"
require "./fizzbuzz"

describe "数値と文字列を扱うFizzBuzzクラス" do
  before { @fizzbuzz = FizzBuzz.new }

  describe "3の倍数のときは文字列の代わりにFizzを返す" do
    it "3を渡すと文字列Fizzを返す" do
      assert_equal "Fizz", @fizzbuzz.convert(3)
    end
  end

  describe "5の倍数のときは文字列の代わりにBuzzを返す" do
    it "5を渡すと文字列Buzzを返す" do
      assert_equal "Buzz", @fizzbuzz.convert(5)
    end
  end

  describe "その他の数値のときはその数値の文字列を返す" do
    it "1を渡すと文字列1を返す" do
      assert_equal "1", @fizzbuzz.convert(1)
    end

    it "2を渡すと文字列2を返す" do
      assert_equal "2", @fizzbuzz.convert(2)
    end
  end
end
