# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Rails Application Management
- `bin/rails server` - Start the Rails development server
- `bin/rails console` - Open Rails interactive console
- `bin/rails generate` - Run Rails generators (controllers, models, migrations, etc.)

### Database Operations
- `bin/rails db:create` - Create development and test databases
- `bin/rails db:migrate` - Run pending migrations
- `bin/rails db:prepare` - Setup database if it doesn't exist, or run migrations
- `bin/rails db:seed` - Populate database with seed data
- `bin/rails db:reset` - Drop, recreate, and migrate database

### Asset Management
- `bin/rails assets:precompile` - Compile assets for production
- `bin/rails assets:clean` - Remove old compiled assets

### Code Quality & Security
- `bundle exec rubocop` - Run Ruby linter (Omakase style)
- `bundle exec brakeman` - Run security vulnerability scanner

### Testing
- `bin/rails test` - Run test suite (Rails minitest)
- `bin/rails test:system` - Run system tests (disabled by default)

### Background Jobs (Sidekiq)
- `bundle exec sidekiq` - Start Sidekiq worker process
- `docker compose up -d` - Start Redis service (required for Sidekiq)
- Access Sidekiq Web UI at `http://localhost:3000/sidekiq` (development only)

### Docker & Deployment
- `docker compose up` - Start Redis service for development
- `bin/kamal setup` - Initial deployment setup using Kamal
- `bin/kamal deploy` - Deploy application
- `bin/kamal console` - Access deployed application console
- `bin/kamal logs` - View deployment logs

## Architecture Overview

### Rails 8 Modern Stack
This is a Rails 8 application using the modern default stack:
- **Asset Pipeline**: Propshaft (modern asset pipeline)
- **JavaScript**: Importmap-rails with Stimulus and Turbo
- **Database**: SQLite3 with multiple databases for production (primary, cache, queue, cable)
- **Background Jobs**: Sidekiq (Redis-backed) - replaced Solid Queue
- **Caching**: Solid Cache (database-backed)
- **WebSockets**: Solid Cable (database-backed) via Action Cable
- **Deployment**: Kamal for containerized deployments

### Database Configuration
- **Development/Test**: Single SQLite database
- **Production**: Multi-database setup with separate databases for:
  - Primary application data
  - Cache storage (Solid Cache)
  - Background jobs (Solid Queue)
  - WebSocket connections (Solid Cable)

### Key Configuration Files
- `config/application.rb` - Main application configuration with Rails 8 defaults
- `config/database.yml` - Multi-database setup for production scaling
- `config/deploy.yml` - Kamal deployment configuration
- `config/routes.rb` - Currently minimal with health check endpoint
- `Gemfile` - Rails 8 with solid_* gems for modern Rails infrastructure

### Application Structure
- Standard Rails MVC structure with minimal controllers/models (new application)
- PWA support built-in (manifest and service worker templates)
- Custom error pages in `public/` directory
- Docker support via Dockerfile and Kamal deployment

### External Services
- **Redis**: Required for Sidekiq background jobs, available via docker-compose
- **Kamal**: Configured for production deployment to servers

### Background Job Implementation
This application includes Sidekiq for background job processing with real-time broadcasting:
- `SimpleJob` - Demonstrates background job with real-time progress updates every second
- Shows elapsed time / remaining time format: "進行状況: 3秒経過 / 残り2秒"
- Uses Turbo Streams to broadcast job status and log messages in real-time

API endpoints:
- `GET /jobs` - Web UI for job execution
- `POST /jobs/simple?message=hello&delay=5` - Enqueue a simple job

## Development Workflow

When working with this codebase:
1. Use `bin/rails` commands instead of global `rails` command
2. Run `bundle install` after pulling changes to update gems
3. Use `bin/rails db:prepare` to setup database for new environments
4. Test security changes with `bundle exec brakeman`
5. Follow Rails conventions and use built-in generators
6. Start Redis with `docker compose up -d` before running background jobs
7. Use `bundle exec sidekiq` to start the job worker process
8. The application has been modified to use Sidekiq instead of Solid Queue for background jobs