Dodgeball Roster Lineup Generator
=================================

Trying to use a computer algorithm to generate lineups for my dodgeball team. 
Ideally all players play the same amount of games, but each lineup should have
the maximum amount of power. Players should not have to play more than 2 games in a row.

Requirements
------------

1. there should be no more than 8 players at a time.
2. A minimum of 3 girls is required at all times.
3. Players should not be able to play more than 2 games in a row.
4. Players should roughtly get equal playing time

Permutations & Combinations
--------------------

### Permutations


__permutations__: consider the same combinations in different orders to be unique. For example: {a, b, c} and {c, b, a} are different permutations, but are a single combination.

### Power Set

If we simply wanted to know how many combinations of lineups regardless of gender distribution and team size, we can simply find the size of the power set. 2**n

	2**14 = 16,384

This assumes that a team size of 1-14 are all valid combinations including a team of no players and a team of all players. This doesn't work as we need teams of 8 with 5 men and 3 women.

Since we don't want to consider the order of the lineup as a unique indicator, we need to use combinations.

### Combinations

__combination__: a way of selecting members from a grouping, such that (unlike permutations) the order of selection does not matter.

For sports teams the order of the lineups does not matter, therefore we need all _combinations_ of 8 men and 3 women lineups without replacement.

__k-combination__: commonly known as "n choose k". a subset of k distinct elements of S. If you have 14 players but can only choose 8 for a single lineup, then all possible combinations of 8 are the k-combinations. Another way to say this is "n choose k" or "14 choose 8" The number of k-combinations is equal to the binomial coefficient:

__k-selection or k-multiset__: refers to combinations where repition of items is allowed. This doesn't apply to the lineup as a player can only be on the lineup once.

- n: number of items
- k: number of chosen items
- !: factorial

	n! / k!(n-k)!

We can split up the lineups into men and women combinations.

	men_combinations = 7! / (7-5)! * 5! = 21
	women_combinations = 7! / (7-3)! * 3! = 35

Then we can multiply the two to get the total number of combinations for lineups.

	21 * 35 = 735 possible combinations of 8m 3w rosters.

### Resources

[Permutations and Combinations](http://ltcconline.net/greenl/courses/102/Probability/permutations_and_combinations.htm)

Potential Algorithms
--------------------

### 0-1 Knapsack Problem

> Given some items, pack the knapsack to get the maximum total value. Each item has some weight and some value. Total weight that we can carry is no more than some fixed number W. So we must consider weights of items as well as their value.

I can translate a Player to have a value or power rating, but how do I handle the weight? Can the weight correlate to the number of games the player has already played? I could use a cool down method to reduce the weight after each linuep is generated.

To Do
-----

- [ ] What is the [Twelvefold way](http://en.wikipedia.org/wiki/Twelvefold_way)?
