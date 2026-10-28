// Copyright (c) 2025-2026, Audionut and the autobrr contributors.
// SPDX-License-Identifier: GPL-2.0-or-later

//go:build !linux && !darwin && !windows

package webserver

import (
	"context"

	"github.com/autobrr/upbrr/pkg/api"
)

func browseVolumeKey(_ context.Context, _ string) string {
	return ""
}

// browseWindowsRoots is only meaningful on Windows.
func browseWindowsRoots(mode string) api.BrowseDirectoryResponse {
	return api.BrowseDirectoryResponse{
		CurrentPath: "",
		ParentPath:  "",
		Mode:        mode,
		Entries:     nil,
		Scope:       "roots",
	}
}
