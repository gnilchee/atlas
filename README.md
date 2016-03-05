# atlas
A simple Dockerized web server written in Go leveraging docker-compose.

## Requirements to run as expected
* Docker (1.10 recommended)
* docker-compose (1.6.2 recommended)
* GO version 1.6+ (handled by Docker within the container)
* Trusted TLS cert chain and key (recommend generating via `letsencrypt certonly`)

## Manual steps
1. git clone https://github.com/gnilchee/atlas.git to directory of your choice.
2. Place static html/css content into "static" subdirectory.
3. Place TLS cert fullchain and private key into the "server" directory.
  * Update Dockerfile with appropriate names for TLS chain and key files.
  * Update atlas.go with appropriate names for TLS chain and key files.
4. Update atlas.go with appropriate names for TLS chain and key files.
5. Within the redirector function in atlas.go update your url to match your https:// so it can redirect your HTTP traffic correctly.

## To do's
* Automate the manual steps above other than static content and cert generation to limit any possible issue.  
