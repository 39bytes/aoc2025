def invalid?(i : Int64)
    return false if i < 10
    digits = i.digits
    half = digits.size // 2
    digits[..half-1] == digits[half..]
end

def invalid2?(i : Int64)
    return false if i < 10
    digits = i.digits
    half = digits.size // 2
    (1..half).any? do |j|
        if digits.size % j != 0
            false
        else
            chunks = digits.in_slices_of(j)
            chunks.all?{ |c| c == chunks[0] }
        end
    end
end

content = File.read("inputs/02/input.txt")
ranges = content.split(",").map do |s| 
    a, b = s.split "-"
    a.to_i64..b.to_i64
end
part1 = ranges.map { |r| r.select { |i| invalid? i } }.flatten.sum
puts "Part 1: #{part1}"

part2 = ranges.map { |r| r.select { |i| invalid2? i } }.flatten.sum
puts "Part 2: #{part2}"
