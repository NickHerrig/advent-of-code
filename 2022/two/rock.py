

point_ledger_one = {
    "A": {"move":"rock", "points": 1},
    "B": {"move": "paper", "points": 2},
    "C": {"move": "scissors", "points": 3},
    "X": {"move": "rock", "points": 1},
    "Y": {"move": "paper", "points": 2},
    "Z": {"move": "scissors","points": 3},
}

point_ledger_two= {
    "A": {"move":"rock", "points": 1},
    "B": {"move": "paper", "points": 2},
    "C": {"move": "scissors", "points": 3},
    "X": "lose",
    "Y": "draw",
    "Z": "win",
}

move_point_ledger = {
    "rock": 1,
    "paper": 2,
    "scissors": 3,
}

# This function returns the move that will win or lose or tie
def determine_move(elfs_move, needed_outcome):
    if needed_outcome == "win":
        if elfs_move == "rock":
            return "paper"
        elif elfs_move == "paper":
            return "scissors"
        else:
            return "rock"
    elif needed_outcome == "lose":
        if elfs_move == "rock":
            return "scissors"
        elif elfs_move == "paper":
            return "rock"
        else:
            return "paper"
    else:
        return elfs_move


# This function returns the points for each specific move
def shape_outcome(my_move, elfs_move):
    return (point_ledger_one[my_move]["points"], point_ledger_one[elfs_move]["points"])

# This function returns the points for each specific move
def shape_outcome_two(my_move, elfs_move):
    return (move_point_ledger[my_move], move_point_ledger[elfs_move])

# This function returns the points outcome of the game
def determine_winner(my_move, elfs_move):
  if my_move == elfs_move:
    return (3, 3)

  # my_move wins
  if (my_move == "rock" and elfs_move == "scissors") or (my_move == "scissors" and elfs_move == "paper") or (my_move == "paper" and elfs_move == "rock"):
    return (6, 0)

  # elfs_move wins
  else:
    return (0, 6)



def main():

    lines = []
    with open('input.txt', 'r') as f:
        turns = [line.split() for line in f]

    my_points_total_one, my_points_total_two = 0, 0

    #Problem 1 elfs_move vs my_move
    for my_move, elfs_move in turns:
        winner_points = determine_winner(point_ledger_one[my_move]["move"], point_ledger_one[elfs_move]["move"])
        shape_points = shape_outcome(my_move, elfs_move) 
        my_points_total_one += winner_points[1] + shape_points[1]

    for elfs_move, outcome in turns:
        my_move = determine_move(point_ledger_two[elfs_move]["move"], point_ledger_two[outcome])
        winner_points = determine_winner(my_move, point_ledger_two[elfs_move]["move"])
        shape_points = shape_outcome_two(my_move, point_ledger_two[elfs_move]["move"])
        my_points_total_two += winner_points[0] + shape_points[0]

    print(my_points_total_one)
    print(my_points_total_two)


if __name__=="__main__":
    main()