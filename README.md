# code-katas

## Game of 3
Given any number, repeatedly divide the number by 3 till you reach 1.   
Add or subtract 1 whenever division by 3 is not possible.  
At each stage output the number.  
e.g. if you start with 100, it’s not divisible by 3, so subtract 1.  
100, 99  
Then keep dividing by 3  
100, 99, 33, 11  
11 is not divisible by 3, so add 1.  
100, 99, 33, 11, 12  
Keep going till you get to 1.  
100, 99, 33, 11, 12, 4, 3, 1  

## Rotate
Write a function that rotates a slice by k elements. For example [1,2,3,4,5,6] rotated by 2 becomes [3,4,5,6,1,2].  
Bonus: Try solving this without creating a copy of the slice.  

## Square or square root
Write a method, that will get an integer array as parameter and will process every number from this array.
Return a new array with processing every number of the input-array like this:  
If the number has an integer square root, take this, otherwise square the number.  

### Example
`[4,3,9,7,2,1] -> [2,9,3,49,4,1]`

### Notes
The input array will always contain only positive numbers, and will never be empty or null.

## Merge lists
Write a function that merges two sorted lists of the same length into a new sorted list.

`[1,4,6],[2,3,5] → [1,2,3,4,5,6]`

You can do this quicker than concatenating them followed by a sort.

## Letter sum value

Assign every lowercase letter a value, from 1 for a to 26 for z.

Given a string of lowercase letters, find the sum of the values of the letters in the string.

`lettersum("") => 0
lettersum("a") => 1
lettersum("z") => 26
lettersum("cab") => 6
lettersum("excellent") => 100
lettersum("microspectrophotometries") => 317`

## Fibonacci

The Fibonacci sequence, named after mathematician Fibonacci, is a sequence of numbers that looks like this:

`0, 1, 1, 2, 3, 5, 8, 13, ...`

You get first two starting numbers, 0 and 1, and the next number in the sequence is always the sum of the previous two numbers.

Write a function fib, that takes one parameter n, and returns the nth number from the Fibonacci sequence, counting from 0.

For example:

`fib(0) returns 0
fib(4) returns 3
fib(15) returns 610`

Note: Can your program compute fib(92)?

## Die Roll

Roll a six-sided die a million times. Print the frequency of each side coming up.

Note: Run the program a few times. Do you get the same numbers every time? How can you fix that?

## Sum of time durations

Given a list of time duration strings, calculate the sum of those durations.

Durations are in the hours:minutes:seconds format. For instance, 3:01:12 denotes 3 hours, 1 minute, and 12 seconds. Leading 0s are optional only for the hour parts.

Example:

SumOfTimes(["12:32", "34:01", "15:23", "9:27", "55:22", "25:56"]) should output 2:32:41
