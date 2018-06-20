node('master') {
    currentBuild.result = "SUCCESS"

    try {
        stage 'Prepare Environment' {
          // Install the desired Go version
              def root = tool name: 'Go 1.6.2', type: 'go'

          // Export environment variables pointing to the directory where Go was installed
          withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
              sh 'go version'
          }
        }
        stage 'Checkout' {
          checkout scm
        }

        stage 'Build' {
          sh 'go install supermarketAPI'
        }

    } catch (err) {
        currentBuild.result = "FAILURE"
        throw err
    }
}
