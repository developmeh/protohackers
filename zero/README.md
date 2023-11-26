## Zero

Also known as the smoke test for https://protohackers.com this helped me get up to 
speed with golang and the basics of the TCPListener and some of the network stack.

## Usage

note the defaults are hard coded in the main.go file

IP: 0.0.0.0

Port: 9999

```bash
go run main.go
```

## Testing

```bash
nc 0.0.0.0 9999 -w 1 -q 1 < test.txt
```

