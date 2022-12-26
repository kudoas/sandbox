# インスタンス化できないクラスのようなもの
# 名前空間として利用できる
module Brainfsck
  class Parser    
  end
end

module Whitespace
  class Parser
  end
end

Brainfsck::Parser
Whitespace::Parser


module Sweet
  def self.lot
    %w[brownie apple-pie bavarois pudding].sample
  end

  module Chocolate
  end
end

Sweet.lot
Sweet::Chocolate
