![Cosmoport Logo](/cosmoport.png)

# Cosmoport
> "Simplicity boils down to two steps: Identify the essential. Eliminate the rest." 
> 
> — Leo Babauta

Cosmoport is a fast, almost-zero weight web server written in Go with human-readable 
delightful YAML configuration.

## Features
Cosmoport is simple, and useful for small homelabs or VPS's with multiple projects on a single machine.
It routes HTTP traffic to different ports, or static folders.

## Configuration

```yaml
rules:                              # define list of routing rules
    todo:                           # name one
        domain:  todo.anafro.com    # set a domain for it
        backend: 3000               # forward it to a port. done!
    images:
        domain:  img.anafro.com
        static:  ~/images           # or deliver files instead!
```