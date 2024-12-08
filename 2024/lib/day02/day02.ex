defmodule Day02 do
  def prepare_input(filename) do
    File.stream!(filename)
    |> Stream.map(&String.split/1)
    |> Stream.map(fn line -> Enum.map(line, &String.to_integer/1) end)
    |> Enum.to_list()
  end

  # Part 01
  def check_level_safety([a, b | rest], decreasing) do
    cond do
      a - b >= 1 and a - b <= 3 and decreasing -> check_level_safety([b | rest], true)
      b - a >= 1 and b - a <= 3 and not decreasing -> check_level_safety([b | rest], false)
      true -> false
    end
  end

  def check_level_safety(_, _), do: true

  def check_levels(levels) do
    Enum.map(levels, fn level ->
      decreasing = List.first(level) > Enum.at(level, -1)
      check_level_safety(level, decreasing)
    end)
  end

  def part01(filename) do
    freq =
      filename
      |> prepare_input()
      |> check_levels()
      |> Enum.frequencies()

    freq[true]
  end

  # Part 2
  def check_levels_part02(levels) do
    Enum.map(levels, fn level ->
      decreasing = List.first(level) > Enum.at(level, -1)

      if check_level_safety(level, decreasing) do
        true
      else
        Enum.any?(0..(length(level) - 1), fn index ->
          reduced_level = List.delete_at(level, index)
          decreasing = List.first(level) > Enum.at(level, -1)
          check_level_safety(reduced_level, decreasing)
        end)
      end
    end)
  end

  def part02(filename) do
    freq =
      filename
      |> prepare_input()
      |> check_levels_part02()
      |> Enum.frequencies()

    freq[true]
  end
end

IO.inspect("Dia 2 - Parte 1:  #{Day02.part01("./inputs/day02/part01.txt")}")
IO.inspect("Dia 2 - Parte 2:  #{Day02.part02("./inputs/day02/part01.txt")}")
