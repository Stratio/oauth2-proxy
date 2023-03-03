@Library('libpipelines') _

hose {
    EMAIL = 'eos@stratio.com'
    BUILDTOOL_IMAGE = 'golang:1.16'
    BUILDTOOL = 'make'
    DEVTIMEOUT = 30
    ANCHORE_POLICY = "production"
    VERSIONING_TYPE = 'stratioVersion-3-3'
    UPSTREAM_VERSION = '7.1.2'
    ANCHORE_TEST = true

    DEV = { config ->
        doUT(conf: config, parameters: "GOCACHE=/tmp")
        doDocker(conf: config, image: 'oauth2-proxy')
    }
}
