defmodule Pipeline.Stage do
  @moduledoc "One stage of the pipeline."

  def run(items) do
    Enum.reduce(items, 0, fn item, acc -> acc + item end)
  end

  defp normalise(item), do: item
end
