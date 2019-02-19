[![Build Status](https://travis-ci.org/tzmfreedom/land.svg?branch=master)](https://travis-ci.org/tzmfreedom/land)

# Land

Salesforce Apex Execution Environment on Local System.

## Installation

For Golang User
```bash
$ go get -u github.com/tzmfreedom/land
```

For Linux
```bash
$ curl -sL http://install.freedom-man.com/land.sh | bash
```

For Windows
```bash
@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile ^
  -InputFormat None -ExecutionPolicy Bypass ^
  -Command "iex ((New-Object System.Net.WebClient).DownloadString('http://install.freedom-man.com/land.ps1'))" ^
  && SET "PATH=%PATH%;%APPDATA%\land\bin"
```

For Powershell
```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; iex ((New-Object System.Net.WebClient).DownloadString('http://install.freedom-man.com/land.ps1'))
```

## Usage

```bash
$ land run -f {file} -a "ClassName#MethodName"
```

```bash
$ land run -d {directory} -a "ClassName#MethodName"
```

## Contribute

Just send pull request if needed or fill an issue!

## License

The MIT License See [LICENSE](https://github.com/tzmfreedom/land/blob/master/LICENSE) file.
