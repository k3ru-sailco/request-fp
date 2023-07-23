**Request FP**

 Very barebones library that logs and makes the TCP and TLS fingerprint of incoming requests accessible using the IP address of the connection. 

 **This could be useful for many different purposes, a few of them being:**
 - blocking or flagging requests.
 - detecting the use of proxies or VPNs.
 - creating a profile of a devices traffic to analyze and detect anomalies.

**Todo** 
- create database of fingerprints in order to detect anomalies.
- built-in anomaly detection and fingerprint verification.
- provide more data from incoming requests (if seen fit in the future).
