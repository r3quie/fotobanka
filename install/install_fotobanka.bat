REM Copyright © 2024 Pavel Kadera info@pavelkadera.cz

REM Licensed under the Apache License, Version 2.0 (the "License");
REM you may not use this file except in compliance with the License.
REM You may obtain a copy of the License at

REM 	http://www.apache.org/licenses/LICENSE-2.0

REM Unless required by applicable law or agreed to in writing, software
REM distributed under the License is distributed on an "AS IS" BASIS,
REM WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
REM See the License for the specific language governing permissions and
REM limitations under the License.

@echo off
setlocal

REM Get the directory of the current script
set SCRIPT_DIR=%~dp0

REM Add the script directory to PATH
setx PATH "%PATH%;%SCRIPT_DIR%"

REM Open a new terminal window with the message
start cmd /k "echo Pro spusteni programu zadejte 'fotobanka banka' pro zobrazeni licence zadejte 'fotobanka licence'"

endlocal