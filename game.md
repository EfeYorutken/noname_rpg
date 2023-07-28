# TODO
- [x] impl character
- [x] impl appreal
- [x] imp weapon
- [x] imp consumable
- [x] imp enemy
-	[x] impl player
# no-name rpg
- rpg game

# magic
- each magic trick implements this
- has one method
- the method takes in 2 parameter
	- caster
	- target
- depending on the magic, the impelmentation of the method does something to either the caster, the target or both

# character
- has the following
	- name
	- appreal
	- weapon
	- health
	- attack (how likely is it to land a hit)
	- defence (how likely is it to avoid a hit)
	- stamina (how many actions can be taken per turn)
	- power (how much base power does a landed hit deal)
	- toughness (how much the power of a recieved hit reduced)

# items
## appreal
- has 
	- name
	- bonuses to base defence
- hat,chest, pants and boots implement this

## weapons
- has
	- name
	- bonus attack and power
- more than one can be held
- has 3 types
	- sword
	- hammer
	- axe

## consumables
- has
	- name
	- increases some value

# enemy
- implements __character__
- all in a level are kept in an array
- each one that is above a certain percentage of hp attack, one that are below but still somewhat high defend, rest try to heal if they can (on later levels enemies get more aggresive)

# player
- implements __character__
- each turn asks the user of an input
	- based on the input attacks, defends or uses an item
	- if incorrect input is entered does nothing

