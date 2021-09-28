package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"

	"github.com/alibaba/sealer/common"

	"github.com/alibaba/sealer/utils"
	"github.com/alibaba/sealer/utils/compress"
	"github.com/opencontainers/go-digest"
)

const emptySHA256TarDigest = "sha256:4f4fb700ef54461cfa02571ae0db9a0dc1e0cdb5577484a6d75e68dc38e8acc1"

type SHA256 struct {
}

func (sha SHA256) CheckSum(reader io.Reader) (digest.Digest, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return "", err
	}
	dig := digest.NewDigestFromEncoded(digest.SHA256, hex.EncodeToString(hash.Sum(nil)))
	return dig, nil
}

func (sha SHA256) TarCheckSum(src string) (*os.File, digest.Digest, error) {
	file, err := compress.RootDirNotIncluded(nil, src)
	if err != nil {
		return nil, "", err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, "", err
	}

	dig, err := sha.CheckSum(file)
	if err != nil {
		return nil, "", err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, "", err
	}
	return file, dig, nil
}

func CheckSumAndPlaceLayer(src string) (digest.Digest, error) {
	sha := SHA256{}
	file, dig, err := sha.TarCheckSum(src)
	if err != nil {
		return "", err
	}

	defer utils.CleanFile(file)
	err = compress.Decompress(file, filepath.Join(common.DefaultLayerDir, dig.Hex()))
	if err != nil {
		return "", err
	}

	return dig, nil
}

func (sha SHA256) EmptyDigest() digest.Digest {
	return emptySHA256TarDigest
}
