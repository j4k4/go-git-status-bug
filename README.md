The application in this repository will create a file in an ignored folder which has a .gitignore, too, in which the 
created file is unignored. go-git will declare the working tree as not clean while on the cli, git will tell us that it
is clean:

```
➜  git-status-bug git:(main) git version
git version 2.43.0

➜  git-status-bug git:(main) go run .          
working tree is clean: true
creating file "this-should-be-ignored/ignored-too.md" which should be ignored
working tree is clean: false

➜  git-status-bug git:(main) git status        
On branch main
nothing to commit, working tree clean
```