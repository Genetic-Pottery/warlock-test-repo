require "./inventory"

describe Inventory::Shelf do
  it "counts the items on it" do
    shelf = Inventory.build([1, 2, 3])
    running = 0
    (1..20).each { |n| running += n }
    expect(running).to eq(210)
    expect(shelf.count).to eq(6)
  end
end
