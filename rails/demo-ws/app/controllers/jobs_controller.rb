class JobsController < ApplicationController
  def index
    respond_to do |format|
      format.html # Render the web UI
      format.json do
        render json: {
          message: "Sidekiq Jobs API",
          endpoints: {
            simple: "/jobs/simple"
          }
        }
      end
    end
  end

  def simple
    message = params[:message] || "Hello from Sidekiq!"
    delay_seconds = params[:delay]&.to_i || 5

    job = SimpleJob.perform_later(message, delay_seconds)

    job_data = {
      status: "Job enqueued",
      job_id: job.job_id,
      message: message,
      delay_seconds: delay_seconds,
      queue: job.queue_name,
      type: "simple"
    }

    respond_to do |format|
      format.json { render json: job_data }
      format.html { redirect_to root_path }
    end
  end
end
