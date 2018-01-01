// Copyright 2017 visualfc. All rights reserved.

package tk

import (
	"errors"
	"fmt"
	"image"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/visualfc/go-tk/tk/interp"
)

type Image struct {
	id    string
	photo *interp.Photo
}

func (i *Image) Id() string {
	return i.id
}

type image_option struct {
	key   string
	value interface{}
}

func ImageOptId(id string) *image_option {
	return &image_option{"id", id}
}

func imageOptGif(file string) *image_option {
	return &image_option{"file", file}
}

func LoadImage(file string, options ...*image_option) (*Image, error) {
	if file == "" {
		return nil, os.ErrInvalid
	}
	var fileImage image.Image
	if filepath.Ext(file) == ".gif" {
		options = append(options, imageOptGif(file))
	} else {
		file, err := os.Open(file)
		if err == nil {
			im, _, err := image.Decode(file)
			file.Close()
			if err != nil {
				return nil, err
			}
			fileImage = im
		}
	}
	im := NewImage(options...)
	if im == nil {
		return nil, errors.New("NewImage failed")
	}
	if fileImage != nil {
		im.SetImage(fileImage)
	}
	return im, nil
}

func NewImage(options ...*image_option) *Image {
	var iid string
	var optList []string
	for _, opt := range options {
		if opt == nil {
			continue
		}
		if opt.key == "id" {
			if v, ok := opt.value.(string); ok {
				iid = v
			}
			continue
		}
		optList = append(optList, fmt.Sprintf("-%v {%v}", opt.key, opt.value))
	}
	if iid == "" {
		iid = MakeImageId()
	}
	script := fmt.Sprintf("image create photo %v", iid)
	if len(optList) > 0 {
		script += " " + strings.Join(optList, " ")
	}
	err := eval(script)
	if err != nil {
		return nil
	}
	photo := interp.FindPhoto(mainInterp, iid)
	if photo == nil {
		return nil
	}
	return &Image{iid, photo}
}

func (i *Image) IsValid() bool {
	return i.id != "" && i.photo != nil
}

func (i *Image) SetImage(img image.Image) *Image {
	err := i.photo.PutImage(img)
	if err != nil {
		dumpError(err)
	}
	return i
}

func (i *Image) SetZoomedImage(img image.Image, zoomX, zoomY, subsampleX, subsampleY int) *Image {
	err := i.photo.PutZoomedImage(img, zoomX, zoomY, subsampleX, subsampleY)
	if err != nil {
		dumpError(err)
	}
	return i
}

func (i *Image) ToImage() image.Image {
	return i.photo.ToImage()
}

func (i *Image) Blank() *Image {
	i.photo.Blank()
	return i
}

func (i *Image) SizeN() (width int, height int) {
	return i.photo.Size()
}

func (i *Image) Size() Size {
	w, h := i.SizeN()
	return Size{w, h}
}

func (i *Image) SetSizeN(width int, height int) *Image {
	err := i.photo.SetSize(width, height)
	if err != nil {
		dumpError(err)
	}
	return i
}

func (i *Image) SetSize(sz Size) *Image {
	return i.SetSizeN(sz.Width, sz.Height)
}
