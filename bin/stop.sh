#!/bin/bash

if [ ! -f $TMPDIR/bosh-softlayer-bmp.pid ]; then
	echo "SoftLayer Baremetal server not running, run it first: bin/start.sh"
	exit 0
fi

echo "Stopping SoftLayer Baremetal Server"

BMP_PID=`cat $TMPDIR/bosh-softlayer-bmp.pid`

kill -9 $BMP_PID

rm $TMPDIR/bosh-softlayer-bmp.pid