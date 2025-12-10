from dataclasses import dataclass
import sys
import z3


@dataclass
class ButtonsAndTarget:
    buttons: list[list[int]]
    target: list[int]


def parse_buttons_and_target(fields: list[str]) -> ButtonsAndTarget:
    buttons = [
        [int(j) for j in i.strip("()").split(",")] for i in fields[1 : len(fields) - 1]
    ]
    target = [int(i) for i in fields[len(fields) - 1].strip("{}").split(",")]
    return ButtonsAndTarget(buttons=buttons, target=target)


def solve_constraints(bt: ButtonsAndTarget) -> int:
    s = z3.Optimize()
    bc = [z3.Int(f"b{i}") for i in range(len(bt.buttons))]
    bc_sum = 0
    for i in bc:
        s.add(i >= 0)
        bc_sum += i
    for n, i in enumerate(bt.target):
        eq = 0
        for m, j in enumerate(bt.buttons):
            if n in j:
                eq += bc[m]
        s.add(i == eq)
    s.minimize(bc_sum)
    if s.check() != z3.sat:
        raise Exception("constraint not solvable")
    model = s.model()
    return model.eval(bc_sum).as_long()


sum = 0
for l in sys.stdin:
    bt = parse_buttons_and_target(l.rstrip().split(" "))
    sum += solve_constraints(bt)

print("Part 2:", sum)
