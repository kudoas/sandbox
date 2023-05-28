Rails.application.routes.draw do
  get '/client/*other', to: 'angular#index'
end
