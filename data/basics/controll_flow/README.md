# If / Else

Here you can see some basic language mechanics for condition statements in `Ruby` & `Go`.
Note that you don't need parentheses around conditions in Go, but that the braces are required.

## Ruby

```ruby
# Here's a basic example.
if 7 % 2 == 0
  puts("7 is even")
else
  puts("7 is odd")
end

# You can have an `if` statement without an else.
if 8 % 4 == 0
  puts("8 is divisible by 4")
end

# A statement can precede conditionals; any variables
# declared in this statement are available in all
# branches.
if num = 9; num < 0
  puts(num, "is negative")
elsif num < 10
  puts(num, "has 1 digit")
else
  puts(num, "has multiple digits")
end
```

## Go

```go
package main

import "fmt"

func main() {
	// Here's a basic example.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// You can have an `if` statement without an else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// A statement can precede conditionals; any variables
	// declared in this statement are available in all
	// branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
```
