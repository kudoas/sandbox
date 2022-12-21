module Sweet
  def self.lot
    %w[brownie apple-pie bavarois pudding].sample
  end

  module Chocolate
  end
end

Sweet.lot
Sweet::Chocolate
