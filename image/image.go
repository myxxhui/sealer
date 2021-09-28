// Copyright © 2021 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package image

import "github.com/alibaba/sealer/image/store"

func NewImageService() (Service, error) {
	imageStore, err := store.NewDefaultImageStore()
	if err != nil {
		return nil, err
	}

	return DefaultImageService{imageStore: imageStore}, nil
}

func NewImageMetadataService() (MetadataService, error) {
	imageStore, err := store.NewDefaultImageStore()
	if err != nil {
		return nil, err
	}
	return DefaultImageMetadataService{
		imageStore: imageStore,
	}, nil
}

func NewImageFileService() (FileService, error) {
	layerStore, err := store.NewDefaultLayerStore()
	if err != nil {
		return nil, err
	}

	imageStore, err := store.NewDefaultImageStore()
	if err != nil {
		return nil, err
	}
	return DefaultImageFileService{
		layerStore: layerStore,
		imageStore: imageStore,
	}, nil
}

func NewDeleteImageService(force bool) (Service, error) {
	imageStore, err := store.NewDefaultImageStore()
	if err != nil {
		return nil, err
	}
	return DefaultImageService{
		imageStore:       imageStore,
		ForceDeleteImage: force,
	}, nil
}
