### Network Fundamentals

#### Default Route

### Route Table

#### Default Gateway

#### Internet Control Message Protocol (ICMP)

#### DNS

#### DNS Resolver

#### Primary and Secondary DNS Servers

#### Understanding 0.0.0.0

#### IP CIDR Notation

#### Network Namespaces

#### Network Interface

#### NAT(Network Address Translation)

#### IPtables

#### Loopback Interface

#### Virtual Ethernet

#### More on Virtual Networking

#### Open vSwitch

#### Bridge networking

#### Setting up a network

# Networking 

## Network Flags

### LOOPBACK

- The Internet Protocol (IP) specifies a loopback network with the (IPv4) address 127.0.0.0/8. Most IP implementations support a loopback interface (lo0) to represent the loopback facility. Any traffic that a computer program sends on the loopback network is addressed to the same computer. The most commonly used IP address on the loopback network is 127.0.0.1 for IPv4 and ::1 for IPv6. The standard domain name for the address is localhost.

### UNICAST

- Unicast is the term used to describe communication where a piece of information is sent from one point to another point. In this case there is just one sender, and one receiver.
Unicast transmission, in which a packet is sent from a single source to a specified destination, is still the predominant form of transmission on LANs and within the Internet. All LANs (e.g. Ethernet) and IP networks support the unicast transfer mode, and most users are familiar with the standard unicast applications (e.g. http, smtp, ftp and telnet) which employ the TCP transport protocol.

### BROADCAST

- A broadcast address is a network address used to transmit to all devices connected to a multiple-access communications network. A message sent to a broadcast address may be received by all network-attached hosts.In short 'device can send traffic to all hosts on the link'.

- Broadcast transmission is supported on most LANs (e.g. Ethernet), and may be used to send the same message to all computers on the LAN (e.g. the address resolution protocol (arp) uses this to send an address resolution query to all computers on a LAN, and this is used to communicate with an IPv4 DHC server). Network layer protocols (such as IPv4) also support a form of broadcast that allows the same packet to be sent to every system in a logical network (in IPv4 this consists of the IP network ID and an all 1's host number).

### MULTICAST

- Multicast is the term used to describe communication where a piece of information is sent from one or more points to a set of other points. In this case there is may be one or more senders, and the information is distributed to a set of receivers (theer may be no receivers, or any other number of receivers).

- Multicasting is the networking technique of delivering the same packet simultaneously to a group of clients. IP multicast provides dynamic many-to-many connectivity between a set of senders (at least 1) and a group of receivers. The format of IP multicast packets is identical to that of unicast packets and is distinguished only by the use of a special class of destination address (class D IPv4 address) which denotes a specific multicast group. Since TCP supports only the unicast mode, multicast applications must use the UDP transport protocol.

- Unlike broadcast transmission (which is used on some local area networks), multicast clients receive a stream of packets only if they have previously elect to do so (by joining the specific multicast group address). Membership of a group is dynamic and controlled by the receivers (in turn informed by the local client applications). The routers in a multicast network learn which sub-networks have active clients for each multicast group and attempt to minimise the transmission of packets across parts of the network for which there are no active clients.

- The multicast mode is useful if a group of clients require a common set of data at the same time, or when the clients are able to receive and store (cache) common data until needed. Where there is a common need for the same data required by a group of clients, multicast transmission may provide significant bandwidth savings (up to 1/N of the bandwidth compared to N separate unicast clients).

- The majority of installed LANs (e.g. Ethernet) are able to support the multicast transmission mode. Shared LANs (using hubs/repeaters) inherently support multicast, since all packets reach all network interface cards connected to the LAN. The earliest LAN network interface cards had no specific support for multicast and introduced a big performance penalty by forcing the adaptor to receive all packets (promiscuous mode) and perform software filtering to remove all unwanted packets. Most modern network interface cards implement a set of multicast filters, relieving the host of the burden of performing excessive software filtering.

### UP

- The UP means that it has been enabled and functioning. This can be controlled by you (or a script) using the ip link set device up of ifconfig device up command.

### - LOWER_UP

- is a physical layer link flag (the layer below the network layer, where IP is generally located). LOWER_UP indicates that an Ethernet cable was plugged in and that the device is connected to the network.

- The LOWER_UP is the state of the Ethernet link (or other link layer protocol). It's defined as Driver signals L1 up, which basically means the cable is fitted and it can see another device on the other end of the cable.

### NO_CARRIER

- NO-CARRIER means that the network jack/bridge detects no signal on the line. This is usually because the network cable is unplugged or broken on physical netwrok or no connection to the otherside on virtual networks. In rare cases it can also be hardware failure or a driver bug. Have you checked the cables and rebooted the system.

- Linux's network stack uses the NO CARRIER status for a network interface that is turned on ("up") but cannot be connected because the physical layer is not operating properly, e.g. because an ethernet cable is not plugged in.
  
## Other stuff

### Eithernet MAC

- The Ethernet network uses two hardware addresses which identify the source and destination of each frame sent by the Ethernet. The MAC destination address (all 1 s) is used to identify a broadcast packet (sent to all connected computers in a broadcast domain) or a multicast packet (lsb of 1st byte=1) (received by a selected group of computers).

- The hardware address is also known as the Medium Access Control (MAC) address, in reference to the IEEE 802.x series of standards that define Ethernet. Each computer network interface card is allocated a globally unique 6 byte MAC source address when the factory manufactures the card (stored in a PROM). This is the normal source address used by an interface for completing the MAC source address field, and also a filter for received frames.

- A computer sends all packets which it creates with its own hardware source address, and receives all packets, which match its hardware address or the broadcast address. When configured to use multicast, a selection of multicast hardware addresses may also be received.

### mtu

- Maximum Transfer Unit (max size of DL frame in IP) (cf MSS)

### qdisc


- qdisc  is  short for 'queueing discipline' and it is elementary to understanding traffic control. Whenever the kernel needs to send a packet to an interface, it is enqueued to the qdisc con‐
figured for that interface. Immediately afterwards, the kernel tries to get as many packets as possible from the qdisc, for giving them to the network adaptor driver.

- A simple QDISC is the 'pfifo' one, which does no processing at all and is a pure First In, First Out queue. It does however store traffic when the network interface can't handle it momentarily.

- Explore tc command

### qdisc noqueue

- Although the noqueue queuing discipline does drop all packets
  queued onto it, in practice that never happens.  Instead when
  a packet is sent over a device it checks if it is using the
  "noqueue" discipline.  If so the device sends the packet
  immediately, or drops it if it can't be sent.  Thus the
  noqueue discipline really means "don't queue this packet".

- noqueue is the queuing discipline that is used by default
  for :virtual: devices, meaning it is the queuing discipline
  installed when a virtual device is first created.  It is
  also the queuing discipline used after you "tc qdisc del"
  another queuing discipline from a virtual device.
  
- You can _not_ manually change a queuing discipline for a
  device or class to noqueue using "tc qdisc add noqueue".  You can
  get around this for virtual devices by deleting their queuing
  discipline.  It is not possible to assign the noqueue queuing
  discipline to physical devices or classes.

### Network trouble shooting 

### Interface Information Breakdown

```ip link show dev eth0```

2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether b8:27:eb:57:91:1c brd ff:ff:ff:ff:ff:ff
    inet 172.31.0.246/24 brd 172.31.0.255 scope global eth0
       valid_lft forever preferred_lft forever
    inet6 fe80::ba27:ebff:fe57:911c/64 scope link
       valid_lft forever preferred_lft forever
To break this down a bit, ip gives the following information:

eth0: name of the interface
BROADCAST: it’s a broadcast interface which means it has a valid broadcast address
MULTICAST: it supports multicast
UP: the interface has been enabled and running (drivers loaded and working)
LOWER_UP: there is signal activity at the physical layer (a cable is plugged in)
mtu 1500: the maximum  transfer unit is 1500
qdisc pfifo_fast: used for packet queueing
state UP: network interface is up
group default: interface group
qlen 1000: transmission queue length

inet 172.31.0.246/24: IPv4 address
brd 172.31.0.255: broadcast address
scope global eth0: valid everywhere
valid_lft forever: valid lifetime for IPv4 address
preferred_lft forever: preferred lifetime for IPv4 address
inet6 fe80::ba27:ebff:fe57:911c/64: IPv6 address
scope link: valid only on this device
valid_lft forever: valid lifetime for IPv6 address
preferred_lft forever: preferred lifetime for IPv6 address

### Interface Statistics

```ip -s link show eth0```

### Routing Information

```ip route show```
