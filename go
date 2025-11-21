@echo off
setlocal enabledelayedexpansion

:: Define input path
set "inputPath=D:\JAGSP-LIVE\JAGSP-P-SCAN\3X9"

:: Initialize counter for BackendReady values
set /a "counter=1"
set "maxCounter=5"

:: Flag to track if any files were processed
set "filesProcessed=false"

:: Loop through all ISBN folders
for /d %%I in ("%inputPath%\*") do (
    if exist "%%I\TIF" (
        set "stsFile=%%I\TIF\%%~nxI.sts"
        
        if exist "!stsFile!" (

            :: Read sts content (trim spaces)
            for /f "usebackq delims=" %%A in ("!stsFile!") do (
                set "fileContent=%%A"
            )
            :: Trim leading/trailing spaces
            for /f "tokens=* delims= " %%B in ("!fileContent!") do set "fileContent=%%B"

            :: Convert to lowercase for matching
            set "lc=!fileContent:~0!"
            set "lc=!lc: =!"
            set "lc=!lc:,=!"
            set "lc=!lc:~0!"

            :: ---------------------------------------------------------
            :: PROCESS Scanning Completed OR LZWCompleted → BackendReadyX
            :: ---------------------------------------------------------
            if /i "!fileContent!"=="Scanning Completed" (
                set "process=true"
            ) else if /i "!fileContent!"=="LZWCompleted" (
                set "process=true"
            ) else (
                set "process=false"
            )

            if "!process!"=="true" (
                set "newContent=BackendReady!counter!"

                powershell -NoProfile -Command "Set-Content -Path '!stsFile!' -Value '!newContent!' -NoNewline"

                echo Updated !stsFile! from "!fileContent!" to "!newContent!"

                set /a "counter+=1"
                if !counter! gtr !maxCounter! set "counter=1"

                set "filesProcessed=true"
            )
        )
    )
)

if not !filesProcessed! (
    echo No files were processed.
)

echo Done processing all files.
pause
