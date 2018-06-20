# DiceRoll-Go
This is a simple dice rolling application I wrote to test my skills in the Go programming langauge.

It can be used either staticly, or as an interactive shell.

To use is staticly, just call it in the command line and put the number of dice and type of dice as an argument:

```
> diceroll-go 3d8
7, 4, 8
Total: 19
```

If you call it without any arguments, it will go into an interactive shell:

```
> diceroll-go
Welcome to DiceRoll-Go.
You can roll dice by typing 'roll #d#'.
>>> roll 3d8
7, 4, 8
Total: 19
```