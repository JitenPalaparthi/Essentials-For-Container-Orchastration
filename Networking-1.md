# Essentials for Containers | docker | nomad | k8s | k3s | simplenetes

## Networking

### Network OSI Model

- The Open Systems Interconnection model is a conceptual model that describes the universal standard of communication functions of a telecommunication system or computing system, without any regard to the system's underlying internal technology and specific protocol suites. 

| Layers | Physical | Datalink | Network | Transport | Session | Presentation | Application |
| -------| -------- | -------- | ------- | --------- | ------- | ------------ | ----------- |
| Protocol Data Unit[PDU] | Bit | Frame | Packet | Segment/Datagram | Data | Data | Data |
| Notes | Physical Connection Between Devices | This layer is the protocol layer that transfers data between nodes on a network segment across the physical layer | Addressing routing and delivery of datagrams between point to point networks | reliable delivery of segments between points on a network | interhost communication , managing session between applications.It provides services such as connection-oriented communication, reliability, flow control, and multiplexing. | Data representation , encryption decryption, converting machine dependent data to machine independent data | Network processes to the applications like web applications , Emails, Messengers |
| Device | Eithernet PHY,Hubs, Repeaters, Cables, Fibers, Wireless | Bridges, Modems, Network cards, 2-layer switches | Routers, Brouters, 3-layer switches | Gateways, Firewalls | Gateways, Firewalls, PC’s | Gateways, Firewalls, PC's | Gateways,Firewalls, all end devices like PC’s, Phones, Servers |
| Communication | Electrons(Signals On-Off) | Frames are sent to a specific switch port based on destination MAC addresses | packets are sent to a specific next-hop IP address, based on destination IP address | Reliable Communication on IP:Ports | Data communication on IPs and Ports on IPs and Ports on Protocals like TCP,UDP,HTTP| Data Communication on IPs and Ports on Protocals like TCP,UDP,HTTP | Data Communication on IPs and Ports on Protocals like TCP,UDP,HTTP |
| Services | The physical layer performs bit-by-bit or symbol-by-symbol data delivery over a physical transmission medium | -Encapsulation of network layer data packets into frames -Physical addressing (MAC addressing) -Error Control (Automatic Repeat Request) - LAN switching (packet switching), including MAC filtering, Spanning Tree Protocol (STP), Shortest Path Bridging (SPB) and TRILL (TRansparent Interconnection of Lots of Links) - Data packet queuing or scheduling | Connectionless communication from source to destination via one or more networks | - Connection oriented communication as a data stream . - Order or delivery . - Reliability . - Flow Control . - Multiplexing: Ports can provide multiple end points on a single node. | - Session checkoint of recovery. - Authentication. - Authorization. | - Data conversion. - Charcode translation . - Compression . - Encryption-Decryption . - Serialization | - Process-Process Communication. - RESTFul Communication - gPRC,SMTP Communication via applications, Client-Server , Peer-Peer networking model implementations |

### Network devices

#### switch

- Switches are very important to form a network.They can connect multiple devices such as computers, printers , wireless access points and servers which are in same building or in the campus.
- Switches comes with ports,firmware or OS.
- Unmanaged switches: Simple switches with no or fewer configuration options. Generally used in homes and small businesses.
- Managed switches: Manged switches gives more configuration options, grater security more features with greater control. In general managed switches comes with software so that most of the configurations can be done through software.Users can access switch through CLI or GUI based interfaces.
![images/switch.jpeg](Switch)
  
##### l2 switch

##### l3 switch

##### cloud managed

#### router
- Routers are used to different netwroks to communicate.
- Router routes data packets between computer networks

#### wireless access points