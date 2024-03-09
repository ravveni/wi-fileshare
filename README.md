```
           _       _____ __          __
 _      __(_)     / __(_) /__  _____/ /_  ____ _________
| | /| / / /_____/ /_/ / / _ \/ ___/ __ \/ __ `/ ___/ _ \
| |/ |/ / /_____/ __/ / /  __(__  ) / / / /_/ / /  /  __/
|__/|__/_/     /_/ /_/_/\___/____/_/ /_/\__,_/_/   \___/

spin up a quick LAN server with file sharing & streaming capabilities.
```
***
changes made to both `index.md` and within the `share/` directory are live.

## usage
- creates `wi-fileshare/` in home directory for the following, skips if exists:
  - creates an `index.md` for customizable landing page, skips if exists
  - creates a `share/` directory for filesharing, skips if exists
- continually converts `index.md` into temp `index.html`
- serves on port `1997`
