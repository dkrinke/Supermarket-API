node('master') {
    currentBuild.result = "SUCCESS"

    try {
        stage('Check Environment') {
              sh 'go version'
              sh 'pwd'
        }

        stage('Checkout') {
          checkout scm
        }

        withEnv ([ "GOPATH=${env.WORKSPACE}" ])  {
          stage('Build') {
            sh 'go install supermarketAPI'
          }

          stage('Test') {
            sh 'go test supermarketAPI'
          }
        }
        
        withEnv ([ "GOPATH=${env.WORKSPACE}/src/supermarketAPI" ]) {
          stage('Build Image') {
           sh 'pwd'
           sh 'ls -a'
           sh 'docker build -t supermarketapi:latest .'
           sh 'pwd'
           sh 'ls -a'
          }
        }
      }
    } catch (err) {
        currentBuild.result = "FAILURE"
        throw err
    }
}
