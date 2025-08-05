@Library('libpipelines') _

hose {
    EMAIL = 'eos@stratio.com'
    BUILDTOOL_IMAGE = 'golang:1.24'
    BUILDTOOL = 'make'
    DEVTIMEOUT = 30
    ANCHORE_POLICY = "production"
    VERSIONING_TYPE = 'stratioVersion-3-3'
    UPSTREAM_VERSION = '7.11.0'
    GRYPE_TEST = false

    DEV = { config ->
        doUT(conf: config, parameters: "GOCACHE=/tmp")
        doDocker(
            conf: config,
            image: 'oauth2-proxy',
            buildargs: [
                "BUILDPLATFORM=linux/amd64",
            ]
        )
    }
}
