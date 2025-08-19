# Installation

```bash
line="export PATH=\"\$HOME/.${APP}/bin:\$PATH\""
rc_file="${HOME}/.$(basename "$SHELL")rc"
echo $line >> $rc_file
source $rc_file

```
