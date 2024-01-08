import random
import argparse
from dataclasses import dataclass
from typing import List


@dataclass
class DieCast:
    """A single die cast."""
    type: str
    val: int


@dataclass
class RollRequest:
    """A request to roll a number of dice with a given number of sides."""
    name: str
    num: int
    sides: int

    def roll(self) -> List[DieCast]:
        """
        Roll the dice the specified number of times and return the results.
        """
        return [
            DieCast(
                type=self.sides,
                val=random.randint(1, self.sides)
            ) for x in range(self.num)
        ]


def get_roll_requests() -> List[RollRequest]:
    """Parse the command line arguments and return a list of roll requests."""
    try:
        parser = argparse.ArgumentParser()
        parser.add_argument('dice', nargs='+')
        args = parser.parse_args()
        requests: List[RollRequest] = []
        for request in args.dice:
            num, sides = request.split('d')
            requests.append(
                RollRequest(name=request, num=int(num), sides=int(sides))
            )
        return requests
    except Exception as e:
        print(f"Usage: python {__file__} NdM [NdM ...]")


def print_results(results: List[DieCast], total: int) -> None:
    """Print the results of the die rolls."""
    print(f"Total: {total}")
    for result in results:
        print(f"   d{result.type}: {result.val}")


def roll(requests: List[RollRequest]) -> List[DieCast]:
    """Roll the dice and return the results."""
    results: List[DieCast] = []
    for req in requests:
        results += req.roll()
    return results


def main() -> None:
    # Parse the command line arguments to get the roll requests.
    roll_requests = get_roll_requests()

    # Roll the dice and get the results.
    results = roll(roll_requests)
    # Calculate the total of all the dice.
    total = sum(r.val for r in results)

    # Print the results.
    print_results(results, total)


if __name__ == "__main__":
    main()