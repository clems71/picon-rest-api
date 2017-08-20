# Picon REST API

Control the PICON hat from a comfy REST API.

## Build

```bash
make
```

This will build both the desktop (`server`) version and the RPI version for ARM processors (`server-arm`), from your desktop machine (for more performance). You could also build it from your raspberry-pi, but it's longer... :)

## Use

On desktop with a 'fake' hat:

```bash
./server
```

On RPI with the real PICON hat:

```bash
USE_PICON=yes ./server-arm
```