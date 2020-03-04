# Simple API

**Simple API** provides a basic API server which is useful for various testing scenarios.

Endpoints available

| Endpoint | Description |
| -------- | ----------- |
| getVersion | Returns the application version number in JSON format. |
| get1KBFile | Returns a 1KB file download |
| get1MBFile | Returns a 1MB file download |
| podReady | This is designed for use when running the application in a Kubernetes cluster. It will return either a ```200/OK``` response if the application is running, or ```500/Internal Server Error``` if the application is being terminated. This endpoint can be used for the liveness and readiness probe endpoints. |
| podTerminate | Endpoint which can be called if running in Kubernetes to let the app know it's being terminated. This can be configured as the preStop endpoint.  After this endpoint is called, the liveness probe via the podReady endpoint will return a ```500``` status indicating that the server is no longer serving requests. The server will wait 30 seconds when this endpoint is called to allow Kubernetes time to perform a failover |

## Use Cases

### Blue / Green / Canary Testing

You can simulate the process of updating an application and deploying a new version by updating the version number returned by the ```getVersionValue``` function. You can then run a previous and the "new" version of the application side by side, demonstrating that each returns a different value when the ```getVersion``` endpoint is called. When you apply your canary or blue / green configuration, you can demonstrate that a certain percentage of hits are now being directed to the "new" version, or that you've switch blue / green environments.

### Performance Testing

This application is written in Go using the fasthttp library and is intended to be lightweight and fast. The application will automatically spawn new threads to handle incoming connections, so it's ideal to use as the target when performance testing an Ingress Controller.