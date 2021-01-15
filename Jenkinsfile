@Library('dst-shared@master') _

dockerBuildPipeline {
        githubPushRepo = "Cray-HPE/hms-bss"
        repository = "cray"
        imagePrefix = "cray"
        app = "bss"
        name = "hms-bss"
        description = "Cray boot script service"
        dockerfile = "Dockerfile"
        slackNotification = ["", "", false, false, true, true]
        product = "csm"
}
