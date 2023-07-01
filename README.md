# The GOLANG Utility
## Summary
The repository includes applications and packages written in the GO language. It is targeting applications that 
works as a gRPC client including a desktop GUI. The GO language is mature and good while its GUI framework FYNE still
is in development mode. It is lacking some controls/widgets so the development of the Event Viewer was interrupted and 
paused until a better FYNE is released.

### The Log Chain Package
The log chain package is a list of loggers. It replaces the GO standard logger functionality with the possibility of
having multiple loggers. Log to console, log to file and log to syslog are supported. 

### The Event Viewer
The Event Viewer is a simple viewer for the Event Log Server. It uses gRPC/Protobuf for 
communicating with the server. 




