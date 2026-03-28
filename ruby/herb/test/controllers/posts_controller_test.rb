require "test_helper"

class PostsControllerTest < ActionDispatch::IntegrationTest
  setup do
    @post = posts(:one)
  end

  test "should get index" do
    get posts_url
    assert_response :success
    assert_select "h1", "気軽に書いて、すぐ読めるブログ"
  end

  test "should get new" do
    get new_post_url
    assert_response :success
  end

  test "should create post" do
    assert_difference("Post.count") do
      post posts_url, params: {
        post: {
          body: "本文が20文字以上ある新しい記事です。ブログとして読める長さにしています。",
          title: "新しい記事"
        }
      }
    end

    assert_redirected_to post_url(Post.last)
  end

  test "should show post" do
    get post_url(@post)
    assert_response :success
    assert_select "h1", @post.title
  end

  test "should get edit" do
    get edit_post_url(@post)
    assert_response :success
  end

  test "should update post" do
    patch post_url(@post), params: {
      post: {
        body: "更新された本文です。最低文字数を超えるように十分な長さを入れています。",
        title: @post.title
      }
    }
    assert_redirected_to post_url(@post)
  end

  test "should destroy post" do
    assert_difference("Post.count", -1) do
      delete post_url(@post)
    end

    assert_redirected_to posts_url
  end

  test "should reject invalid post" do
    assert_no_difference("Post.count") do
      post posts_url, params: { post: { body: "short", title: "" } }
    end

    assert_response :unprocessable_entity
    assert_select "h2", "入力内容を確認してください"
  end
end
