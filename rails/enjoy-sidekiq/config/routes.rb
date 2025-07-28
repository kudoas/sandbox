Rails.application.routes.draw do
  # Define your application routes per the DSL in https://guides.rubyonrails.org/routing.html

  # Reveal health status on /up that returns 200 if the app boots with no exceptions, otherwise 500.
  # Can be used by load balancers and uptime monitors to verify that the app is live.
  get "up" => "rails/health#show", as: :rails_health_check

  # Render dynamic PWA files from app/views/pwa/* (remember to link manifest in application.html.erb)
  # get "manifest" => "rails/pwa#manifest", as: :pwa_manifest
  # get "service-worker" => "rails/pwa#service_worker", as: :pwa_service_worker

  # Action Cable endpoint
  mount ActionCable.server => "/cable"

  # Sidekiq Web UI
  require "sidekiq/web"
  mount Sidekiq::Web => "/sidekiq"

  # Sidekiq routes
  resources :jobs, only: [ :index, :create ] do
    collection do
      delete :clear_results
      post :send_email
    end
  end

  # Defines the root path route ("/")
  root "jobs#index"
end
