"""
https://app.codesignal.com/arcade/python-arcade/itertools-kit/d9Ru2ARE5tXoQ9KgR

The greatest ever Rock-Paper-Scissors Championship will take place in your town! The best players will battle each other to see who's the best player in the world. Each player will compete against each other player twice, once home and once away.

Given the list of the players, your task is to come up with the list of all the games between them, and return them sorted lexicographically.

Example

For players = ["trainee", "warrior", "ninja"], the output should be

rockPaperScissors(players) = [["ninja", "trainee"], ["ninja", "warrior"], 
                              ["trainee", "ninja"], ["trainee", "warrior"], 
                              ["warrior", "ninja"], ["warrior", "trainee"]]
"""
from itertools import permutations

def rockPaperScissors(players):
    return sorted(permutations(players, 2))
