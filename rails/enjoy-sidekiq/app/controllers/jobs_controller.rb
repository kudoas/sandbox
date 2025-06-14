class JobsController < ApplicationController
  def index
    @job_results = read_job_results
  end

  def create
    name = params[:name].presence || "Anonymous"
    message = params[:message].presence || "Hello from Sidekiq!"

    # 非同期ジョブをキューに追加
    SampleJob.perform_later(name, message)

    flash[:notice] = "🚀 ジョブがキューに追加されました！ 5秒後に完了予定です。"
    redirect_to jobs_path
  end

  def clear_results
    result_file = Rails.root.join("tmp", "job_results.txt")
    File.delete(result_file) if File.exist?(result_file)

    flash[:notice] = "🗑️ 結果をクリアしました。"
    redirect_to jobs_path
  end

  private

  def read_job_results
    result_file = Rails.root.join("tmp", "job_results.txt")
    return [] unless File.exist?(result_file)

    File.readlines(result_file).reverse.first(10)
  end
end
