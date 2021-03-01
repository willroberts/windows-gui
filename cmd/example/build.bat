@echo off

REM Clean up old build artifacts.
del gui.exe

REM Compile the manifest.
rsrc -manifest main.manifest -o rsrc.syso

REM Build the server for Windows.
go build -o gui.exe -ldflags="-H windowsgui"