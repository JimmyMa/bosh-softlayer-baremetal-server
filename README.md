bosh-softlayer-baremetal-server
===============================

A server to create, provision and manage SoftLayer baremetals servers.
## Setup
The server can be deployed by BOSH with a release, [bosh-softlayer-baremetal-server-release][1]. The release contains three jobs named `baremetal-provision-server`, `xcat-server` and `redis`. Job `baremetal-provision-server` and `xcat-server` should always be on the same node.

This is an example [deployment.yml][2] file for your reference. 


## Usage
Once deployment is completed, you can login the server and take a check. 

Run `monit summary`,  normally if everything works well you can see an output like this, 
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
If any job is not running,  run `monit restart <job-name>` to restart it. If this doesn't work out, you can check logs under `/var/vcap/sys/log` and do further investigation. 

The server exposes RESTful api which you can find details in [api.md][3]
 and you can use a CLI written in GO to leverage it. Here is the project [link][4].


To get the lasted executable `bmp`, you can run following `go` command (Assuming that you have `go1.42` or later version installed.) 

>$ go get github.com/cloudfoundry-community/bosh-softlayer-tools/bmp

You can find an executable `bmp` in your `$GOPATH/bin` directory and move it to `/user/local/bin` on linux or Mac OS X

To use bmp CLI, firstly you need to set the target and login, then you can use the CLI to communicate with server to create and provision a baremetal.  
>**Example:**
>$ bmp target -t http://10.12.20.17:8080
>
>$ bmp login -u admin -p admin
>Login Successful!
>
>$ bmp status
>BMP server info
 name:    Bluemix Provision Server
 version: 0.1

When to create a baremetal on Softlayer, you need to create a deployment.yml firstly and specify your required configurations like `cores` and `memory`. `size` represent the number of baremetals . Please find an example yml file [here][5].

>**Example:**
>$ bmp create-baremetals -d deployment.yml

While creating baremetals, you can check the task log as below. 
>**Example:**
>$ bmp task --task_id=12

After deployment completes, you can check the baremetal status as below. Normally the status for a newly created baremetal is `new`

>**Example:**
>$ bmp bms -d deployment.yml

[1]: https://github.com/cloudfoundry-community/bosh-softlayer-baremetal-server-release
[2]: https://github.com/cloudfoundry-community/bosh-softlayer-baremetal-server-release/docs/baremetal-provision-server.yml
[3]: https://github.com/cloudfoundry-community/bosh-softlayer-baremetal-server/blob/master/docs/api.md
[4]: https://github.com/cloudfoundry-community/bosh-softlayer-tools
[5]: https://github.com/cloudfoundry-community/bosh-softlayer-baremetal-server/blob/master/docs/deployment.yml
