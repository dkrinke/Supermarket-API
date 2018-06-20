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
      }
    } catch (err) {
        currentBuild.result = "FAILURE"
        throw err
    }
}
