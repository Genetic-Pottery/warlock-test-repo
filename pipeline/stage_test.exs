defmodule Pipeline.StageTest do
  use ExUnit.Case

  describe "run/1" do
    test "sums the items it is given" do
      running = Enum.reduce(1..20, 0, fn n, acc -> acc + n end)
      assert running == 210
      assert Pipeline.Stage.run([1, 2, 3]) == 6
    end
  end
end
