bosh-softlayer-baremetal-server
===============================

A server to create, provision and manage SoftLayer baremetals servers.

## Setup
The server could be deployed by BOSH with a release as below. The release contains three jobs named baremetal-provision-server, xcat-server and redis. Job baremetal-provision-server and xcat-server should always be on the same node.
https://github.com/cloudfoundry-community/bosh-softlayer-baremetal-server-releas

This is an example deployment yaml file.
https://github.com/cloudfoundry-community/bosh-softlayer-baremetal-server-release/docs/baremetal-provision-server.yml

## Usage
Once deployment is completed, you could login the server and take check. 

Run `monit summary`,  normally you could see output like this,
```
~# monit summary
The Monit daemon 5.2.5 uptime: 15d 9h 41m

Process 'xcat-server'               running
Process 'redis'                     running
Process 'baremetal-provision-server' running
Process 'baremetal-provision-worker-1' running
Process 'baremetal-provision-worker-2' running
Process 'baremetal-provision-worker-3' running
Process 'baremetal-provision-worker-4' running
Process 'baremetal-provision-worker-5' running
Process 'baremetal-provision-worker-6' running
System 'system_localhost'           running
```
If any job is down, you could go to `/var/vcap/sys/log` and take a look at the logs.

The server exposes RESTful api which you could find details in `api.md` and you could use a CLI written in GO to leverage it. Firstly you need to set the target and login, then you could use the server to create and provision a baremetal.  Here is the link of CLI.
https://github.com/cloudfoundry-community/bosh-softlayer-tools

>**Examples:**
>$ bmp target -t http://{server_ip}:8080
>
>$ bmp login -u admin -p admin
>{"status":200,"data":null}
>
>$ bmp status
>BMP server info
 name:    Bluemix Provision Server
 version: 0.1

When to create a baremetal on Softlayer, you need to create a `deployment.yml` firstly and specify the your required configuration like core or mem.

>**Example:**
>$ bmp create-baremetals -d deployment.yml

After deployment completes, you could check the baremetal status as below.

>**Example:**
>$ bmp bms -d deployment.yml