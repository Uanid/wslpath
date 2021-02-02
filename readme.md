# WSLPATH

### How to use
```shell script
wslpath C:\Users\musong\Documents

-> /mnt/c/Users/musong/Documents
```

### Man Page
```shell script
WSLPATH https://github.com/uanid/wslpath v1.0
Usage: wslpath C:\\Users\\musong
return: /mnt/c/Users/musong
```

### How to Download Pre-Compiled Binary
Go to [Release Page](https://github.com/Uanid/wslpath/releases)

### How to Build
```shell script
go build -ldflags="-s -w" -o wslpath.exe main.go

or

go build -o wslpath.exe main.go
```

