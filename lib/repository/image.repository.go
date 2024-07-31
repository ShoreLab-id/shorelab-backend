package repository

import (
	"io"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func (r *Repository) GetImageList() ([]string, error) {
	obj := r.db.StorageBucket.Objects(r.ctx, nil)
	var blobList []string

	for {
		attr, err := obj.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		blobList = append(blobList, attr.Name)
	}
	return blobList, nil
}

func (r *Repository) GetImage(img string) ([]byte, string, error) {
	obj := r.db.StorageBucket.Object(img)
	attr, err := obj.Attrs(r.ctx)
	if err == storage.ErrObjectNotExist {
		log.Default().Println(err.Error())
		return nil, "", storage.ErrObjectNotExist
	}
	if err != nil {
		log.Default().Println(err.Error())
		return nil, "", err
	}

	reader, err := obj.NewReader(r.ctx)
	if err != nil {
		log.Default().Println(err.Error())
		return nil, "", err
	}
	defer reader.Close()

	p, err := io.ReadAll(reader)
	if err != nil {
		log.Default().Println(err.Error())
		return nil, "", err
	}

	return p, attr.ContentType, nil
}
