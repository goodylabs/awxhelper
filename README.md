# Installation

```bash
# Prerequisites
APP="awxhelper"
rc_file="${HOME}/.$(basename "$SHELL")rc"

line="export PATH=\"\$HOME/.${APP}/bin:\$PATH\""
echo $line >> $rc_file

source $rc_file

# Download and install binary
curl -s "https://raw.githubusercontent.com/goodylabs/${APP}/refs/heads/main/scripts/download_scripts.sh" | bash -s

# Check AWXHelper version
awxhelper --version
```

# Configuration

```bash
awxhelper configure
```

# Available Commands

```bash
Usage:
  awxhelper [flags]
  awxhelper [command]

Available Commands:
  configure   Configure connection to your AWX (url, username, password)
  downloaddb  Download database backup to your local /tmp dir
  help        Help about any command
  runbackup   Run job that makes a database backup
  runjob      Run any job you want!

Flags:
  -x, --debug     Run in debug mode
  -h, --help      help for awxhelper
  -v, --version   Print version and exit
```
