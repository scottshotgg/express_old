#!/bin/sh

# os.tmpdir from node.js
for OS_TMPDIR in "$TMPDIR" "$TMP" "$TEMP" /tmp
do
  test -n "$OS_TMPDIR" && break
done

# kill any currently running Discord
if pgrep Discord ; then
  pkill Discord
  sleep 1
  pkill -9 Discord
fi

# This is probably just paranoia, but some people claim that clearing out
# cache and/or the sock file fixes bugs for them, so here we go
for DIR in /home/* ; do
  rm -rf "$DIR/.config/discordstable/Cache"
  rm -rf "$DIR/.config/discordstable/GPUCache"
done
rm -f "$OS_TMPDIR/discordstable.sock"
