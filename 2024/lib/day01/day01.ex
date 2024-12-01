defmodule Day01 do
  def prepare_input(filename) do
    File.stream!(filename)
    |> Stream.map(&String.split/1)
    |> Stream.map(fn pair -> Enum.map(pair, &String.to_integer/1) end)
    |> Enum.to_list()
    |> Enum.reduce([[], []], fn [a, b], [list1, list2] -> [[a | list1], [b | list2]] end)
  end

  # Parte 01
  def calculate_total_distance([_list1, _list2] = l) do
    Enum.map(l, &Enum.sort/1)
    |> List.zip()
    |> Enum.reduce(0, fn {a, b}, acc -> acc + abs(a - b) end)
  end

  # Parte 02
  def calculate_similarity([list1, list2]) do
    list2_freq = Enum.frequencies(list2)

    # Enum.reduce(list1, 0, fn x, acc -> list2_freq[x] * x + acc end)
    Enum.reduce(list1, 0, fn x, acc -> x * Map.get(list2_freq, x, 0) + acc end)
  end

  def part01(filename) do
    prepare_input(filename)
    |> calculate_total_distance()
  end

  def part02(filename) do
    prepare_input(filename)
    |> calculate_similarity()
  end
end

IO.inspect("Dia 1 - Parte 1:  #{Day01.part01("./inputs/part01.txt")}")
IO.inspect("Dia 1 - Parte 2:  #{Day01.part02("./inputs/part01.txt")}")
