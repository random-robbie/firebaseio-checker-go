# firebaseio-checker-go

Firebase url checker in go it will dump the contents of the database to a json file for you.

Requirements
---

```
go get github.com/logrusorgru/aurora
go get github.com/remeh/sizedwaitgroup
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

[![Capture.png](https://i.postimg.cc/52ngnsqp/Capture.png)](https://postimg.cc/75JSLMRT)


it will output everything to the cfg folder with the database name .json and it's contents for easy grepping.

Use a VPS from DO

[![DigitalOcean Referral Badge](https://web-platforms.sfo2.cdn.digitaloceanspaces.com/WWW/Badge%201.svg)](https://www.digitalocean.com/?refcode=e22bbff5f6f1&utm_campaign=Referral_Invite&utm_medium=Referral_Program&utm_source=badge)
