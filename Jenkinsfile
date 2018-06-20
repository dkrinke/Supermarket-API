node('master') {
    currentBuild.result = "SUCCESS"

    try {
        stage 'Checkout'
          checkout scm

        stage 'Build'
          sh 'go install supermarketAPI'

    } catch (err) {
        currentBuild.result = "FAILURE"
        throw err
    }
}
