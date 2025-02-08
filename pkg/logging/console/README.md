# Stdout Logging

We will try at least 5 most popular loggers to log data in `JSON` format to stdout/stderr, the most popular format for logging nowadays.
- log/slog
- zap
- zerolog
- logrus
- TBD

### Implementation
Various logger provide different levels of logging and methods. So we will implement common wrapper for them to swap them easily.

## Requirements
### Our wrapper logs messages with 5 levels of logging.
- Debug
- Info
- Warning
- Error
- Critical

### It can enrich logging message with info:
- name of logger
- timestamp
- logging level
- main message
- extra arguments any type, which will be added as new fields for our json object
