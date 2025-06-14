class JobsController < ApplicationController
  def index
    @job_results = read_job_results
  end

  def create
    name = params[:name].presence || "Anonymous"
    message = params[:message].presence || "Hello from Sidekiq!"

    # Sidekiqを直接使用してジョブをキューに追加
    SampleJob.perform_async(name, message)

    flash[:notice] = "🚀 ジョブがキューに追加されました！ 5秒後に完了予定です。"
    redirect_to jobs_path
  end

  def send_email
    email = params[:email].presence || "test@example.com"
    subject = params[:subject].presence || "Test Email from Sidekiq"
    body = params[:body].presence || "This is a test email sent via Sidekiq!"

    # 高優先度キューでメールジョブを実行
    EmailJob.perform_async(email, subject, body)

    flash[:notice] = "📧 メールジョブがキューに追加されました！ 3秒後に完了予定です。"
    redirect_to jobs_path
  end

  def clear_results
    result_file = Rails.root.join("tmp", "job_results.txt")
    File.delete(result_file) if File.exist?(result_file)

    # Turbo Streamsで結果リストを更新
    respond_to do |format|
      format.html do
        flash[:notice] = "🗑️ 結果をクリアしました。"
        redirect_to jobs_path
      end
      format.turbo_stream do
        # 結果リストを空の状態で更新
        render turbo_stream: turbo_stream.replace(
          "job_results_content",
          partial: "results_content",
          locals: { job_results: [] }
        )
      end
    end
  end

  private

  def read_job_results
    result_file = Rails.root.join("tmp", "job_results.txt")
    return [] unless File.exist?(result_file)

    File.readlines(result_file).reverse.first(10)
  end
end
