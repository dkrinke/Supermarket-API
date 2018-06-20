node('master') {
    currentBuild.result = "SUCCESS"

    try {
      withEnv ([ "GOPATH=${env.WORKSPACE}/src/supermarketAPI" ])  {
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
         sh 'pwd'
         sh 'ls -a'
         sh 'docker build -t supermarketapi:latest .'
         sh 'pwd'
         sh 'ls -a'
        }
      }
    } catch (err) {
        currentBuild.result = "FAILURE"
        throw err
    }
}
