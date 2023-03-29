# interactiveShell
A fully interactive TCP cmd reverse shell for Windows OS. 
Automatically tries to reconnect to the listener if the connection is lost

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
