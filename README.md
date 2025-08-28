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
```

# Configuration

```bash
awxhelper configure
```

# Available Commands

```bash
configure     Configure your AWX connection
downloaddb    Download database dump
forceupdate   Force check for new updates and install if available
```
