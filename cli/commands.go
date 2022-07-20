package cli

import (
	"context"
	"fmt"
	"git_with_filecoin/util"
	gocid "github.com/ipfs/go-cid"
	"github.com/web3-storage/go-w3s-client"
	"io/fs"
	"log"
	"os"
	"os/exec"
)

func (cli *CLI) commit(token string) {
	if util.IsGitRepo() {
		fmt.Println("OK, it is within a git repo.")
		userHome, err := os.UserHomeDir()
		if err != nil {
			log.Panicln(err)
		}
		home := userHome + "/.git_with_filecoin"
		if !util.PathExist(home) {
			err = os.Mkdir(home, 0755)
			if err != nil {
				log.Panic(err)
			}
		}
		patchLoc := home + "/patch"
		cmd := exec.Command("git", "diff>"+patchLoc)
		resp, err := cmd.Output()
		fmt.Println(string(resp))
		//err = cmd.Start()
		//if err != nil {
		//	log.Panic(err)
		//}
		//err = cmd.Wait()
		//if err != nil {
		//	log.Panic(err)
		//}
		util.PutFile(token, patchLoc)
	} else {
		fmt.Println("NO! It is not a git repo!")
	}
}

func (cli *CLI) apply(token string, cid string) {
	c, _ := w3s.NewClient(w3s.WithToken(token))
	decodedCid, _ := gocid.Parse(cid) //gocid.Decode(cid)
	res, err := c.Get(context.Background(), decodedCid)
	if err != nil {
		log.Panic(err)
	}
	// res is a http.Response with an extra method for reading IPFS UnixFS files!
	f, fsys, err := res.Files()
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
	fmt.Println("QAQ")
	//patch, _ := fsys.Open("patch")
	patch, err := fs.ReadFile(fsys, "/patch")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Read successful.")
	fmt.Println(string(patch))

	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Panicln(err)
	}
	home := userHome + "/.git_with_filecoin"
	if !util.PathExist(home) {
		err = os.Mkdir(home, 0755)
		if err != nil {
			log.Panic(err)
		}
	}
	patchLoc := home + "/patch"
	err = os.WriteFile(patchLoc, patch, 0644)
	if err != nil {
		log.Panic(err)
	}
	cmd := exec.Command("git", "apply", patchLoc)
	resp, err := cmd.Output()
	fmt.Println(string(resp))
}
