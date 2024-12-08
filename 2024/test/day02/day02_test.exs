defmodule Day02.Tests do
  use ExUnit.Case
  alias Day02

  test "testing part 01 with example" do
    example_file = Path.join(__DIR__, "example.txt")
    assert Day02.part01(example_file) == 2
  end

  test "testing part 02 with example" do
    example_file = Path.join(__DIR__, "example.txt")
    assert Day02.part02(example_file) == 4
  end
end
