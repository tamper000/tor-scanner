# Tor-scanner

**Tor-scanner** is a fast and simple tool to find working bridges for Tor.
It checks a list of bridges for availability and selects only those that respond quickly and reliably.

Suitable for users living in censored regions who need reliable bridges to connect to Tor.

---

## 🚀 How to Use

Run the application with the desired parameters:

```bash
tor-scanner \
  --count 10 \
  --threads 20 \
  --ping 1000 \
  --country ru,es,us \
  --output tor.conf \
  --input bridges.list
```

### Parameters:

| Parameter | Description |
|----------|-------------|
| `--count` | Number of working bridges to find (default: 5) |
| `--threads` | Number of concurrent requests (default: 50) — higher values mean faster scanning |
| `--ping` | Maximum delay (in milliseconds) for a bridge to be considered responsive (default: 250) |
| `--country` | Filter bridges by country codes. Example: `ru,es,us` — checks only bridges from these countries |
| `--output` | File to save found bridges (e.g., `bridges.txt`) |
| `--input` | URL or path to a file containing a list of bridges. Uses the official list by default |

---

## 💡 Examples

### Find 15 fast bridges from Russia and USA, save to file:

```bash
tor-scanner --count 15 --threads 100 --ping 100 --country ru,us --output tor.conf
```

### Check a list of bridges and save only working ones:

```bash
tor-scanner --input my-bridges.list --output working-bridges.txt --ping 200
```

or

```bash
tor-scanner --input "https://onionoo.torproject.org/details?type=relay&running=true&fields=fingerprint,or_addresses,country" --output working-bridges.txt --ping 200
```

---

## 🛠️ How to Build

Make sure you have Go (1.19+) installed:

```bash
git clone https://github.com/tamper000/tor-scanner.git
cd tor-scanner
go build -o tor-scanner cmd/main.go
```

Done! Now you can run `./tor-scanner`.

[🐳 Using Docker](docker/howto-en.md)
