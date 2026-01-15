@Library('libpipelines') _

hose {
    EMAIL = 'eos@stratio.com'
    BUILDTOOL_IMAGE = 'golang:1.25'
    BUILDTOOL = 'make'
    DEVTIMEOUT = 30
    ANCHORE_POLICY = "production"
    VERSIONING_TYPE = 'stratioVersion-3-3'
    UPSTREAM_VERSION = '7.13.0'
    GRYPE_TEST = true

    DEV = { config ->
        doUT(conf: config, parameters: "GOCACHE=/tmp")
        doDocker(
            conf: config,
            image: 'oauth2-proxy',
            buildargs: [
                "BUILD_IMAGE=golang:1.25-bookworm",
                "RUNTIME_IMAGE=distroless/static:nonroot",
                "BUILDPLATFORM=linux/amd64",
            ]
        )
        doPushDockerECR(conf: config,AWS_CREDENTIALS_ID: 'AWS_CREDENTIALS_ECR_TEST',AWS_REGION: 'us-east-1')
    }
}
