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
