# ⚙️ Build and Run Docker Image with Automatic Tor Bridge Discovery

---

## 🔨 Build the Docker Image

First, navigate to the **project root directory** (one level above `docker/`). Then build:

```bash
cd ../
docker build -t tor-scanner -f docker/Dockerfile .
```

---

## 🚀 Run the Container

```bash
docker run -p 9050:9050 --rm -d --name tor-node tor-scanner
```

Once running, Tor will be available as a SOCKS5 proxy at `localhost:9050`.

### ✅ Verify the proxy is working:

```bash
curl -x socks5h://localhost:9050 https://ifconfig.co
```

---

## ❓ Tor Takes Too Long to Start?

Even with working bridges, Tor may take up to a minute to fully bootstrap.
Check the logs to monitor progress:

```bash
docker logs tor-node
```

A successful startup will show lines like:

```
Nov 16 09:31:15.000 [notice] Bootstrapped 56% (loading_descriptors): Loading relay descriptors
Nov 16 09:31:15.000 [notice] Bootstrapped 62% (loading_descriptors): Loading relay descriptors
Nov 16 09:31:16.000 [notice] Bootstrapped 67% (loading_descriptors): Loading relay descriptors
Nov 16 09:31:16.000 [notice] Bootstrapped 75% (enough_dirinfo): Loaded enough directory info to build circuits
Nov 16 09:31:17.000 [notice] Bootstrapped 90% (ap_handshake_done): Handshake finished with a relay to build circuits
Nov 16 09:31:17.000 [notice] Bootstrapped 95% (circuit_create): Establishing a Tor circuit
Nov 16 09:31:18.000 [notice] Bootstrapped 100% (done): Done
```

> ✅ Look for: `Bootstrapped 100% (done): Done` — this means Tor is ready.

### 🔁 If stuck, restart:

```bash
docker stop tor-node
docker run -p 9050:9050 --rm -d --name tor-node tor-scanner
```
