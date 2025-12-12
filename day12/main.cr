struct Piece
    property shape
    def initialize(@shape : Array(Array(Char)))
    end

    def spaces
        @shape.map { |row| row.count('#') }.sum
    end
end
 
struct Region
    property width, height, required
    def initialize(@width : Int32, @height : Int32, @required : Array(Int32))
    end
end

def part1(regions : Array(Region), pieces : Array(Int32))
    regions.count do |r|
        total = r.width * r.height
        filled = r.required.zip(pieces).map { |(a, b)| a * b }.sum
        filled <= total
    end
end

parts = File.read("inputs/12/input.txt").strip.split("\n\n")
pieces = parts[...-1].map do |p|
    p.split('\n')[1..].map { |l| l.chars.count('#') }.sum
end

regions = parts[-1].split('\n').map do |line|
    dims, reqs = line.split(": ")
    width, height = dims.split('x')
    Region.new(width.to_i32, height.to_i32, reqs.split.map { |x| x.to_i32 })
end

puts "Part 1: #{part1(regions, pieces)}"
