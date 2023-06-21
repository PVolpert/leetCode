# Number Manipulation

Numbers are a important factor in programming languages. Advanced number manipulation often is done by manipulating the converted string of a number. Most of the time string conversion is not necessary. The following methods purely use mathematic properties to manipulate numbers.

## Base Conversion

### Doubling and Adding

1. Start with the given number in base x.
2. Initialize the result to 0.
3. Begin with the leftmost digit (the most significant digit) of the number.
4. Multiply the current value of the result by the base y.
5. Add the value of the current digit (in base y) to the result.
6. Move to the next digit to the right and repeat steps 4 and 5 until you reach the rightmost digit.
7. The final value of the result will be the equivalent representation of the given number in base y.

Doubling and Adding can be used to convert bases. It runs in O(n) and has the advantage that it starts from the leftmost digit.

## Extracting digits

### Rightmost digit

1. Start with the given number in base x.
2. digit = number % base
3. (optional) number = number / base

### Leftmost digit & Length of number

1. Get the absolute value of the number as temp
2. Set length = 1
3. Temp = Temp / base
4. Increment length
5. Repeat 4&5 until temp < base
6. The final value of temp is the leftmost digit and final value of length is length of number

With pure mathematical manipulation we can neither access length or leftmost digit of a number directly. As such we need to extract all digits until we only have one remaining digit.
