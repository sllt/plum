#
# Printer
#

defmodule Printer do
    def pr_str(value, readably) do
        case value do
            nil ->
                "nil"
            true ->
                "true"
            false ->
                "false"
            {:integer, num} ->
                integer_to_list(num)
            {:string, str} ->
                IO.puts(Printer.escape(str,readably))
            {:keyword, keyword} ->
                IO.puts ?keyword
            {:symbol, sym} ->
                IO.puts sym
            {:list, list} ->
                Printer.pr_list(list, ?(, ?), readably)
            {:vector, vec} ->
                Printer.pr_list(vec, ?(, ?), readably)
            {:map, map} ->
                Printer.pr_map(map, readably)
        end
    end

    def pr_list(seq, seq_start, seq_end, readably) do
        
    end

    def pr_map(map, readably) do
        
    end

    def escape(str, false) do
        "\"" <> str <> "\""
    end

    def escape(str, true) do
        
    end
    
    
end