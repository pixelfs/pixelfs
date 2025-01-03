#!/bin/sh
. /etc/os-release

echo "Checking and stopping pixelfs service for all users"

awk -F: '{if ($3 >= 0 && $3 < 65534 && $1 != "nobody" && $7 != "/usr/sbin/nologin" && $7 != "/bin/sync") print $1}' /etc/passwd | while read -r user; do

    su - "$user" -c '
        if command -V systemctl >/dev/null 2>&1; then
            if systemctl --user is-active pixelfs >/dev/null 2>&1; then
                echo "Stopping pixelfs service for $USER"
                systemctl --user stop pixelfs >/dev/null 2>&1 || true
                systemctl --user disable pixelfs >/dev/null 2>&1 || true
                systemctl --user daemon-reload || true
            else
                echo "PixelFS service is not running for $USER"
            fi
        fi
    ' || echo "Failed to process $user"

    user_home=$(eval echo "~$user")
    if [ -d "$user_home/.pixelfs" ]; then
        echo "Removing cache directory for $user: $user_home/.pixelfs/cache"
        rm -rf "$user_home/.pixelfs/cache"
    fi
done

echo "All users processed successfully"
