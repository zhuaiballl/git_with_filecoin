package util

import (
	"context"
	"fmt"
	gocid "github.com/ipfs/go-cid"
	"github.com/web3-storage/go-w3s-client"
	"io/fs"
	"log"
	"os"
)

func PutFile(token string, loc string) {
	c, _ := w3s.NewClient(w3s.WithToken(token))
	fmt.Println("Puting", loc, "to web3.storage...")
	f, _ := os.Open(loc)
	cid, _ := c.Put(context.Background(), f)
	fmt.Printf("https://%v.ipfs.dweb.link\n", cid)
}

func GetFile(token string, cid string) {
	c, _ := w3s.NewClient(w3s.WithToken(token))
	decodedCid, _ := gocid.Parse(cid) //gocid.Decode(cid)
	res, err := c.Get(context.Background(), decodedCid)
	if err != nil {
		log.Panic(err)
	}
	// res is a http.Response with an extra method for reading IPFS UnixFS files!
	f, _, err := res.Files()
	if err != nil {
		log.Panic(err)
	}
	// List directory entries
	if d, ok := f.(fs.ReadDirFile); ok {
		ents, _ := d.ReadDir(0)
		for _, ent := range ents {
			fmt.Println(ent.Name())
		}
	}

	// Walk whole directory contents (including nested directories)
	//fs.WalkDir(fsys, "/", func(path string, d fs.DirEntry, err error) error {
	//	info, _ := d.Info()
	//	fmt.Printf("%s (%d bytes)\n", path, info.Size())
	//	return err
	//})
}
