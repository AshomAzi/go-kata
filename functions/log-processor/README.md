# Log Processor
Let's combine defer and a variance concept to test your understanding!

Write a program with a function named LogBatch.

## The Rules:

The function signature must accept a string prefix tag (like "INFO" or "ERROR"), followed by a variadic number of message strings.

The very first line inside LogBatch must use defer to print a closing summary statement: "[SYSTEM] Batch processing complete."

Loop through all the passed messages and print them out formatted like this: "[PREFIX] Message text".

## Signature Guide:

```go
func LogBatch(prefix string, messages ...string)
```