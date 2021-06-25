package main

/**
 * This file should be idempotent, meaning all seeds that
 * you run add to this should be resilient to being run
 * multiple times (and should be no-op after the first run).
 */

import (
	"log"

	"github.com/atomicjolt/atomiclti/config"
	"github.com/atomicjolt/atomiclti/model"
	"github.com/atomicjolt/atomiclti/repo"
)

func main() {
	rootRepo := repo.NewRepo()
	ltiAdvantageConfig := config.GetLtiAdvantageConfig()
	serverConfig := config.GetServerConfig()

	application := &model.Application{
		Name:                  "Atomic LTI",
		ClientApplicationName: "atomiclti",
		DefaultConfig:         &config.ApplicationConfig{},
		Description:           "LTI App",
		Key:                   "atomiclti",
		LtiAdvantageConfig:    &ltiAdvantageConfig,
	}

	ltiInstall := &model.LtiInstall{
		ClientID: serverConfig.ClientId,
		Iss:      "https://canvas.instructure.com",
		JwksUrl:  "https://canvas.instructure.com/api/lti/security/jwks",
		OidcUrl:  "https://canvas.instructure.com/api/lti/authorize_redirect",
		TokenUrl: "https://canvas.instructure.com/login/oauth2/token",
	}

	applicationInstance := &model.ApplicationInstance{
		ClientApplicationName: "atomiclti",
		Config:                &config.ApplicationConfig{},
		Description:           "LTI App",
		Key:                   "atomiclti",
	}

	ltiDeployment := &model.LtiDeployment{
		DeploymentID: "TODO",
	}

	jwk := model.NewJwk()

	err := rootRepo.Application.Insert(application)

	if err != nil {
		log.Fatal(err)
	}

	applicationInstance.ApplicationID = application.ID
	err = rootRepo.ApplicationInstance.Insert(applicationInstance)

	if err != nil {
		panic(err)
	}

	ltiInstall.ApplicationID = application.ID
	err = rootRepo.LtiInstall.Insert(ltiInstall)

	if err != nil {
		panic(err)
	}

	jwk.ApplicationID = application.ID
	err = rootRepo.Jwk.Insert(jwk)

	if err != nil {
		panic(err)
	}

	ltiDeployment.ApplicationInstanceID = applicationInstance.ID
	ltiDeployment.LtiInstallID = ltiInstall.ID
	err = rootRepo.LtiDeployment.Insert(ltiDeployment)

	if err != nil {
		panic(err)
	}
}
