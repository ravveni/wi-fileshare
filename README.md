```
           _       _____ __          __
 _      __(_)     / __(_) /__  _____/ /_  ____ _________
| | /| / / /_____/ /_/ / / _ \/ ___/ __ \/ __ `/ ___/ _ \
| |/ |/ / /_____/ __/ / /  __(__  ) / / / /_/ / /  /  __/
|__/|__/_/     /_/ /_/_/\___/____/_/ /_/\__,_/_/   \___/

spin up a quick LAN server with file sharing & streaming capabilities.
```
***
changes made to both `index.md` and within the shared directory are live.

## usage
### basic:
`$ ./wi-fileshare`

within the current working directory:
- creates an `index.md` for landing page, skips if exists
- creates a `share` directory for filesharing, skips if exists
- converts `index.md` to `index.html`
- serves on port 8080

### custom:
`$ ./wi-fileshare -p=PORT -i=INDEX.MD-FILEPATH -s=SHARE-DIRECTORY-FILEPATH`

all arguments are optional, defaults same from basic usage
