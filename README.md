# cleanmyfolder
## Description

These are my first steps to learn golang.

## Use

This Clean older files (Tested in linux),  You need pass folder and days to set the antique.

Use: 
```bash
go run cmd/main.go -folder=${HOME}/Downloads -days=15
```

## TODO

- [ ] Remove feak delete.
- [ ] Check if folder exist.
- [ ] CleanUP empty folders.
- [ ] Boolean parameter to implicit confimation
    - [ ] If bool is false, ask by terminal: Do you want delete these files?
    - [ ] If is yes, delete list.