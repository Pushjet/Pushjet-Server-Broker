Pushjet Server Broker [![License](http://img.shields.io/badge/license-BSD-blue.svg?style=flat)](/LICENSE)
=====================
This part of pushjet manages the communication between clients that are 
listening to pushjet and the server. It uses ZeroMQ for the messaging 
which allows the server to be scaled at ease. The method for sending out 
messages is called pub/sub. It makes sure that only the messages that 
matter are processed by the right connections. 

## Installation 
```sh
make 
sudo make install
cp init.d/pushjet-broker.sh /etc/init.d/pushjet-broker
/etc/init.d/pushjet-broker start
```
