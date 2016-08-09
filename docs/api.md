Baremetal Provision Server API
===============================

This is a brief introduction of Baremetal Provision Server API.

#Login bmp server
`GET`  /login/{username}/{password}

**Request**
`GET /login/{username}/{password}`

**cURL**
```
curl -X GET "http://{serverIP}:8080/login/{username}/{password}
```
**Response**
Status
`200 OK`

#Check bmp server status
`GET`    /info

**Request**
`GET /info`

**cURL**
```
curl -X GET "http://{serverIP}:8080/info
```
**Response**
Status
`200 OK`

Body
```
{
    "name" : "fake-name",
    "version" : "fake-version"
}
```

#List bmp server stemcells
`GET`  /stemcells

**Request**
`GET /stemcells`

**cURL**
```
curl -X GET "http://{serverIP}:8080/stemcells
```
**Response**
Status
`200 OK`

Body
```
{
  [
    "fake-stemcell-0",
    "fake-stemcell-1"
  ]
}
```

#List Softlayer packages
`GET`  /sl/packages

**Request**
`GET /sl/packages`

**cURL**
```
curl -X GET "http://{serverIP}:8080/sl/packages
```
**Response**
Status
`200 OK`

Body
```
{
    "packages":[
      {
        "id": 0,
        "name": "name0"
      },
      {
        "id": 1,
        "name": "name1"
      }
    ]
}
```

#List Softlayer package options
`GET`  /sl/package/{packageID}/options

**Request**
`GET /sl/package/{packageID}/options`

**cURL**
```
curl -X GET "http://{serverIP}:8080//sl/package/{packageID}/options
```
**Response**
Status
`200 OK`

Body
```
 {
    "categories": [
      {
        "code": "code0",
        "name": "name0",
        "options": [
          {
            "id": 0,
            "description": "description0"
          },
          {
            "id": 1,
            "description": "description1"
          }
        ],
        "required": true
      },
      {
        "code": "code1",
        "name": "name1",
        "options": [
          {
             "id": 0,
            "description": "description0"
          }
        ],
        "required": false
      }
    ],
    "datacenters" : [
      "datacenter0 - location0",
      "datacenter1 - location1"
    ]
}
```

#List baremetals
`GET`  /bms/{deployment_name}

**Request**
`GET /bms/{deployment_name}`

**cURL**
```
curl -X GET "http://{serverIP}:8080/bms/{deployment_name}
```
**Response**
Status
`200 OK`

Body
```
{
  [
    {
      "id":   0,
      "hostname": "hostname0",
      "private_ip_address": "private_ip_address0",
      "public_ip_address": "public_ip_address0",
      "tags": [
        "bm.state.new",
        "bm.bps"
      ],
      "memory": 0,
      "cpu": 0,
      "provision_date": "2016-01-01T00:00:00-00:00"
    },
    {
      "id":   1,
      "hostname": "hostname1",
      "private_ip_address": "private_ip_address1",
      "public_ip_address": "public_ip_address1",
      "tags": [
        "bm.state.new",
        "bm.bps"
      ],
      "memory": 1,
      "cpu": 1,
      "provision_date": "2016-01-01T00:00:00-00:00"

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
curl -X GET "http://{serverIP}:8080/tasks?latest={number}
```
**Response**
Status
`200 OK`

Body
```
{
  [
    {
      "id": 0,
      "description": "fake-description-0",
      "start_time": "fake-start-time-0",
      "status": "fake-status-0",
      "end_time": "fake-end-time-0"
    },
    {
      "id": 1,
      "description": "fake-description-1",
      "start_time": "fake-start-time-1",
      "status": "fake-status-1",
      "end_time": "fake-end-time-1"
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
curl -X GET "http://{serverIP}:8080/task/{task_id}/txt/event
```
**Response**
Status
`200 OK`

Body
```
{
  [
    "INFO -- event0",
    "ERROR -- event1"
  ]
}
```

#Update baremetal server state
`PUT`  /baremetal/{server_id}/{state}

**Request**
`PUT /baremetal/{server_id}/{state}`

**cURL**
```
curl -X PUT "http://{serverIP}:8080/baremetal/{server_id}/{state}
```
**Response**
Status
`200 OK`

#Create baremetal servers
`PUT`  /baremetals

**Request**
`PUT /baremetals`

**Parameters**
```
{
    "baremetal_specs": [
        {
            "bosh_ip": "{bosh_ip}",
            "datacenter": "{datacenter}",
            "domain": "{domain}",
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
            "bm_stemcell": "{bm_stemcell}",
            "bm_netboot_image": "{bm_netboot_image}",
            "size": 1
        }
    ],
    "deployment": "bps"
}
```

**cURL**
```
curl -X PUT "http://{serverIP}:8080/baremetals
```
**Response**
Status
`200 OK`

Body
```
{
    "task_id": 1
}
```
