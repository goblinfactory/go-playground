# random thoughts (possibly not correct)

-   It looks like basic GO lang objects are thinner, smaller, simpler.

    -   for example; `var ct = context.Background()` when set a breakpoint and inspect
        ct, it has only 1 property, pointer to data. I dont get the same impression with C# base objects. This may be because of C# OO nature. To be confirmed.

-   I love how GO right from "get-go" throws you into a world where you can easily "block" a thread, until something cool happens using select and channels. This is ... very liberating.

-   would be super if I could program in VS code and have debug console output be sent
    to a small external display, e.g. mobile app or tablet. Or simply undock the
    terminal window (floating) like chrome debugger window.
    So that I can keep coding in full screen when working on laptop with minimal screen res.
    I'm constanting expanding and resizing the debug console so that i can see the output.

-   I'm amazed at how you can casually throw in a check against a channel in any select statement.

-   mmm, organisations structure tend to reflect the culture?? or people?
    -   so what I need to do is ask, how do i find people that think like me who are
        in control of companies or consultancies?

## GO observations

-   func return tuples can't have names describing what the return value is. My guess is this is to encourage you to create a struct for it.
-   generics causes a developer to focus too much on writing re-usable code, instead of just getting the job done. By not re-using, there's no coupling, and no exostential burden of worrying about very complex package management. The lack of generics is a godsend. (possibly/very likely!)

# nice to check

-   see if it's possible to detect if a function is being called static by 2 goroutines, instead of being safely synchronised using a select and channels?

# surprised that go doesnt

-   Go does not run defer for the main thread when a user presses control+c inside a console application in Go. I would have thought this a great feature for the language to automatically support. Unfortunately not, and you have some non-intuitive hoops to jump through to make a dabase or console app cleanup up properly when encounting control+c see : https://medium.com/@matryer/make-ctrl-c-cancel-the-context-context-bd006a8ad6ff

# nice to check

-   I'm amazed at how you can casually throw in a check against a channel in any select statement.
