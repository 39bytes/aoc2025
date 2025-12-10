def part1(lines : Array(String))
    nums = [] of Array(Int64)
    ops = [] of String
    lines.each do |line|
        if line.starts_with?('*') || line.starts_with?('+')
            ops = line.split
        else
            nums << line.split.map { |x| x.to_i64 }
        end
    end

    nums.transpose.zip(ops).map do |nums, op| 
        if op == "*"
            nums.product
        else
            nums.sum
        end
    end.sum
end

def part2(lines : Array(String))
    ans : Int64 = 0
    nums = [] of Int64
    (lines[0].size - 1).downto(0).each do |j|
        num = ""
        (0...lines.size).each do |i|
            c = lines[i][j]
            if c.hex?
                num += c
            end
            if c == '+' || c == '*' 
                nums << num.to_i64
                num = ""
                ans += nums.sum if c == '+'
                ans += nums.product if c == '*'
                nums = [] of Int64
            end
        end 
        if num != ""
            nums << num.to_i64
            num = ""
        end
    end

    ans
end

lines = File.read("inputs/06/input.txt").strip("\n").split("\n")

puts "Part 1: #{part1(lines)}"
puts "Part 2: #{part2(lines)}"
