Rails.application.routes.draw do
  # Angular で設定しているルーティングと一致させればAngular Router のURLを保持できる
  # 設定しない場合 (eg. '*path') とかだと、/ にリダイレクトされる
  get '/hero', to: 'angular#index'
end
