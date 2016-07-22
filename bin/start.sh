#!/bin/bash

mkdir -p logs

if [ -f $TMPDIR/bosh-softlayer-bmp.pid ]; then
	echo "SoftLayer Baremetal server already running, stop it first: bin/stop.sh"
	exit 0
fi

echo "Starting SoftLayer Baremetal Server"

nohup bin/api_ctl 1 >> logs/api_ctl.stdout 2 >> logs/api_ctl.stderr &

mktemp -q $TMPDIR/bosh-softlayer-bmp.pid
echo `ps -A | grep "ruby bin/api_ctl" | head -1 | cut -c 1-5` >> $TMPDIR/bosh-softlayer-bmp.pid

nohup bin/worker_ctl 1 >> logs/worker_1.stdout 2 >> logs/worker_1.stderr &
nohup bin/worker_ctl 1 >> logs/worker_2.stdout 2 >> logs/worker_2.stderr &
nohup bin/worker_ctl 1 >> logs/worker_3.stdout 2 >> logs/worker_3.stderr &
nohup bin/worker_ctl 1 >> logs/worker_4.stdout 2 >> logs/worker_4.stderr &
nohup bin/worker_ctl 1 >> logs/worker_5.stdout 2 >> logs/worker_5.stderr &
nohup bin/worker_ctl 1 >> logs/worker_6.stdout 2 >> logs/worker_6.stderr &
nohup bin/worker_ctl 1 >> logs/worker_7.stdout 2 >> logs/worker_7.stderr &
nohup bin/worker_ctl 1 >> logs/worker_8.stdout 2 >> logs/worker_8.stderr &
nohup bin/worker_ctl 1 >> logs/worker_9.stdout 2 >> logs/worker_9.stderr &
