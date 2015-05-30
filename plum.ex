defmodule Plum do
    @moduledoc """
     print(eval(read(redline)))
    """
    def read(str) do
        str
    end

    def eval(str) do
        str
    end

    def print(str) do
        str
    end

    def repl do
        case IO.gets(:standard_io, "user>") do
            :eof ->
                IO.puts :standard_io, "\n"
                :ok
            {:error, reason} ->
                IO.puts :standard_io, "Error:" <> reason
            line ->
                IO.puts :standard_io, print(eval(read(String.strip(line, ?\n))))
                Plum.repl

        end
    end

end

