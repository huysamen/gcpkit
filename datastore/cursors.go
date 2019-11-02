package datastore

import (
	"encoding/base64"
	"errors"
	"os"

	"github.com/huysamen/gcpkit/utils/encryption"

	"cloud.google.com/go/datastore"
)

func NewCursorCodecs() (encoder func(datastore.Cursor) (string, error), decoder func(string) (datastore.Cursor, error)) {
	key := encryption.NewEncryptionKey(os.Getenv("GCPKIT_DATASTORE_CURSOR_SECRET"))

	enc := func(cursor datastore.Cursor) (string, error) {
		encrypted, err := encryption.Encrypt([]byte(cursor.String()), key)
		if err != nil {
			return "", err
		}

		return base64.URLEncoding.EncodeToString(encrypted), nil
	}

	dec := func(encoded string) (datastore.Cursor, error) {
		if encoded == "" {
			return datastore.Cursor{}, errors.New("cursor not valid")
		}

		decoded, err := base64.URLEncoding.DecodeString(encoded)
		if err != nil {
			return datastore.Cursor{}, err
		}

		decrypted, err := encryption.Decrypt(decoded, key)
		if err != nil {
			return datastore.Cursor{}, err
		}

		cursor, err := datastore.DecodeCursor(string(decrypted))
		if err != nil {
			return datastore.Cursor{}, err
		}

		return cursor, nil
	}

	return enc, dec
}
