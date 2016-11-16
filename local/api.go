package local

import (
	"errors"
	"os"

	"golang.org/x/net/context"

	api_api "github.com/james-nesbitt/kraut-api/api"
	api_builder "github.com/james-nesbitt/kraut-api/builder"
	handlers_bytesource "github.com/james-nesbitt/kraut-handlers/bytesource"
	handlers_local "github.com/james-nesbitt/kraut-handlers/local"
	handlers_null "github.com/james-nesbitt/kraut-handlers/null"
)

const (
	WUNDERTOOLS_PROJECT_CONF_FOLDER = ".kraut" // If the project has existing setitngs, they will be in this subfolder, somewhere up the file tree.
	WUNDERTOOLS_USER_CONF_SUBPATH   = "kraut"  // If the user has user-scope config, they will be in this subfolder
)

/**
 * Build a local API, by scanning for project settings based on the
 * path.  First a number of "conf" folders are determinged, and these
 * are used to build the localAPI.
 */

// Construct a LocalAPI by checking some paths for the current user.
func MakeLocalAPI() (api_api.API, error) {
	var err error

	workingDir, _ := os.Getwd()
	settings := handlers_local.LocalAPISettings{
		BytesourceFileSettings: handlers_bytesource.BytesourceFileSettings{
			ExecPath:    workingDir,
			ConfigPaths: &handlers_bytesource.Paths{},
		},
		Context: context.Background(),
	}

	// Discover paths for the user like ~ and ~/.config/wundertools
	DiscoverUserPaths(&settings)
	DiscoverProjectPaths(&settings)

	/**
	 * We could here add more paths for settings.ConfigPaths, for
	 * configurations of a higher priority.  For example, a feature
	 * or environment concept might want to override user and
	 * project level confs
	 */

	/**
	 * Now that we have local settings, let's start to build our API
	 *
	 * We will build it using a BuilderAPI, and adding the local
	 * Builder, which will be used at a minimum for a ConfigWrapper,
	 * to determine how to build the rest of the project
	 */
	localApi := api_builder.BuilderAPI{}
	localApi.AddBuilder(handlers_local.New_LocalBuilder(settings))

	/**
	 * If we have discovered that there is no local project folder,
	 * then we will enable a minimum API, which can be used to create
	 * a local folder
	 */
	if settings.ProjectDoesntExist {

		// allow local project operations, which could be used to build a project
		localApi.ActivateBuilder("local", *api_builder.New_Implementations([]string{"project"}), nil)

		// Use null wrappers for those handlers that we don't have (to play safe)
		localApi.AddBuilder(handlers_null.New_NullBuilder())
		localApi.ActivateBuilder("null", *api_builder.New_Implementations([]string{"config", "seting", "command"}), nil)

		err = errors.New("No project found.")

	} else {
		/**
		 * The automated build is complex enough that it deserves
		 * it's own method
		 */

		LocalBuild(&localApi)

	}

	return api_api.API(&localApi), err
}