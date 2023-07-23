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


**Example Usage**

```go
package main

import (
	"github.com/gofiber/fiber/v2"
	request_fp "github.com/k3ru-sailco/request-fp/lib"
)

func main() {
	/* start listening on the wifi network interface (i am using a macbook in this example)*/
	go request_fp.StartListening("en0")
	router := fiber.New()
	router.Get("/fingerprint-me", func(c *fiber.Ctx) error {
		/* returns the fingerprint of the incoming request */
		data, _ := request_fp.RequestFP(c.IP())
		return c.JSON(data)
	})
	router.Listen(":8080")
}
```

**When Containerizing**

```dockerfile
# libpcap must be installed prior to building your application
# in order for it to build properly or at all.
RUN apt-get update && apt-get install -y libpcap-dev
```

**Issues & Requests**

Please open an issue if you have something you'd like integrated or if you're unable to fix something. PRs welcome if you are contributing something useful.