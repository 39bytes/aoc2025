content = File.read("inputs/04/input.txt")
grid = content.strip.split("\n").map {|s| s.chars }

DIRS = [[0, -1], [0, 1], [-1, 0], [1, 0], [1, 1], [1, -1], [-1, 1], [-1, -1]]

def get_accessible(grid)
    n = grid.size
    m = grid[0].size
    accessible = [] of {Int32, Int32}
    (0...n).each do |i|
        (0...m).each do |j|
            next if grid[i][j] != '@'
            neighbors = DIRS.count do |(di, dj)|
                (0...n).covers?(i+di) && (0...m).covers?(j+dj) && grid[i+di][j+dj] == '@'
            end
            accessible << {i, j} if neighbors < 4
        end
    end
    accessible
end

def part1(grid)
    get_accessible(grid).size
end

def strip_accessible(grid, accessible)
    r = 0
    grid.map_with_index do |row, i|
        row.map_with_index do |char, j|
            if accessible[r]? == {i, j}
                r += 1
                '.'
            else
                char
            end
        end
    end
end

def part2(grid)
    ans = 0
    accessible = get_accessible(grid)
    until accessible.empty?
        ans += accessible.size
        grid = strip_accessible(grid, accessible)
        accessible = get_accessible(grid)
    end
    ans
end

puts "Part 1: #{part1(grid)}"
puts "Part 2: #{part2(grid)}"
