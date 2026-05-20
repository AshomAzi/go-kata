## The Secure String Capitalizer (Truncator)
* **The Goal**: Master basic multi-value returns and coordinate bounds validation without using external libraries.
```
Write a function named SecureTruncate that accepts two parameters: a string text target and an integer limit boundary.
```

* **The Rules**:
```
If the limit parameter is less than or equal to 0, return an empty string and a descriptive error using fmt.Errorf.

If the length of the string is greater than the limit, slice the string safely down to that limit, append "..." to the end, and return the modified string along with a nil error.

If the string length is within bounds, return the original string completely intact and a nil error.```


During implementation, I converted the input into slice of rune to handle characters which takes more than one byte and then convert it back to string