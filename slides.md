---
author: I Like Bad Code
date: " by Tom"
paging: "%d / %d"
---

```
~~~figlet -f banner
I  Like  Bad  Code
~~~
~~~figlet -f bell
by Tom
~~~
```

---

Consider a `desk-printer` program that outputs the following:

```
~~~go run .

~~~

```

It has a title, a legend, a main section and a footer.

---

A less experienced developer might structure their code like this:

```go
func printBasic(desks [][]string) {
    // title
    fmt.Printf("\n~DESK PRINTER~\n\n")
    // legend
    fmt.Println(available + " = Available Desk")
    fmt.Println(normal + " = Normal Person")
    fmt.Println(annoying + " = Annoying Person")
    fmt.Println(sick + " = Sick Person")
    fmt.Println("")
    // desks
    for _, row := range desks {
        rowStr := ""
        for _, desk := range row {
            rowStr += desk + " "
        }
        fmt.Println(rowStr)
    }
    // footer
    fmt.Println("\n© copyright tomontheinternet")
}

```

Everything happens in a single function.

---

As a more experienced developer, I would encourage them to break their function up into smaller functions.

```go
func printFactored(desks [][]string) {
    printTitle()
    printLegend()
    printDesks(desks)
    printFooter()
}
```

Here's why this is better.

- Reads like English.
- Clear separation.
- Self documenting.
- Clear what depends on what.
- Easy to test individual pieces.

---

```
~~~figlet -f banner
Sorry, but...
~~~
```

I like `printBasic` more than `printFactored`, even though `printFactored` is what I've been recommending for years.

---

Here's why I like this better.

- Everything is right in front of me.
- Comments can be more expressive than function names.
- Easier to refactor.

```go
func printBasic(desks [][]string) {
    // title
    fmt.Printf("\n~DESK PRINTER~\n\n")
    // legend
    fmt.Println(available + " = Available Desk")
    fmt.Println(normal + " = Normal Person")
    fmt.Println(annoying + " = Annoying Person")
    fmt.Println(sick + " = Sick Person")
    fmt.Println("")
    // desks
    for _, row := range desks {
        rowStr := ""
        for _, desk := range row {
            rowStr += desk + " "
        }
        fmt.Println(rowStr)
    }
    // footer
    fmt.Println("\n© copyright tomontheinternet")
}
```

---

```
~~~figlet -f banner
Who cares, Tom?
~~~
```

_I care._ It's strange realizing the patterns I liked can actually be anti-patterns. I find it hard to let go of positions I've defended. But it's also scary to think of becoming someone who can't change. Perhaps I'm not alone?

_Good code tells a story._ I used to think good code should be nice to look at. I don't think that anymore. Sometimes, making your code nice to read makes it hard to understand. Good code is easy to understand, even if it's ugly and needs comments.

---

```
~~~figlet -f banner
So, longer functions?
~~~
```

Yes. Sometimes. But sometimes, no.

If a longer function will make the code easier to understand, then yes.

If a longer function makes the code harder to understand, then... no.

It's hard to reevaluate my instincts. But if I don't, I'll never improve as a programmer.
