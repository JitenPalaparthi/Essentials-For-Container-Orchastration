# FAQS

- Pinging systems from another network depends on the actual implementation of the network. If these two subnets are connected via a switch or a hub it can be difficult to find the system from another network but if its connected via a router it may forward packets from one subnet to another. If you can communicate to systems at different networks then all your physical connection of your network goes through a router. There is also another possiblity where for security purposes ping and port scanning functions are disable by the firewall either at the destination or at the router. If you can see other computers but no reply from ping requests then probably the firewall must be blocking the communication. This typically happens when you ping from a Linux based computer with a computer running windows because windows firewall blocks these type of ping and port scan requests.

- What is North-North vs East-West traffic?
    "east-west" traffic refers to traffic within a data center -- i.e. server to server traffic.
    "North-south" traffic is client to server traffic, between the data center and the rest of the network