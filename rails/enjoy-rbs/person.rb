# rbs_inline: enabled

class Person
  attr_reader :name #: String

  attr_reader :addresses #: Array[String]

  # You can write the type of parameters and return types.
  #
  # @rbs name: String
  # @rbs addresses: Array[String]
  # @rbs return: void
  def initialize(name:, addresses:)
    @name = name
    @addresses = addresses
  end

  # Or write the type of the method just after `@rbs` keyword.
  #
  # @rbs () -> String
  def to_s
    "Person(name = #{name}, addresses = #{addresses.join(", ")})"
  end

  # @rbs () -> Integer
  def hash
    [name, addresses].hash
  end

  # @rbs &block: (String) -> void
  def each_address(&block) #: void
    addresses.each(&block)
  end
end

person = Person.new(name: "Alice", addresses: ["123 Main St", "456 Elm St"])
person.each_address do |address|
  puts "Address: #{address}"
end
