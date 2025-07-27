# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Rails 8 application demonstrating Sidekiq background job processing with Turbo Streams for real-time UI updates. The app showcases asynchronous job processing with immediate user feedback through server-sent events.

## Commands

### Development Setup
```bash
# Start the entire application stack
docker compose up --build

# View logs
docker compose logs web        # Rails app logs
docker compose logs sidekiq    # Sidekiq worker logs
docker compose logs redis      # Redis logs

# Access containers
docker compose exec web bash
docker compose exec sidekiq bash
```

### Rails Commands
```bash
# Database operations
bundle exec rails db:create
bundle exec rails db:migrate
bundle exec rails db:prepare

# Console access
bundle exec rails console

# Manual job execution (in console)
SampleJob.perform_async("Test User", "Manual job execution")
EmailJob.perform_async("test@example.com", "Test Subject", "Test Body")
ScheduledJob.perform_async("health_check")

# Scheduled jobs management (in console)
Sidekiq.schedule  # View current schedule
SidekiqScheduler::Scheduler.instance.reload_schedule!  # Reload schedule
```

### Code Quality & Security
```bash
# Run RuboCop linter
bundle exec rubocop

# Run Brakeman security scanner
bundle exec brakeman
```

### Testing
```bash
# Run all tests
bundle exec rails test

# Run specific test file
bundle exec rails test test/path/to/test_file.rb
```

## Architecture

### Core Components
- **Rails 8 Web App**: Main application server serving UI and API endpoints
- **Sidekiq Workers**: Background job processors running in separate containers
- **Redis**: Job queue storage and Turbo Streams message broker
- **Turbo Streams**: Real-time DOM updates without JavaScript

### Job Processing Flow
1. User submits job via web form (`JobsController`)
2. Job queued to Redis via Sidekiq (`SampleJob.perform_async`)
3. Sidekiq worker processes job in background
4. Job completion triggers Turbo Streams broadcast
5. UI updates automatically via WebSocket connection

### Key Files
- `app/jobs/sample_job.rb`: Default queue job (5s processing time)
- `app/jobs/email_job.rb`: High priority queue job (3s processing time)
- `app/jobs/scheduled_job.rb`: Scheduled background jobs (health check, cleanup, daily summary)
- `app/controllers/jobs_controller.rb`: Main controller handling job submission
- `config/initializers/sidekiq.rb`: Sidekiq and sidekiq-scheduler configuration
- `config/schedule.yml`: Scheduled job definitions
- `app/views/jobs/index.html.erb`: Main UI with Turbo Streams integration

### Turbo Streams Integration
- Jobs broadcast completion via `Turbo::StreamsChannel.broadcast_prepend_to`
- Real-time notifications appear in `#notifications` div
- Results list updates automatically in `#job_results_content`
- No JavaScript required for real-time updates

### Queues
- `default`: Standard priority jobs (SampleJob)
- `high_priority`: High priority jobs (EmailJob)
- `scheduled`: Scheduled jobs (ScheduledJob)

## Development Notes

### Accessing Services
- Main app: http://localhost:3000
- Sidekiq Web UI: http://localhost:3000/sidekiq
- Redis: localhost:6379

### Job Results Storage
- Results stored in `tmp/job_results.txt` for demo purposes
- Last 10 results displayed in UI
- Can be cleared via "Clear Results" button

### Turbo Streams Testing
```ruby
# Manual broadcast test (in Rails console)
Turbo::StreamsChannel.broadcast_prepend_to(
  "job_notifications", 
  target: "notifications", 
  html: "<div>Test notification</div>"
)
```

### Queue Monitoring
```ruby
# Check queue status (in Rails console)
Sidekiq::Stats.new.queues
Sidekiq::Queue.new("default").size
Sidekiq::Queue.new("high_priority").size
Sidekiq::Queue.new("scheduled").size

# Check scheduled jobs
Sidekiq::Cron::Job.all  # View all scheduled jobs
```

### Scheduled Jobs
The application includes three types of scheduled jobs:

1. **Health Check** (`*/2 * * * *`): Runs every 2 minutes
   - Monitors system metrics (memory, CPU usage)
   - Checks queue sizes
   - Provides system status updates

2. **Cleanup Task** (`*/10 * * * *`): Runs every 10 minutes  
   - Removes old job result entries
   - Keeps only the latest 50 results
   - Maintains clean data storage

3. **Daily Summary** (`0 0 * * *`): Runs daily at 9 AM JST
   - Generates job processing statistics
   - Reports success/failure rates
   - Provides daily performance metrics

Schedule configuration is managed in `config/schedule.yml` with cron expressions.