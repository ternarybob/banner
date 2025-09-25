# Version management script for local development (PowerShell)
param(
    [Parameter(Position=0)]
    [ValidateSet("current", "get", "patch", "minor", "major")]
    [string]$Action = "help"
)

$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$RootDir = Split-Path -Parent $ScriptDir
$VersionFile = Join-Path $RootDir "VERSION"

function Get-CurrentVersion {
    if (Test-Path $VersionFile) {
        Get-Content $VersionFile -Raw | ForEach-Object { $_.Trim() }
    } else {
        "v0.0.0"
    }
}

function Increment-PatchVersion {
    $current = Get-CurrentVersion
    $parts = $current -replace '^v', '' -split '\.'
    $major = [int]$parts[0]
    $minor = [int]$parts[1]
    $patch = [int]$parts[2] + 1

    $newVersion = "v$major.$minor.$patch"
    Set-Content -Path $VersionFile -Value $newVersion -NoNewline
    return $newVersion
}

function Increment-MinorVersion {
    $current = Get-CurrentVersion
    $parts = $current -replace '^v', '' -split '\.'
    $major = [int]$parts[0]
    $minor = [int]$parts[1] + 1

    $newVersion = "v$major.$minor.0"
    Set-Content -Path $VersionFile -Value $newVersion -NoNewline
    return $newVersion
}

function Increment-MajorVersion {
    $current = Get-CurrentVersion
    $parts = $current -replace '^v', '' -split '\.'
    $major = [int]$parts[0] + 1

    $newVersion = "v$major.0.0"
    Set-Content -Path $VersionFile -Value $newVersion -NoNewline
    return $newVersion
}

switch ($Action) {
    "current" { Get-CurrentVersion }
    "get" { Get-CurrentVersion }
    "patch" { Increment-PatchVersion }
    "minor" { Increment-MinorVersion }
    "major" { Increment-MajorVersion }
    default {
        Write-Host "Usage: .\version.ps1 {current|patch|minor|major}"
        Write-Host ""
        Write-Host "Commands:"
        Write-Host "  current  - Show current version"
        Write-Host "  patch    - Increment patch version (x.x.X)"
        Write-Host "  minor    - Increment minor version (x.X.0)"
        Write-Host "  major    - Increment major version (X.0.0)"
        Write-Host ""
        Write-Host "Current version: $(Get-CurrentVersion)"
    }
}