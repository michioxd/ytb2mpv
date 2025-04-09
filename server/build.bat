@echo off

set GOOS=windows

if exist dist rmdir /s /q dist

echo == [BUILD] Building amd64 ============================================
set GOARCH=amd64
go build -o dist/ytb2mpv-amd64.exe -ldflags "-s -w -H windowsgui" -trimpath -v -a -gcflags "all=-trimpath=$GOPATH" .

echo == [BUILD] Building x86 ============================================
set GOARCH=386
go build -o dist/ytb2mpv-i386.exe -ldflags "-s -w -H windowsgui" -trimpath -v -a -gcflags "all=-trimpath=$GOPATH" .

echo Done i guess :)