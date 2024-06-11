# Indicator-parser
An IOC parser library written in Rust, made to handle multithreaded IOC management. Will be used in both apps and backend of Shuffle.

## TODO
- Prototype it single threaded
- Prototype multithreaded
- Ensure it's FAST

## Using it in Python
```python
pip install indicator-parser
```

```
import indicator-parser

types = ["urls", "ipv4s"]
output = indicator-parser.parse("as10.0.0.10df1234 1.2.3.4 https://google.com", types)

print(output)
```

Expected output:
```
########
[{
  "type": "ipv4",
  "data": "10.0.0.10",
  "internal": True,
 },
  {
  "type": "ipv4",
  "data": "1.2.3.4",
  "internal": False,
 },
 {
  "type": "url",
  "data": "https://google.com"
 },
 {
  "type": "domain",
  "data": "google.com"
 }
]
```

## Datatypes (asap)
- IPv4
- URL
- Domains

### Future: 
- Mitre Att&ck tactics & techniques
- File paths
- Registry keys
- Email related stuff

### Optimizing (long-term)
- Process Trees: Specific sequences of process creation that indicate malicious behavior.
- CyberChef-related toolkit 

