# sandbox currency and Money types

Messy experiment playing with types and interfaces, and experiment with currency (Money) types

## useful links

-   https://floating-point-gui.de/formats/integer/ : see the note and warning NOT to use integers if at all possible.
-   https://floating-point-gui.de/xkcd/

#### Creating currency objects

```go
    zar1 := money.NewZAR(18325.01)
	zar2 := money.NewZAR(18325.99)
	gbp1 := money.NewGBP(50.95))
```

#### Comparing

`.Greater()`, `.GreaterThan()`, `.Equal()`, `.LessThan()` etc

```go
    if zar1.Greater(zar2) {
        ...
    }
```

#### Printing

```go
    fmt.Println("|", gbp1.Wide(), "|")
    fmt.Println("|", gbp1.Narrow(), "|")
```

**_produces_**

14 and 7 digit wide formats.

```dos
| £          50.95 |
| £   50.95 |
```

#### Cloning with new Value

Clone the currency type and formatting;

```go
    gbp2:= gbp1.Clone(123.45)
```

#### type safety

You can't add or substract different currency types. Compiler will prevent that.
