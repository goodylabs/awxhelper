# Installation

```bash
# Prerequisites
line="export PATH=\"\$HOME/.${APP}/bin:\$PATH\""
rc_file="${HOME}/.$(basename "$SHELL")rc"
echo $line >> $rc_file
source $rc_file

# Download and install binary
curl -s https://raw.githubusercontent.com/goodylabs/awxhelper/refs/heads/main/scripts/download.sh | bash -s
```
