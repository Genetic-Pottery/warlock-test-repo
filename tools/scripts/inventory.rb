module Inventory
  SHELF_LIMIT = 40

  class Shelf
    def initialize(items)
      @items = items
    end

    def count
      total = 0
      @items.each { |i| total += i }
      total
    end
  end

  def self.build(items)
    Shelf.new(items)
  end
end
