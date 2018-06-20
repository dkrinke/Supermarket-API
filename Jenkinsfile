node('master') {
    currentBuild.result = "SUCCESS"

    try {
      withEnv ([ "GOPATH=${env.WORKSPACE}" ])  {
        stage('Check Environment') {
              sh 'go version'
              sh 'pwd'
        }

        stage('Checkout') {
          checkout scm
        }

        stage('Build') {
          sh 'go install supermarketAPI'
        }

        stage('Test') {
          sh 'go test supermarketAPI'
        }

        stage('Build Image') {
         sh 'docker build -t supermarketapi:latest . -f ./src/supermarketAPI/Dockerfile'
        }

        stage('Publish Image') {
          withDockerRegistry([ credentialsId: "7f19ca19-c670-4382-a759-978c181f242c", url: "" ]) {
            sh 'docker push docker pull dkrinke/supermarketapi:latest'
          }
        }

        stage('Deploy') {}
      }
    } catch (err) {
        currentBuild.result = "FAILURE"
        throw err
    }
}
