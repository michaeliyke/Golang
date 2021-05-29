@ECHO off
SETLOCAL
CLS
ECHO. 
ECHO. Executing ^go run %* & ECHO.
go run %* & ECHO. & ECHO. Done!
ENDLOCAL