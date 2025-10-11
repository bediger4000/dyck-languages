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

https://stackoverflow.com/questions/2509358/how-to-find-validity-of-a-string-of-parentheses-curly-brackets-and-square-brack
