require "test_helper"

class PostTest < ActiveSupport::TestCase
  test "is invalid without a title" do
    post = Post.new(body: "十分な長さの本文です。ブログ記事として成立する長さにしています。")

    assert_not post.valid?
    assert_includes post.errors[:title], "can't be blank"
  end

  test "is invalid with a short body" do
    post = Post.new(title: "タイトル", body: "short")

    assert_not post.valid?
    assert_includes post.errors[:body], "is too short (minimum is 20 characters)"
  end

  test "orders recent posts first" do
    older = Post.create!(title: "古い記事", body: "これは十分な長さの本文です。古い記事として作成します。")
    newer = Post.create!(title: "新しい記事", body: "これは十分な長さの本文です。新しい記事として作成します。")

    assert_equal [newer, older], Post.recent_first.limit(2)
  end
end
