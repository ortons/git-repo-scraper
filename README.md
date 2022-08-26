# git-repo-scraper

### Usage
This will create a shell script you can copy / paste straight into the target machine which will create the same
directory structure and clone repos underneath.

`git-repo-scaper -o write ~/code`

output redirection to a file works:
`./git-repo-scraper ~/code -o create -v > output.txt`

Input flags: 
* Verbosity `-v, -vv, -vvv`
* Operation `-o [read | create]`
* Root folder: positional & required

### Improvements
* improve the logging
* read the file
* options to just have urls, not the entire shell creation one-liner. 



---
### Note:
You can do some of this in a find one-liner, and some sed magic would help too. Probably should compare and profile the two! 

```find . -path '*/.git/config' -execdir git remote get-url origin \;```
