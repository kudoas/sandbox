class Person
  attr_reader :name
  attr_reader :contacts

  def initialize(name:)
    @name = name
    @contacts = []
  end

  def speak
    "I'm #{@name} and I love Ruby!"
  end
end

person = Person.new(name: "John")
person.speak
