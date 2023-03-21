package cliutil

import (
	"bytes"
	"io"
	"os"

	"github.com/aperturerobotics/bifrost/keypem"
	"github.com/aperturerobotics/bifrost/peer"
	"github.com/aperturerobotics/bifrost/peer/ssh"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh"
)

// RunGeneratePrivate runs the generate-private util command.
func (a *UtilArgs) RunGeneratePrivate(_ *cli.Context) error {
	npeer, err := peer.NewPeer(nil)
	if err != nil {
		return err
	}

	priv, err := npeer.GetPrivKey(a.GetContext())
	if err != nil {
		return err
	}

	pemd, err := keypem.MarshalPrivKeyPem(priv)
	if err != nil {
		return err
	}
	err = writeIfNotExists(a.OutPath, bytes.NewReader(pemd))
	if err != nil {
		return err
	}
	le := a.GetLogger()
	le.Debugf("generated private key: %s", npeer.GetPeerID().Pretty())
	return nil
}

// RunDerivePublic derives the public key from a private pem.
func (a *UtilArgs) RunDerivePublic(_ *cli.Context) error {
	rp, err := a.readInputFilePrivKey()
	if err != nil {
		return err
	}
	pemd, err := keypem.MarshalPubKeyPem(rp.GetPubKey())
	if err != nil {
		return err
	}
	err = writeIfNotExists(a.OutPath, bytes.NewReader(pemd))
	if err != nil {
		return err
	}
	return nil
}

// RunDerivePublic derives the ssh public key from a private or public pem.
func (a *UtilArgs) RunDeriveSshPublic(_ *cli.Context) error {
	rp, err := a.readInputFilePubKey()
	if err != nil {
		return err
	}
	pkey, err := peer_ssh.NewPublicKey(rp.GetPubKey())
	if err != nil {
		return err
	}
	dat := ssh.MarshalAuthorizedKey(pkey)
	err = writeIfNotExists(a.OutPath, bytes.NewReader(dat))
	if err != nil {
		return err
	}
	return nil
}

// readInputFile reads the input file path or stdin.
func (a *UtilArgs) readInputFile() ([]byte, error) {
	if fp := a.FilePath; fp != "" {
		return os.ReadFile(fp)
	}

	return io.ReadAll(os.Stdin)
}

// readInputFilePrivKey reads the input file path or stdin.
func (a *UtilArgs) readInputFilePrivKey() (peer.Peer, error) {
	dat, err := a.readInputFile()
	if err != nil {
		return nil, err
	}

	key, err := keypem.ParsePrivKeyPem(dat)
	if err != nil {
		return nil, err
	}

	le := a.GetLogger()
	npeer, err := peer.NewPeer(key)
	if err != nil {
		return nil, err
	}
	le.Debugf("loaded private key: %s", npeer.GetPeerID().Pretty())
	return npeer, nil
}

// readInputFilePubKey reads the input file path or stdin.
func (a *UtilArgs) readInputFilePubKey() (peer.Peer, error) {
	dat, err := a.readInputFile()
	if err != nil {
		return nil, err
	}

	key, err := keypem.ParsePubKeyPem(dat)
	if err != nil {
		return nil, err
	}

	le := a.GetLogger()
	npeer, err := peer.NewPeerWithPubKey(key)
	if err != nil {
		return nil, err
	}
	le.Debugf("loaded public key: %s", npeer.GetPeerID().Pretty())
	return npeer, nil
}

func writeIfNotExists(outPath string, input io.Reader) error {
	var of *os.File
	var out io.Writer
	if outPath != "" {
		_, err := os.Stat(outPath)
		if !os.IsNotExist(err) {
			return errors.Wrap(os.ErrExist, outPath)
		}
		of, err = os.OpenFile(outPath, os.O_CREATE|os.O_RDWR, 0600)
		if err != nil {
			return err
		}
		out = of
		defer of.Close()
		if pos, err := of.Seek(0, io.SeekEnd); err != nil || pos != 0 {
			if err == nil {
				// file must have existed
				return errors.Wrap(os.ErrExist, outPath)
			}
			return err
		}
	} else {
		out = os.Stdout
	}
	if _, err := io.Copy(out, input); err != nil {
		return err
	}
	if of != nil {
		return of.Close()
	}
	return nil
}
