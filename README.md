# Dyck Languages

There's an older programming job interview question:

---

Given a string of round, curly, and square open and closing brackets,
return whether the brackets are balanced (well-formed).

For example, given the string "([])[](\{\})", you should return true.

Given the string "([)]" or "((()", you should return false.

---

I got this from the old Daily Coding Problem
email list as "Daily Coding Problem: Problem #712 [Easy]",
and I'm certain those folks sent it out several times
while they were a going concern.

When I [initially tried](https://github.com/bediger4000/balanced-parens)
this problem, I thought it was a coding problem that had been around for a while.
I thought it was one of the coding problems that has a clever and a non-clever
solution, and you had to have seen the clever solution before,
because it isn't obvious.

It turns out that deciding whether strings of parentheses, brackets, braces, etc etc
are balanced or not is the problem of recognizing [Dyck Languages](https://en.wikipedia.org/wiki/Dyck_language).
It's a well-studied problem, and has applications in a lot of fields.

---

Four Programs doing Dyck word recognition differently:

`counter.go` counts characters, +1 for an opening parenthesis, -1 for a closing.
If the counter becomes negative, or is non-zero at the end of the word,
the parentheses do not balance.

`reduction.go` - excises matching pairs of opening and closing parens/braces/brackets
until it can't.
A zero-length resulting string means the initial string was a Dyck word.

`balanced2.go` - the standard push opening characters on a stack,
if you find a closing character, pop the stack, if that popped character
isn't the matching opening character, the string isn't balanced.

`depth.go` - find maximum nesting of opening and closing characters
regardless of match.
For each nesting depth, finding opening character sets the appropriate
matching closing character.
The next character at that nesting depth must be the matching closing character,
or else the string is unbalanced.

For the string `((([{}{}{()}))))`, the nesting depth works out like this:

```
1  (              )
2   (            ) 
3    (          )
4     [        )
5      {}{}{  }    
6           ()
```

Depth 4, the '[' doesn't match the ')'.
This is an unbalanced string.
