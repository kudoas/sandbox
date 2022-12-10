# raise "error!" => RuntimeError

begin
  1 / 0
rescue StandardError => e
  p e.class
  p e.message
  p e.backtrace
ensure # 必ず実行する
  p "ERROR!!"
end

# 戻り値はbegin/rescueの最後の式の評価になる
returned = 
  begin
    value = "return value"
    raise
  rescue
    value
  ensure
    "this is not return value"
  end

puts returned

# retry
begin
  failed ||= 0
  p "try"

  1 / 0
rescue
  failed += 1

  retry if failed < 5
end

# Exception#cause: 例外処理中に別の例外処理が発生したときに情報を取得srうう
begin
  begin
    raise "original error!"
  rescue => e1
    raise "another error!"
  end
rescue => e2
  p e2
  p e2.cause
end
