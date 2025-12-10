class UnionFind
    getter components : Int64

    def initialize(size : Int64)
        @parent = (0...size).to_a
        @size = Array(Int64).new(size, 1)
        @components = size
    end

    def find(x : Int32)
        return x if @parent[x] == x
        @parent[x] = find(@parent[x])
        @parent[x]
    end

    def union(x : Int32, y : Int32)
        x = find(x)
        y = find(y)
        return if x == y

        if @size[x] < @size[y]
            x, y = y, x
        end

        @parent[y] = x
        @size[x] += @size[y]
        @components -= 1
    end

    def top3
        @size.sort.reverse[...3]
    end
end
 
struct Node
    property i, x, y, z

    def initialize(@i : Int32, @x : Int64, @y : Int64, @z : Int64)
    end

    def dist_squared(other : Node)
        (@x - other.x) ** 2 + (@y - other.y) ** 2 + (@z - other.z) ** 2
    end
end

def compute_dist_pairs(nodes : Array(Node))
    dist_pairs = nodes.combinations(2).map do |(a, b)|
        {a.dist_squared(b), a.i, b.i}
    end.sort_by { |x| x[0] }
end

MAX_CONNECTIONS = 1000
def part1(nodes : Array(Node), pairs : Array(Tuple(Int64, Int32, Int32)))
    uf = UnionFind.new(nodes.size)
    (0...MAX_CONNECTIONS).each do |i|
        uf.union(pairs[i][1], pairs[i][2])
    end
    uf.top3.product
end

def part2(nodes : Array(Node), pairs : Array(Tuple(Int64, Int32, Int32)))
    uf = UnionFind.new(nodes.size)
    pairs.each do |p|
        uf.union(p[1], p[2])
        if uf.components == 1
            return nodes[p[1]].x * nodes[p[2]].x
        end
    end
    -1
end
 
nodes = File.read("inputs/08/input.txt").strip.split('\n').map_with_index do |line, i|
    x, y, z = line.split(',')
    Node.new(i, x.to_i64, y.to_i64, z.to_i64)
end
pairs = compute_dist_pairs(nodes)

puts "Part 1: #{part1(nodes, pairs)}"
puts "Part 2: #{part2(nodes, pairs)}"
