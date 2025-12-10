from dataclasses import dataclass
import numpy as np
from scipy.optimize import milp, LinearConstraint


@dataclass
class Machine:
    buttons: list[list[int]]
    joltages: list[int]


with open("inputs/10/input.txt") as f:
    machines: list[Machine] = []

    for line in f.readlines():
        parts = line.strip().split()
        btns = parts[1:-1]

        buttons = [list(map(int, b[1:-1].split(","))) for b in btns]
        joltages = list(map(int, parts[-1][1:-1].split(",")))

        machines.append(Machine(buttons, joltages))


def solve(m: Machine):
    obj = [1] * len(m.buttons)

    constraints = [[0] * len(m.buttons) for _ in range(len(m.joltages))]
    for i, button in enumerate(m.buttons):
        for j in button:
            constraints[j][i] = 1
    res = milp(
        c=obj,
        constraints=LinearConstraint(lb=m.joltages, ub=m.joltages, A=constraints),
        integrality=np.ones(len(m.buttons)),
    )

    return int(sum(res.x))


print(sum(solve(m) for m in machines))
