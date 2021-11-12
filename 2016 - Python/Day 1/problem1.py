from enum import Enum, auto


class Direction(Enum):
    NORTH = auto()
    EAST = auto()
    SOUTH = auto()
    WEST = auto()


def move(direction, position, distance):
    x, y = position
    if direction is Direction.NORTH:
        y += distance
    elif direction is Direction.EAST:
        x += distance
    elif direction is Direction.SOUTH:
        y -= distance
    elif direction is Direction.WEST:
        x -= distance

    return (x, y)


def turn(direction, lr):
    if lr == "R":
        if direction is Direction.NORTH:
            return Direction.EAST
        elif direction is Direction.EAST:
            return Direction.SOUTH
        elif direction is Direction.SOUTH:
            return Direction.WEST
        elif direction is Direction.WEST:
            return Direction.NORTH
    else:
        if direction is Direction.NORTH:
            return Direction.WEST
        elif direction is Direction.EAST:
            return Direction.NORTH
        elif direction is Direction.SOUTH:
            return Direction.EAST
        elif direction is Direction.WEST:
            return Direction.SOUTH


def walk(walk_path):
    position = (0, 0)
    direction = Direction.NORTH

    for command in walk_path.split(", "):
        lr, distance = command[0], int(command[1:])
        direction = turn(direction, lr)
        position = move(direction, position, distance)

    x, y = position
    return abs(x) + abs(y)


def walk_with_stop(walk_path):
    position = (0, 0)
    direction = Direction.NORTH
    seen = set()
    seen.add(str(position))

    for command in walk_path.split(", "):
        lr, distance = command[0], int(command[1:])
        direction = turn(direction, lr)
        position = move(direction, position, distance)
        print(position)
        if str(position) in seen:
            break
        else:
            seen.add(str(position))

    x, y = position
    return abs(x) + abs(y)


# Tests from problem
print(walk("R2, L3"), 5)
print(walk("R2, R2, R2"), 2)
print(walk("R5, L5, R5, R3"), 12)

print(walk(
    "R1, L3, R5, R5, R5, L4, R5, R1, R2, L1, L1, R5, R1, L3, L5, L2, R4, L1, R4, R5, L3, R5, L1, R3, L5, R1, "
    "L2, R1, L5, L1, R1, R4, R1, L1, L3, R3, R5, L3, R4, L4, R5, L5, L1, L2, R4, R3, R3, L185, R3, R4, L5, L4, "
    "R48, R1, R2, L1, R1, L4, L4, R77, R5, L2, R192, R2, R5, L4, L5, L3, R2, L4, R1, L5, R5, R4, R1, R2, L3, "
    "R4, R4, L2, L4, L3, R5, R4, L2, L1, L3, R1, R5, R5, R2, L5, L2, L3, L4, R2, R1, L4, L1, R1, R5, R3, R3, "
    "R4, L1, L4, R1, L2, R3, L3, L2, L1, L2, L2, L1, L2, R3, R1, L4, R1, L1, L4, R1, L2, L5, R3, L5, L2, L2, "
    "L3, R1, L4, R1, R1, R2, L1, L4, L4, R2, R2, R2, R2, R5, R1, L1, L4, L5, R2, R4, L3, L5, R2, R3, L4, L1, "
    "R2, R3, R5, L2, L3, R3, R1, R3"))

print(walk_with_stop("R8, R4, R4, R8"), 4)

print(walk_with_stop(
    "R1, L3, R5, R5, R5, L4, R5, R1, R2, L1, L1, R5, R1, L3, L5, L2, R4, L1, R4, R5, L3, R5, L1, R3, L5, R1, "
    "L2, R1, L5, L1, R1, R4, R1, L1, L3, R3, R5, L3, R4, L4, R5, L5, L1, L2, R4, R3, R3, L185, R3, R4, L5, L4, "
    "R48, R1, R2, L1, R1, L4, L4, R77, R5, L2, R192, R2, R5, L4, L5, L3, R2, L4, R1, L5, R5, R4, R1, R2, L3, "
    "R4, R4, L2, L4, L3, R5, R4, L2, L1, L3, R1, R5, R5, R2, L5, L2, L3, L4, R2, R1, L4, L1, R1, R5, R3, R3, "
    "R4, L1, L4, R1, L2, R3, L3, L2, L1, L2, L2, L1, L2, R3, R1, L4, R1, L1, L4, R1, L2, L5, R3, L5, L2, L2, "
    "L3, R1, L4, R1, R1, R2, L1, L4, L4, R2, R2, R2, R2, R5, R1, L1, L4, L5, R2, R4, L3, L5, R2, R3, L4, L1, "
    "R2, R3, R5, L2, L3, R3, R1, R3"))
