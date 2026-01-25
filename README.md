# Tor-scanner

[English](README-en.md)

**Tor-scanner** — это быстрый и простой инструмент для поиска рабочих мостов (bridges) для Tor.
Он проверяет список мостов на доступность и отбирает только те, которые отвечают быстро и стабильно.

Подходит для тех, кто живёт в странах с цензурой и нуждается в надёжных мостах для подключения к Tor.

---

![Demo](.github/demo.gif)

## 🚀 Как использовать

Запустите приложение с нужными параметрами:

```bash
tor-scanner \
  --count 10 \
  --threads 20 \
  --ping 1000 \
  --country ru,es,us \
  --output tor.conf \
  --input bridges.list
```

### Параметры:

| Параметр | Описание |
|----------|----------|
| `--count` | Сколько рабочих мостов нужно найти (по умолчанию: 5) |
| `--threads` | Сколько запросов делать одновременно (по умолчанию: 50) |
| `--ping` | Максимальная задержка (в миллисекундах), чтобы мост считался рабочим (по умолчанию: 250) |
| `--country` | Фильтр по странам. Пример: `ru,es,us` — будут проверяться только мосты из этих стран |
| `--output` | Имя файла, куда сохранить найденные мосты (например: `bridges.txt`) |
| `--input` | Ссылка или путь к файлу, содержащие список мостов. По умолчанию используется оффициальный список |

---

## 💡 Примеры

### Найти 15 быстрых мостов из России и США, сохранить в файл:

```bash
tor-scanner --count 15 --threads 100 --ping 100 --country ru,us --output tor.conf
```

### Проверить список мостов и сохранить только рабочие:

```bash
tor-scanner --input my-bridges.list --output working-bridges.txt --ping 200
```

или

```bash
tor-scanner --input "https://onionoo.torproject.org/details?type=relay&running=true&fields=fingerprint,or_addresses,country" --output working-bridges.txt --ping 200
```

---

## 🛠️ Как собрать

Убедитесь, что у вас установлен Go (1.19+):

```bash
git clone https://github.com/tamper000/tor-scanner.git
cd tor-scanner
go build -o tor-scanner cmd/main.go
```

Готово! Теперь можно запускать `./tor-scanner`.

[🐳 Использование Docker](docker/howto.md)
