# tileserver-proxy

Use tile servers in your favorite GIS application that only allow access with a referer

## Usage

```
TARGET_HOSTNAME=tiles.example.org REFERER=example.org ./tileserver-proxy
```

Then replace `https://tiles.example.org/map/foo/{z}/{x}/{y}` with `http://localhost:8888/map/foo/{z}/{x}/{y}`.
