module ProgressTrackable
  extend ActiveSupport::Concern

  def update_progress(current_step, total_steps, message = nil)
    progress_percentage = (current_step.to_f / total_steps * 100).round(1)

    # Redisにプログレス情報を保存
    Redis.new(url: ENV.fetch("REDIS_URL", "redis://localhost:6379/0")).setex(
      progress_key,
      300, # 5分でexpire
      {
        current_step: current_step,
        total_steps: total_steps,
        percentage: progress_percentage,
        message: message || "処理中...",
        updated_at: Time.current.iso8601
      }.to_json
    )

    # Turbo Streamsでプログレスバーを更新
    broadcast_progress_update(progress_percentage, current_step, total_steps, message)
  end

  def complete_progress(final_message = "完了しました！")
    # プログレスを100%に設定
    update_progress(100, 100, final_message)

    # 3秒後にプログレスバーを非表示
    sleep(1)
    broadcast_progress_hide
  end

  private

  def progress_key
    "job_progress:#{self.class.name}:#{jid}"
  end

  def broadcast_progress_update(percentage, current_step, total_steps, message)
    Turbo::StreamsChannel.broadcast_replace_to(
      "job_notifications",
      target: progress_target_id,
      partial: "shared/progress_bar",
      locals: {
        job_id: jid,
        job_type: self.class.name,
        percentage: percentage,
        current_step: current_step,
        total_steps: total_steps,
        message: message || "処理中..."
      }
    )
  end

  def broadcast_progress_hide
    Turbo::StreamsChannel.broadcast_remove_to(
      "job_notifications",
      target: progress_target_id
    )
  end

  def progress_target_id
    "progress_#{self.class.name}_#{jid}"
  end
end
