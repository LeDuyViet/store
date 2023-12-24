#!/bin/sh
git config core.autocrlf false

# Don’t forget the dot at the end
git rm --cached -r .      

git reset --hard
