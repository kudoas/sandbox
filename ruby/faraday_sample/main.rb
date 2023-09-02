require 'csv'
require 'faraday'

class User
  def initialize(name, age)
    @name = name
    @age = age
  end
end

User.new('test', 10)
c = CSV.new
client = Faraday.new

arr = []
arr.flatten
