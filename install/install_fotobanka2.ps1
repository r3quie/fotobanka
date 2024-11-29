# Copyright © 2024 Pavel Kadera info@pavelkadera.cz

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

# 	http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Get the current working directory
$newPath = Get-Location

# Get the current user PATH
$userPath = [Environment]::GetEnvironmentVariable("PATH", "User")

# Check if the new path is already in the user's PATH
if (-not $userPath -or $userPath -notcontains $newPath) {
    # Append the new path
    $updatedPath = if ($userPath) { "$userPath;$newPath" } else { "$newPath" }
    [Environment]::SetEnvironmentVariable("PATH", $updatedPath, "User")
    Write-Host "PATH updated successfully!"
} else {
    Write-Host "The directory is already in the PATH."
}

# Launch Command Prompt with the specified message
Start-Process cmd -ArgumentList "/k echo Pro spusteni programu zadejte 'fotobanka banka' pro zobrazeni licence zadejte 'fotobanka licence'"
