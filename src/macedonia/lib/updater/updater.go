package updater

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"macedonia/lib/setting"

	"google.golang.org/appengine/log"
	"gopkg.in/src-d/go-billy.v3/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/client"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	gioutil "gopkg.in/src-d/go-git.v4/utils/ioutil"
)

const (
	repoURL = "ssh://git@github.com/delphinus/homebrew-macvim-kaoriya"
)

// Updater updates cask for macvim-kaoriya
func Updater(ctx context.Context) (err error) {
	var pkey, pw []byte
	if _, err := base64.StdEncoding.Decode(pkey, setting.Setting.PrivateKey); err != nil {
		return fmt.Errorf("error in Decode: %v", err)
	}
	if _, err := base64.StdEncoding.Decode(pw, setting.Setting.Password); err != nil {
		return fmt.Errorf("error in Decode: %v", err)
	}
	auth, err := ssh.NewPublicKeys("delphinus", pkey, pw)
	if err != nil {
		return fmt.Errorf("error in NewPublicKeys: %v", err)
	}
	client.InstallProtocol("ssh", ssh.NewClient(&ssh.ClientConfig{
		Auth: {auth},
	}))
	fs := memfs.New()
	r, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		return fmt.Errorf("error in Clone: %v", err)
	}
	f, err := fs.Open("Casks/macvim-kaoriya.rb")
	if err != nil {
		return fmt.Errorf("error in Open: %v", err)
	}
	defer gioutil.CheckClose(f, &err)
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("error in ReadAll: %v", err)
	}
	log.Infof(ctx, "content: %s", content)
	return nil
}
