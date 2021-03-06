package cleaning

import (
	"context"
	"fmt"

	"github.com/werf/logboek"
	"github.com/werf/logboek/pkg/style"
	"github.com/werf/logboek/pkg/types"

	"github.com/werf/werf/pkg/stages_manager"
	"github.com/werf/werf/pkg/storage"
)

type CleanupOptions struct {
	ImagesCleanupOptions
	StagesCleanupOptions
}

func Cleanup(ctx context.Context, projectName string, imagesRepo storage.ImagesRepo, storageLockManager storage.LockManager, stagesManager *stages_manager.StagesManager, options CleanupOptions) error {
	m := newCleanupManager(projectName, imagesRepo, stagesManager, options)

	if lock, err := storageLockManager.LockStagesAndImages(ctx, projectName, storage.LockStagesAndImagesOptions{GetOrCreateImagesOnly: false}); err != nil {
		return fmt.Errorf("unable to lock stages and images: %s", err)
	} else {
		defer storageLockManager.Unlock(ctx, lock)
	}

	if err := logboek.Context(ctx).Default().LogProcess("Running images cleanup").
		Options(func(options types.LogProcessOptionsInterface) {
			options.Style(style.Highlight())
		}).
		DoError(func() error {
			return m.imagesCleanupManager.run(ctx)
		}); err != nil {
		return err
	}

	repoImages := m.imagesCleanupManager.getImageRepoImageList()
	m.stagesCleanupManager.setImagesRepoImageList(flattenRepoImages(repoImages))

	if err := logboek.Context(ctx).Default().LogProcess("Running stages cleanup").
		Options(func(options types.LogProcessOptionsInterface) {
			options.Style(style.Highlight())
		}).
		DoError(func() error {
			return m.stagesCleanupManager.run(ctx)
		}); err != nil {
		return err
	}

	return nil
}

func newCleanupManager(projectName string, imagesRepo storage.ImagesRepo, stagesManager *stages_manager.StagesManager, options CleanupOptions) *cleanupManager {
	return &cleanupManager{
		imagesCleanupManager: newImagesCleanupManager(projectName, imagesRepo, stagesManager, options.ImagesCleanupOptions),
		stagesCleanupManager: newStagesCleanupManager(projectName, imagesRepo, stagesManager, options.StagesCleanupOptions),
	}
}

type cleanupManager struct {
	*imagesCleanupManager
	*stagesCleanupManager
}
