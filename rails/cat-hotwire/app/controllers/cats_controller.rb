class CatsController < ApplicationController
  before_action :set_cat, only: %i[ show edit update destroy ]

  include Pagy::Backend

  # GET /cats
  def index
   @search = Cat.ransack(params[:q])
   @search.sorts = "id desc" if @search.sorts.empty?
   @pagy, @cats = pagy(@search.result, limit: 10)
  end

  # GET /cats/1
  def show
  end

  # GET /cats/new
  def new
    @cat = Cat.new
  end

  # GET /cats/1/edit
  def edit
  end

  # POST /cats
  def create
    @cat = Cat.new(cat_params)

    if @cat.save
      redirect_to @cat, notice: "Cat was successfully created."
    else
      render :new, status: :unprocessable_entity
    end
  end

  # PATCH/PUT /cats/1
  def update
    if @cat.update(cat_params)
      redirect_to @cat, notice: "Cat was successfully updated.", status: :see_other
    else
      render :edit, status: :unprocessable_entity
    end
  end

  # DELETE /cats/1
  def destroy
    @cat.destroy!
    redirect_to cats_path, notice: "Cat was successfully destroyed.", status: :see_other
  end

  private
    # Use callbacks to share common setup or constraints between actions.
    def set_cat
      @cat = Cat.find(params.expect(:id))
    end

    # Only allow a list of trusted parameters through.
    def cat_params
      params.expect(cat: [ :name, :age ])
    end
end
