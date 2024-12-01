defmodule Day01.Tests do
  use ExUnit.Case
  alias Day01

  test "testing part 01 with example" do
    example_file = Path.join(__DIR__, "example.txt")
    assert Day01.part01(example_file) == 11
  end

  test "testing part 02 with example" do
    example_file = Path.join(__DIR__, "example.txt")
    assert Day01.part02(example_file) == 31
  end
end
