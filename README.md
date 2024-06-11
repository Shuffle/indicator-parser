# Indicator-parser
An IOC parser library written in Rust, made to handle multithreaded IOC management. Will be used in both apps and backend of Shuffle.

## TODO
- Prototype it single threaded
- Prototype multithreaded
- Ensure it's FAST

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

Example: A process explorer.exe spawning cmd.exe and then powershell.exe
Anomalous Activity: Unusual system or network activities, such as high CPU usage, unexpected network connections, or unusual file modifications.

