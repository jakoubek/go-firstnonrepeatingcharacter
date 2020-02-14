# go-firstnonrepeatingcharacter

Get the first non-repeating character in a string - my solution in Go

## Approach

- create a map for the characters of the string
- iterate over all characters (runes) of the string and increment the number of times this character has been seen
- iterate a second time over all characters
- look in the map for the current character and return it if the number is 1
