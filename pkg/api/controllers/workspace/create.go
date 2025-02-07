// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package workspace

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/pkg/api/util"
	"github.com/daytonaio/daytona/pkg/models"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/daytonaio/daytona/pkg/services"
	"github.com/gin-gonic/gin"
)

// CreateWorkspace 			godoc
//
//	@Tags			workspace
//	@Summary		Create a workspace
//	@Description	Create a workspace
//	@Param			workspace	body	CreateWorkspaceDTO	true	"Create workspace"
//	@Produce		json
//	@Success		200	{object}	WorkspaceDTO
//	@Router			/workspace [post]
//
//	@id				CreateWorkspace
func CreateWorkspace(ctx *gin.Context) {
	var createWorkspaceReq services.CreateWorkspaceDTO
	err := ctx.BindJSON(&createWorkspaceReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err))
		return
	}

	server := server.GetInstance(nil)

	w, err := server.WorkspaceService.Create(ctx.Request.Context(), createWorkspaceReq)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to create workspace: %w", err))
		return
	}

	apiKeyType, ok := ctx.Get("apiKeyType")
	if !ok || apiKeyType == models.ApiKeyTypeClient {
		util.HideDaytonaEnvVars(&w.EnvVars)
		util.HideDaytonaEnvVars(&w.Target.EnvVars)
		w.ApiKey = ""
		w.Target.ApiKey = ""
	}

	ctx.JSON(200, w)
}
