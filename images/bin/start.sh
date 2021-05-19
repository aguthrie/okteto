#!/bin/sh

log() {
        echo "$(date +%Y-%m-%dT%H:%M:%S)" "$1"
}
set -e

userID="$(id -u)"
echo "USER:$userID"
log "development container starting"
if [ -d "/var/okteto/cloudbin" ]; then
        if [ -w "/usr/local/bin" ]; then
                cp /var/okteto/cloudbin/* /usr/local/bin
        fi
fi

if [ -d "/var/okteto/hooks/start.sh" ]; then
  /var/okteto/hooks/start.sh
done

remote=""
reset=""
verbose="--verbose=false"
while getopts ":s:rev" opt; do
        case $opt in
        e)
                reset="--reset"
                ;;
        r)
                remote="--remote"
                ;;
        v)
                verbose="--verbose"
                ;;
        s)
                sourceFILE="$(echo "$OPTARG" | cut -d':' -f1)"
                destFILE="$(echo "$OPTARG" | cut -d':' -f2)"
                dirName="$(dirname "$destFILE")"

                if [ ! -d "$dirName" ]; then
                        mkdir -p "$dirName"
                fi

                log "Copying secret $sourceFILE to $destFILE"
                if [ "/var/okteto/secret/$sourceFILE" != "$destFILE" ]; then
                        cp "/var/okteto/secret/$sourceFILE" "$destFILE"
                fi
                ;;
        \?)
                log "Invalid option: -$OPTARG" >&2
                exit 1
                ;;
        esac
done

syncthingHome=/var/syncthing
log "Copying configuration files to $syncthingHome"
cp /var/syncthing/secret/* $syncthingHome
chmod 644 $syncthingHome/cert.pem $syncthingHome/config.xml $syncthingHome/key.pem

log "Executing okteto-supervisor $remote $reset $verbose"
exec /var/okteto/bin/okteto-supervisor $remote $reset $verbose
