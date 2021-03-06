# A Software gNB for free5GC

**/!\ This project is deprecated and has been integrated into freecli https://github.com/srajdax/free5gc-cli**

![5G gNB](https://img.shields.io/badge/Golang-5G%20gNB-blue?logo=go)

- [Disclaimer](#disclaimer)
- [Roadmap](#roadmap)
- [Installation](#installation)
- [Before Launch](#before-launch)
- [Configuration](#configuration)
- [Service Exposed by REST Interface](#service-exposed-by-rest-interface)
  - [HTTP GET - /run/establish_pdu/:index](#http-get---runestablish_pduindex)
  - [HTTP GET - /run/ping_device/:index/:device](#http-get---runping_deviceindexdevice)
- [Usage](#usage)
- [Limitations](#limitations)

The gNB function was built on the model of the other free5GC CN functions using all the pattern and helper class defined by the free5GC team.

It ensures a seamless and immediate integration into free5GC without requiring any other dependencies.

The build and exection process is therefore the same as for the free5GC CN functions.

The gNB was tested using free5gc v3.0.3, v3.0.4 and free5gc-compose v3.0.4

Feel free to contribute !

## Disclaimer

This project provides an unofficial gNB for the free5gc project as the official test scripts does not cover all use-cases. This gNB is designed and tested only against free5gc, thus it follows some asumptions I have found into free5gc code:

- The TEID for the GTP-U tunnel is incremented by 1 at each PDU session request, therefore I assume the the N-PDU session established with the UPF would have the TEID = N value for the tunnel 

## Roadmap

- [ ] Update the registration procedure using the version v3.0.4
- [x] Add parameters to the configuration instead of hardcoding them directly
- [x] Separate the PDU Session Establishment and the processing of the Data Plane
- [x] Maintain the GTP UDP Socket Open
- [x] Refactor the base code of the gNB using the latest version of free5gc 3.0.4
- [ ] Use the new helper class of free5gc v3.0.4
- [x] Remove the use of mongo database
- [ ] Forge other packets than ICMP: HTTP and raw UDP socket
- [ ] Implement QoS at IP level
- [ ] Implement PDU Session Release procedure
- [ ] Establish a kernel based tunnel to allow other traffic generation other than the hardcoded one in go language

**Last commit needs to be tested**

## Installation

Follow the installation instructions provided by free5gc repository, but instead of clonine the project: `https://github.com/free5gc/free5gc.git`. Clone this forked version of the project `https://github.com/Srajdax/free5gc`.

The compilation and installation procedure of the gNB is the same as the other free5gc core functions, you can compile the functions using 

``` bash
cd ~/free5gc
go build -o bin/gnb -x src/gnb/gnb.go
```

Execute the function with the following command

``` bash
cd bin/gnb
./gnb
```

## Before Launch

You need to ensure that:

- The configuration of the `gnbcfg.cfg` is consistent with your free5gc `uerouting.yaml`, `smfcfg.cfg` and `upfcfg.cfg`
- Ensure that mongodb is running on the gNB host and also have the credentials loaded into the mongo free5gc database on the Core Network host

You will find a `dump` folder which contains the credentials for the imsi *imsi-2089300007487*. The credentials can be restored to the free5gc mongo database by using the mongorestore command `mongorestore dump`

## Configuration

The gNB `gnbcfg.cfg` configuration file is located in `free5gc/config` folder. A sample is also present into `gnb/config` folder.

``` yaml
info:
  version: 1.0.1
  description: "5G gNB initial local configuration"

configuration:
  ranName: gNB
  amfInterface:
    ipv4Addr: "127.0.0.1"
    port: 38412
  upfInterface:
    ipv4Addr: "10.200.200.102"
    port: 2152
  ngranInterface:
    ipv4Addr: "127.0.0.1"
    port: 9487
  gtpInterface:
    ipv4Addr: "10.200.200.1"
    port: 2152
  ueSubnet: "60.60.0.0/24"
  plmn:
    mcc: "208"
    mnc: "93"
  security:
    networkName: 5G:mnc093.mcc208.3gppnetwork.org
    k: 5122250214c33e723a5dd523fc145fc0
    opc: 981d464c7c52eb6e5036234984ad0bcf
    op: c9e8763286b5b9ffbdf56e1297d0887b
    sqn: 16f3b3f70fc2
  snssai:
    sst: 1
    sd: "010203"
  ue:
    - SUPI: imsi-2089300007487
      ipv4: 60.60.0.1
    - SUPI: imsi-2089300007486
      ipv4: 60.60.0.2
  sbi:
    scheme: http
    ipv4Addr: 127.0.0.1
    port: 32000
  networkName:
    full: free5GC
    short: free
```

The following Diagram represents the configuration file above

![diagram_gNB](https://user-images.githubusercontent.com/41422704/88692144-07d6a700-d0fe-11ea-836d-56df98ffa93a.png)

## Service Exposed by REST Interface

The gNB exposes two command interfaces

| Service                 | Url                                  | Status      |
| ----------------------- | ------------------------------------ | ----------- |
| Establish a PDU Session | /run/establish_pdu/:index            | Implemented |
| Ping a Device           | /run/ping_device/:identifier/:device | Implemented |

### HTTP GET - /run/establish_pdu/:index 

GET Parameters:

- index: The identifier of the UE in the provided UE config list (gnbcfg.cfg)

### HTTP GET - /run/ping_device/:index/:device

GET Parameters:

- index: The identifier of the PDU Session. Currently the index of the List containing all the PDU Sessions
- device: The destination IP to ping.

## Usage

After launching the gnb, with simple tools such as curl, you can control the gNB using:

``` bash
curl -d {} http://localhost:32000/run/ping_device/0/60.60.0.101
```

## Limitations

For the moment, only one PDU session could be established per UE to match with the UE IP configuration