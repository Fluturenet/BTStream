# BTStream
Streaming Media over DHT Bittorrent Network

The aim of this software is to create a uncensorable (video)content network. It's BackBone is the Bittorrent [DHT Network] ~~~ and [Storing arbitrary data in the DHT] specifications~~~.

![uncached image](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://github.com/Fluturenet/BTStream/raw/master/doc/diagram1.txt)

## The relationship with HTTP Live Streaming
![diagram2](doc/diagram2.png)

The basic idea of this project is heavily based on [HTTP Live Streaming] protocol. It's a sort of HLS without web server(CDN) where the M3U8 is stored under the DHT network and the video fragments are .torrent files.
The force of HLS is the direct control of the CDN over the users and the ability to select who and how can connect to a particular stream and who can upload the stream itself, on the other side the CDN itself can be censored or blocked.

With BTStream we are givin back the power to the users:
* Content Providers can stream without censorship and a low bandwidth usage
* Users can watch their streams all over the world

# Quick Start

> btstream


> vlc http://localhost:8080/tr?ih=[hash of a torrent]

[DHT Network]: http://bittorrent.org/beps/bep_0005.html
[Storing arbitrary data in the DHT]: http://bittorrent.org/beps/bep_0044.html
[HTTP Live Streaming]:https://en.wikipedia.org/wiki/HTTP_Live_Streaming
