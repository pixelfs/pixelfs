#!/bin/sh

set -e
. /etc/os-release

PIXELFS_HOME="$HOME/.pixelfs"

create_home_dir() {
	mkdir -p "$PIXELFS_HOME" || {
        echo "Error: Failed to create $PIXELFS_HOME. Please check permissions." >&2
        exit 1
    }
}

summary() {
	echo "----------------------------------------------------------------------"
	echo " pixelfs package has been successfully installed."
	echo ""
	echo " Please follow the next steps to start the software:"
	echo ""
	echo "    systemctl --user start pixelfs"
	echo "    systemctl --user enable pixelfs"
	echo "    loginctl enable-linger"
	echo ""
    echo " Configuration settings can be adjusted here:"
    echo "    ~/.pixelfs/config.toml"
    echo ""
	echo "----------------------------------------------------------------------"
}

#
# Main body of the script
#
{
	create_home_dir
	summary
}
