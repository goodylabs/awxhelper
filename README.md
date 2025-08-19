# Installation

```bash
# Prerequisites
APP="awxhelper"
rc_file="${HOME}/.$(basename "$SHELL")rc"

line="export PATH=\"\$HOME/.${APP}/bin:\$PATH\""
echo $line >> $rc_file

completion_file="${HOME}/.$(basename "$SHELL")rc_${APP}"
touch $completion_file
echo "source \"$completion_file\"" >> $rc_file

source $rc_file

# Download and install binary
curl -s "https://raw.githubusercontent.com/goodylabs/${APP}/refs/heads/main/scripts/download_script.sh" | bash -s
```
