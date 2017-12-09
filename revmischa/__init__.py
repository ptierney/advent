from revmischa.computer import Computer  # noqa
from logging import getLogger


getLogger('flake8').propagate = False

def main(computer_cls: Computer, puzzle_input: str=None):
    computer = computer_cls.new_from_puzzle_input(puzzle_input)
    answer_1 = computer.run_part1()
    print(f'Part I Answer: {answer_1}')
    computer = computer_cls.new_from_puzzle_input(puzzle_input)
    answer_2 = computer.run_part2()
    print(f'Part II Answer: {answer_2}')