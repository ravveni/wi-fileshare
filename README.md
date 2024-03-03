# quick-server
spin up a quick LAN server with file sharing & streaming.
***
changes made to `index.md` will appear after server restart.

changes made within the `share/` directory are live.

## requirements
- golang [Download latest release](https://go.dev/dl/)

## usage
1. edit `index.md` however you want
2. add any files to share locally to `share/`
3. `go build && ./quick-server`
4. ???
5. profit

re-building the program for subsequent uses isn't necessary unless changes were made in `go` code.

## tested streaming formats
* audio
  - mp3

* video
  - mp4
