package libcompose

import (
	"errors"
	"io"

	"golang.org/x/net/context"

	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Some usefull Base classes used by other libcompose operations
 * and Properties
 */

// A base libcompose operation with Properties for staying attached
type BaseLibcomposeStayAttachedOperation struct {
	properties *operation.Properties
}

// Provide static Properties for the operation
func (base *BaseLibcomposeStayAttachedOperation) Properties() *operation.Properties {
	if base.properties == nil {
		newProperties := &operation.Properties{}

		newProperties.Add(operation.Property(&LibcomposeAttachFollowProperty{}))

		base.properties = newProperties
	}
	return base.properties
}

// A handoff function to make a base orchestration operation, which is really just a lot of linear code.
// @NOTE this needs the "config.get" operation to already be available
func New_BaseLibcomposeOrchestrateNameFilesOperation(projectName string, dockerComposeFiles []string, runContext context.Context, outputWriter io.Writer, errorWriter io.Writer) (BaseLibcomposeOrchestrateNameFilesOperation, operation.Result) {
	result := operation.BaseResult{}
	result.Set(true, nil)

	// This Base operations will be at the root of all of the libCompose operations
	baseLibcomposeOrchestrate := BaseLibcomposeOrchestrateNameFilesOperation{}
	orchestrateProperties := baseLibcomposeOrchestrate.Properties()

	// Set a project name
	if projectNameConf, found := orchestrateProperties.Get(OPERATION_PROPERTY_LIBCOMPOSE_PROJECTNAME); found {
		if !projectNameConf.Set(projectName) {
			result.Set(false, []error{errors.New("Could not set base libCompose project name.  Config set error on base Orchestration operation")})
		}
	} else {
		result.Set(false, []error{errors.New("Could not set base libCompose project name.  Config value not found on base Orchestration operation")})
	}

	// Add project docker-compose yml files
	if projectComposeFilesConf, found := orchestrateProperties.Get(OPERATION_PROPERTY_LIBCOMPOSE_COMPOSEFILES); found {
		if !projectComposeFilesConf.Set(dockerComposeFiles) {
			result.Set(false, []error{errors.New("Could not set base libcompose docker-compose file conf.  Config set error on base Orchestration operation")})
		}
	} else {
		result.Set(false, []error{errors.New("Could not set base libcompose docker-compose file conf.  Config not found on base Orchestration operation")})
	}
	// Add project context
	if projectContextConf, found := orchestrateProperties.Get(OPERATION_PROPERTY_LIBCOMPOSE_CONTEXT); found {
		if !projectContextConf.Set(runContext) {
			result.Set(false, []error{errors.New("Could not set base libcompose net context.  Config set error on base Orchestration operation")})
		}
	} else {
		result.Set(false, []error{errors.New("Could not set base libcompose net context.  Config not found on base Orchestration operation")})
	}
	// Add Stdout as an output writer
	if projectOutputConf, found := orchestrateProperties.Get(OPERATION_PROPERTY_LIBCOMPOSE_OUTPUT); found {
		if !projectOutputConf.Set(outputWriter) {
			result.Set(false, []error{errors.New("Could not set base libcompose output handler.  Config set error on base Orchestration operation")})
		}
	} else {
		result.Set(false, []error{errors.New("Could not set base libcompose output handler.  Config not found on base Orchestration operation")})
	}
	if projectErrorConf, found := orchestrateProperties.Get(OPERATION_PROPERTY_LIBCOMPOSE_ERROR); found {
		if !projectErrorConf.Set(errorWriter) {
			result.Set(false, []error{errors.New("Could not set base libcompose error handler.  Config set error on base Orchestration operation")})
		}
	} else {
		result.Set(false, []error{errors.New("Could not set base libcompose error handler.  Config not found on base Orchestration operation")})
	}

	return baseLibcomposeOrchestrate, operation.Result(&result)
}

// A base libcompose operation with Properties for project-name, and yml files
type BaseLibcomposeOrchestrateNameFilesOperation struct {
	properties *operation.Properties
}

// Provide static Properties for the operation
func (base *BaseLibcomposeOrchestrateNameFilesOperation) Properties() *operation.Properties {
	if base.properties == nil {
		newProperties := &operation.Properties{}

		newProperties.Add(operation.Property(&LibcomposeProjectnameProperty{}))
		newProperties.Add(operation.Property(&LibcomposeComposefilesProperty{}))
		newProperties.Add(operation.Property(&LibcomposeContextProperty{}))

		newProperties.Add(operation.Property(&LibcomposeOutputProperty{}))
		newProperties.Add(operation.Property(&LibcomposeErrorProperty{}))

		base.properties = newProperties
	}
	return base.properties
}

// Base Up operation
type BaseLibcomposeOrchestrateUpOperation struct {
	properties *operation.Properties
}

// Provide static Properties for the operation
func (base *BaseLibcomposeOrchestrateUpOperation) Properties() *operation.Properties {
	if base.properties == nil {
		newProperties := &operation.Properties{}

		newProperties.Add(operation.Property(&LibcomposeOptionsUpProperty{}))

		base.properties = newProperties
	}
	return base.properties
}

// Base Down operation
type BaseLibcomposeOrchestrateDownOperation struct {
	properties *operation.Properties
}

// Provide static Properties for the operation
func (base *BaseLibcomposeOrchestrateDownOperation) Properties() *operation.Properties {
	if base.properties == nil {
		newProperties := &operation.Properties{}

		newProperties.Add(operation.Property(&LibcomposeOptionsDownProperty{}))

		base.properties = newProperties
	}
	return base.properties
}
