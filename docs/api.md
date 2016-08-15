Baremetal Provision Server API
===============================

This is a brief introduction of Baremetal Provision Server API.

#Login bmp server
`GET`  /login/{username}/{password}

**Request**
`GET /login/{username}/{password}`

**cURL**
```
curl -X GET "http://10.12.20.17:8080/login/{username}/{password}
```
**Response**
Status
`200`

#Check bmp server status
`GET`    /info

**Request**
`GET /info`

**cURL**
```
curl -X GET "http://10.12.20.17:8080/info
```
**Response**
Status
`200`

Body
```
{
    "status": 200,
    "data": {
        "name": "Bluemix Provision Server",
        "version": "0.1"
    }
}
```

#List bmp server stemcells
`GET`  /stemcells

**Request**
`GET /stemcells`

**cURL**
```
curl -X GET "http://10.12.20.17:8080/stemcells
```
**Response**
Status
`200`

Body
```
{
    "status": 200,
    "data": [
        "bosh-stemcell-0.1-softlayer-baremetal",
        "bosh-stemcell-0.3-softlayer-baremetal"
    ]
}
```

#List Softlayer packages
`GET`  /sl/packages

**Request**
`GET /sl/packages`

**cURL**
```
curl -X GET "http://10.12.20.17:8080/sl/packages
```
**Response**
Status
`200`

Body
```
{
    "status": 200,
    "data": {
        "packages": [
            {
                "id": 56,
                "name": "Quad Processor Multi Core Nehalem EX"
            },
            {
                "id": 126,
                "name": "Single Xeon 1200 Series (Sandy Bridge / Haswell)"
            },
            {
                "id": 142,
                "name": "Single Xeon 2000 Series (Sandy Bridge)"
            },
            {
                "id": 143,
                "name": "Dual Xeon 2000 Series (Sandy Bridge)"
            },
            {
                "id": 144,
                "name": "Specialty Server: GPU"
            },
            {
                "id": 146,
                "name": "Sandy Bridge 1270"
            },
            {
                "id": 147,
                "name": "Specialty Server: 4u Mass Storage Dual Xeon 2000 (Sandy Bridge) Series"
            },
            {
                "id": 148,
                "name": "Specialty Server: 2u Mass Storage Dual Xeon 2000 (Sandy Bridge) Series"
            },
            {
                "id": 158,
                "name": "Quad Xeon 4000 Series (Sandy Bridge)"
            },
            {
                "id": 234,
                "name": "Quad Xeon E7-4800 v2 (Ivy Bridge) Series"
            },
            {
                "id": 248,
                "name": "Dual E5-2600 v3 Series (36 Drives)"
            },
            {
                "id": 251,
                "name": "Dual E5-2600 v3 Series (12 Drives)"
            },
            {
                "id": 253,
                "name": "Dual E5-2600 v3 Series (4 Drives)"
            },
            {
                "id": 255,
                "name": "Single E3-1270 (4 Drives)"
            },
            {
                "id": 257,
                "name": "Single E3-1270 v3 (4 Drives)"
            },
            {
                "id": 259,
                "name": "Single E5-2600 Series (4 Drives)"
            },
            {
                "id": 261,
                "name": "Single E3-1270 (2 Drives)"
            },
            {
                "id": 263,
                "name": "Dual E5-2600 Series (36 Drives)"
            },
            {
                "id": 265,
                "name": "Dual E5-2600 Series (12 Drives)"
            },
            {
                "id": 267,
                "name": "Quad E5-4600 Series (24 Drives)"
            },
            {
                "id": 269,
                "name": "Quad E7-4800  Series (6 Drives)"
            },
            {
                "id": 271,
                "name": "Quad E7-4800 v2 Series (24 Drives)"
            },
            {
                "id": 273,
                "name": "Dual E5-2600 (4 Drives)"
            },
            {
                "id": 295,
                "name": "SAP HANA Certified Servers"
            },
            {
                "id": 297,
                "name": "SAP NetWeaver Certified Servers"
            },
            {
                "id": 551,
                "name": "Dual E5-2600 v4 Series (4 Drives)"
            },
            {
                "id": 553,
                "name": "Dual E5-2600 v4 Series (12 Drives)"
            },
            {
                "id": 555,
                "name": "Dual E5-2600 v4 Series (36 Drives)"
            }
        ]
    }
}
```

#List Softlayer package options
`GET`  /sl/package/{packageID}/options

**Request**
`GET /sl/package/{packageID}/options`

**cURL**
```
curl -X GET "http://10.12.20.17:8080//sl/package/{packageID}/options
```
**Response**
Status
`200`

Body
```
{
    "status": 200,
    "data": {
        "categories": [
            {
                "code": "server",
                "name": "Server",
                "options": [
                    {
                        "id": 29366,
                        "description": "Quad Intel Xeon E7-4850 (40 Cores, 2.00 GHz)"
                    },
                    {
                        "id": 100099,
                        "description": "Quad Intel Xeon E7-4850 (40 Cores, 2.00 GHz)"
                    },
                    {
                        "id": 100101,
                        "description": "Quad Intel Xeon E7-4850 (40 Cores, 2.00 GHz)"
                    }
                ],
                "required": true
            },
            {
                "code": "os",
                "name": "Operating System",
                "options": [
                    {
                        "id": 44992,
                        "description": "CentOS 7.x (64 bit)"
                    },
                    {
                        "id": 22609,
                        "description": "CentOS 5.x (64 bit)"
                    },
                    {
                        "id": 30396,
                        "description": "CentOS 5.x (32 bit)"
                    },
                    {
                        "id": 170169,
                        "description": "CentOS 7.0 (64 bit)"
                    }
                ],
                "required": true
            },
            {
                "code": "port_speed",
                "name": "Uplink Port Speeds",
                "options": [
                    {
                        "id": 22829,
                        "description": "10 Mbps Public & Private Network Uplinks"
                    },
                    {
                        "id": 26737,
                        "description": "100 Mbps Public & Private Network Uplinks"
                    }
                ],
                "required": true
            }
        ],
        "datacenters": [
            "hkg02 - Hong Kong 2",
            "lon02 - London 2",
            "mon01 - Montreal 1",
            "ams03 - Amsterdam 3",
            "wdc04 - Washington 4",
            "sao01 - Sao Paulo 1",
            "sjc03 - San Jose 3"
        ]
    }
}
```

#List baremetals
`GET`  /bms/{deployment_name}

**Request**
`GET /bms/{deployment_name}`

**cURL**
```
curl -X GET "http://10.12.20.17:8080/bms/bps
```
**Response**
Status
`200`

Body
```
{
    "status": 200,
    "data": [
        {
            "id": 981961,
            "hostname": "baremetal-165-20160705-073112-025",
            "private_ip_address": "10.113.109.99",
            "public_ip_address": "169.50.68.83",
            "tags": [
                "bm.state.loading",
                "bm.bps",
                "bm.p.bm-pipeline"
            ],
            "hardware_status": "ACTIVE",
            "memory": 8,
            "cpu": 4,
            "provision_date": "2016-07-05T19:38:08+08:00"
        },
        {
            "id": 290736,
            "hostname": "baremetal-165-20160805-103440-936",
            "private_ip_address": "10.113.205.88",
            "public_ip_address": "159.122.228.174",
            "tags": [
                "bm.p.bm-pipeline",
                "bm.state.deleted",
                "bm.bps"
            ],
            "hardware_status": "ACTIVE",
            "memory": 8,
            "cpu": 4,
            "provision_date": "2016-08-05T19:57:02+08:00"
        },
        {
            "id": 780185,
            "hostname": "baremetal-165-20160805-103447-792",
            "private_ip_address": "10.113.109.202",
            "public_ip_address": "159.8.154.243",
            "tags": [
                "bm.p.bm-pipeline",
                "bm.state.using",
                "bm.bps"
            ],
            "hardware_status": "ACTIVE",
            "memory": 8,
            "cpu": 4,
            "provision_date": "2016-08-05T21:24:35+08:00"
        }
    ]
}
```

#List bmp server tasks
`GET`  /tasks

**Request**
`GET /tasks?latest={number}`

**cURL**
```
curl -X GET "http://10.12.20.17:8080/tasks?latest=4
```
**Response**
Status
`200`

Body
```
{
    "status": 200,
    "data": [
        {
            "id": 22,
            "description": "load stemcell",
            "start_time": "2016-07-14 09:28:24 +0000",
            "status": "running"
        },
        {
            "id": 21,
            "description": "load stemcell",
            "start_time": "2016-07-14 09:19:32 +0000",
            "status": "failed",
            "end_time": "2016-07-14 09:57:50 +0000"
        },
        {
            "id": 20,
            "description": "load stemcell",
            "start_time": "2016-07-14 08:52:55 +0000",
            "status": "completed",
            "end_time": "2016-07-14 08:59:53 +0000"
        },
        {
            "id": 19,
            "description": "load stemcell",
            "start_time": "2016-07-14 08:10:26 +0000",
            "status": "failed",
            "end_time": "2016-07-14 08:48:46 +0000"
        }
    ]
}
```

#List bmp server task output
`GET`  /task/{task_id}/txt/{level}

**Request**
`GET /task/{task_id}/txt/{level}`

**cURL**
```
curl -X GET "http://10.12.20.17:8080/task/19/txt/event
```
**Response**
Status
`200`

Body
```
{
    "status": 200,
    "data": [
        "INFO -- Begin to load stemcell bosh-stemcell-3262-softlayer-baremetal into bare metal baremetal-165-20160712-110628-713\n",
        "INFO -- Failed to load stemcell bosh-stemcell-3262-softlayer-baremetal into bare metal baremetal-165-20160712-110628-713\n"
    ]
}
```

#Update baremetal server state
`PUT`  /baremetal/{server_id}/{state}

**Request**
`PUT /baremetal/{server_id}/{state}`

**cURL**
```
curl -X PUT "http://10.12.20.17:8080/baremetal/981961/bm.state.new
```
**Response**
Status
`200`

#Create baremetal servers
`PUT`  /baremetals

**Request**
`PUT /baremetals`

**Parameters**
```
{
    "baremetal_specs": [
        {
            "bosh_ip": "10.11.20.119",
            "datacenter": "lon02",
            "domain": "softlayer.com",
            "name_prefix": "baremetal-",
            "server_spec": {
                "cores": 4,
                "memory": 4,
                "max_port_speed": 100,
                "public_vlan_id": 524956,
                "private_vlan_id": 524954,
                "hourly": true
            },
            "baremetal": true,
            "bm_stemcell": "bosh-stemcell-0.3-softlayer-baremetal",
            "bm_netboot_image": "bmp-netboot-ixgbe-lon02",
            "size": 1
        }
    ],
    "deployment": "bps"
}
```

**cURL**
```
curl -X PUT "http://10.12.20.17:8080/baremetals
```
**Response**
Status
`200`

Body
```
{
    "status": 200,
    "data": {
        "task_id": 23
    }
}
```