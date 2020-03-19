package main

import (
	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
"github.com/anacrolix/torrent/storage"

	log "github.com/sirupsen/logrus"
	"net/http"
	"io"
	"fmt"
)

var tr *torrent.Client

func init_torrent() (err error) {
	cf := torrent.NewDefaultClientConfig()
	cf.DataDir = dataDir
	cf.DefaultStorage =  storage.NewFileByInfoHash(dataDir)
	tr, err = torrent.NewClient(cf)
	return err
}

func tr_handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	infoHashString := query.Get("ih")
		infoHash := new(metainfo.Hash)
		errHash:= infoHash.FromHexString(infoHashString)
		if errHash!=nil {
			                        log.Warn("error loading hash:", infoHashString)
                        w.WriteHeader(http.StatusBadRequest)
                        io.WriteString(w,"Badformed Hash")
                        return

		}
		torrentFH, isNew := tr.AddTorrentInfoHash(*infoHash)
		if isNew != true {
			<-torrentFH.GotInfo()
			torrentFH.DownloadAll()
		}
	<-torrentFH.GotInfo()
	var biggerFile = new (torrent.File)
	for _,file := range torrentFH.Files() {
		if file.Length()>biggerFile.Length(){biggerFile=file}
	}
	log.Warn(r.Header)
	log.Info("BiggerFile: ",biggerFile.Path())
	rangeReq,okRange := r.Header["Range"]
	fileReader := biggerFile.NewReader()
	if okRange {
		var start,end int64
		fmt.Sscanf(rangeReq[0],"bytes=%d-%d", &start,&end)
		w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d",start,biggerFile.Length(),biggerFile.Length()))
		w.WriteHeader(http.StatusPartialContent)
		fileReader.Seek(start,0)
		log.Warn("start end:",start,end)
		}
	io.Copy(w,fileReader)
}

