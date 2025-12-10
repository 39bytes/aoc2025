struct Button
    def initialize(@toggles : Array(Int32))
    end

    def press_lights(lights : Int32)
        @toggles.each do |t|
            lights ^= (1 << t)
        end
        lights
    end

    def increment_joltages(joltages : Array(Int32))
        nxt = joltages.dup
        @toggles.each do |t|
            nxt[t] += 1
        end
        nxt
    end
end

struct Machine
    property lights, buttons, joltages

    def initialize(@lights : Int32, @buttons : Array(Button), @joltages : Array(Int32))
    end
end

def min_light_presses(machine : Machine)
    q = Deque.new([{0, 0}])
    visited = Set(Int32).new
    visited.add(0)

    until q.empty?
        cur, dist = q.shift
        
        machine.buttons.each do |b|
            nxt = b.press_lights(cur)
            if nxt == machine.lights
                return dist + 1
            end
            next if visited.includes?(nxt)
            q << {nxt, dist + 1}
            visited.add(nxt)
        end
    end
    raise "oops"
end

def part1(machines : Array(Machine))
    machines.map { |m| min_light_presses(m) }.sum
end

machines = File.read("inputs/10/example.txt").strip.split("\n").map do |line|
    parts = line.split
    light_config = parts[0][1...-1]
    lights = 0
    light_config.chars.reverse.each do |c|
        lights <<= 1
        lights |= 1 if c == '#'
    end
    buttons = parts[1...-1].map do |button|
        nums = button[1...-1].split(',').map { |x| x.to_i32 }
        Button.new(nums)
    end
    joltages = parts[-1][1...-1].split(',').map { |x| x.to_i32 }
    
    Machine.new(lights, buttons, joltages)
end

puts "Part 1: #{part1(machines)}"
