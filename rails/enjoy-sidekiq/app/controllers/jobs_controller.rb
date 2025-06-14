class JobsController < ApplicationController
  def index
    @job_results = read_job_results
  end

  def create
    name = params[:name].presence || "Anonymous"
    message = params[:message].presence || "Hello from Sidekiq!"

    # Sidekiqを直接使用してジョブをキューに追加
    job_id = SampleJob.perform_async(name, message)

    # プログレスバーを即座に表示
    show_initial_progress("SampleJob", job_id)

    flash[:notice] = "🚀 ジョブがキューに追加されました！進行状況をリアルタイムで確認できます。"
    redirect_to jobs_path
  end

  def send_email
    email = params[:email].presence || "test@example.com"
    subject = params[:subject].presence || "Test Email from Sidekiq"
    body = params[:body].presence || "This is a test email sent via Sidekiq!"

    # 高優先度キューでメールジョブを実行
    job_id = EmailJob.perform_async(email, subject, body)

    # プログレスバーを即座に表示
    show_initial_progress("EmailJob", job_id)

    flash[:notice] = "📧 メールジョブがキューに追加されました！進行状況をリアルタイムで確認できます。"
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

  def show_initial_progress(job_type, job_id)
    # 初期プログレスバーを表示
    Turbo::StreamsChannel.broadcast_prepend_to(
      "job_notifications",
      target: "progress_area",
      partial: "shared/progress_bar",
      locals: {
        job_id: job_id,
        job_type: job_type,
        percentage: 0,
        current_step: 0,
        total_steps: job_type == "SampleJob" ? 5 : 4,
        message: "ジョブをキューに追加しました..."
      }
    )
  end

  def read_job_results
    result_file = Rails.root.join("tmp", "job_results.txt")
    return [] unless File.exist?(result_file)

    File.readlines(result_file).reverse.first(10)
  end
end
