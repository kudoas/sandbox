class Post < ApplicationRecord
  validates :title, presence: true, length: { maximum: 80 }
  validates :body, presence: true, length: { minimum: 20 }

  scope :recent_first, -> { order(created_at: :desc) }

  def excerpt
    body.to_s.truncate(140)
  end
end
