class CreateHedgehogs < ActiveRecord::Migration[8.0]
  def change
    create_table :hedgehogs do |t|
      t.string :name
      t.integer :age

      t.timestamps
    end
  end
end
