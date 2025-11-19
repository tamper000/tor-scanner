#!/bin/sh
set -e

echo "Running tor-scanner with user parameters: $@"
/usr/local/bin/tor-scanner "$@" -output /etc/tor/tor.conf

if [ ! -f /etc/tor/tor.conf ]; then
    echo "Failed to generate tor.conf."
    exit 1
fi

echo -e "\nAvoidDiskWrites 1\nSOCKSPort 0.0.0.0:9050 IsolateDestAddr\nHardwareAccel 1\nReducedConnectionPadding 1" >> /etc/tor/tor.conf

echo "Starting Tor with configuration: /etc/tor/tor.conf"
exec tor -f /etc/tor/tor.conf
