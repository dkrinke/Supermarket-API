node('master') {
    currentBuild.result = "SUCCESS"

    try {
        stage 'Checkout'
          checkout scm
    } catch (err) {
        currentBuild.result = "FAILURE"
        throw err
    }
}
