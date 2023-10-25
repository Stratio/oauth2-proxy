@Library('libpipelines') _

hose {
    EMAIL = 'eos@stratio.com'
    BUILDTOOL_IMAGE = 'golang:1.19'
    BUILDTOOL = 'make'
    DEVTIMEOUT = 30
    ANCHORE_POLICY = "production"
    VERSIONING_TYPE = 'stratioVersion-3-3'
    UPSTREAM_VERSION = '7.5.1'
    GRYPE_TEST = false

    DEV = { config ->
        doUT(conf: config, parameters: "GOCACHE=/tmp")
        doDocker(conf: config, image: 'oauth2-proxy', buildargs:"DOCKER_BUILD_PLATFORM=linux/amd64")
    }
}
