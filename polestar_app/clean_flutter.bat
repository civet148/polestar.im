@echo off
taskkill /F /IM dart.exe
del /q %FLUTTER_SDK%\bin\cache\lockfile
pause
