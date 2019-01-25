$name = "land"
$bin_name = "land"
$version = "0.1.2"
$github_user = "tzmfreedom"
If ($Env:PROCESSOR_ARCHITECTURE -match "64") {
  $arch = "amd64"
} Else {
  $arch = "386"
}
$archive_file = "${name}-${version}-windows-${arch}.zip"
$url = "https://github.com/${github_user}/${name}/releases/download/v${version}/${archive_file}"
$dest_dir = "$Env:APPDATA\${name}"

[System.Net.ServicePointManager]::SecurityProtocol = [System.Net.SecurityProtocolType]::Tls12;

wget $url -OutFile "${dest_dir}\tmp.zip"
Expand-Archive -Path "${dest_dir}\tmp.zip" -DestinationPath "${dest_dir}\tmp" -Force
if (!(Test-Path -path "${dest_dir}\bin")) {  New-Item "${dest_dir}\bin" -Type Directory }
mv "${dest_dir}\tmp\windows-${arch}\${bin_name}" "${dest_dir}\bin\${bin_name}.exe" -Force
rm "${dest_dir}\tmp.zip" -Force
rm "${dest_dir}\tmp" -Force -Recurse

# ensure added to path in registry
$bin_path = "${dest_dir}\bin"
$reg_path = [Environment]::GetEnvironmentVariable("PATH", "User")
If ($reg_path.Contains($bin_path) -eq $false) {
  [Environment]::SetEnvironmentVariable("PATH", "${reg_path};${bin_path}", "User")
}
# ensure added to path for current session
If ($($Env:Path).Contains($bin_path) -eq $false) {
  $Env:Path += ";${bin_path}"
}
Write-Host "${name} installed. Run '${name} help' to try it out."
