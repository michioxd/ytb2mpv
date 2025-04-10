@echo off

set GOOS=windows

if exist dist rmdir /s /q dist

echo == [BUILD] Building amd64 ============================================
set GOARCH=amd64
rsrc -manifest=app.manifest -ico ../media/ytb2mpv.ico -o ytb2mpv.syso
go build -o dist/ytb2mpv.exe -ldflags "-s -w -H windowsgui" -trimpath -v -a -gcflags "all=-trimpath=$GOPATH" .
del ytb2mpv.syso

echo Done i guess :)