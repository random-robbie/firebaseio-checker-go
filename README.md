# firebaseio-checker-go

Firebase url checker in go it will dump the contents of the database to a json file for you.

Requirements
---

```
go get github.com/logrusorgru/aurora
```

Build
---

```
go build
```

How to
---

Put all your firebase urls in a file called list.txt with out https://

```
./firebaseio-checker
```

it will output everything to the cfg folder with the database name .json and it's contents for easy grepping.
