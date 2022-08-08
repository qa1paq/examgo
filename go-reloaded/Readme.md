# go-reloaded

### Description
In this project you will use some of your old functions made in your old repository. You will use them with the objective of making a simple text completion/editing/auto-correction tool.
One more detail. This time the project will be corrected by auditors. The auditors will be other students and you will be an auditor as well.
We advise you to create your own tests for yourself and for when you will correct your auditees.


### Usage
```console
$ cat sample.txt
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.

$ go run . sample.txt result.txt

$ cat result.txt
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.

$ cat sample.txt
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.

$ go run . sample.txt result.txt

$ cat result.txt
Simply add 66 and 2 and you will see the result is 68.

$ cat sample.txt
There is no greater agony than bearing a untold story inside you.

$ go run . sample.txt result.txt

$ cat result.txt
There is no greater agony than bearing an untold story inside you.

$ cat sample.txt
Punctuation tests are ... kinda boring ,don't you think !?

$ go run . sample.txt result.txt

$ cat result.txt
Punctuation tests are... kinda boring, don't you think!?
```