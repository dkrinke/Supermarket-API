node('master') {

    try {
      notifyBuild('STARTED')

      withEnv ([ "GOPATH=${env.WORKSPACE}" ])  {
        stage('Check Environment') {
              sh 'go version'
              sh 'pwd'
        }

        stage('Checkout') {
          checkout scm
        }

        stage('Build') {
          sh 'go get github.com/gorilla/mux'
          sh 'go install supermarketAPI'
        }

        stage('Test') {
          sh 'go test ./...'
        }

        if (env.BRANCH_NAME == 'master') {
          stage('Build Image') {
           sh 'docker build -t dkrinke/supermarketapi:latest . -f ./src/supermarketAPI/Dockerfile'
          }

          stage('Publish Image') {
            withDockerRegistry([ credentialsId: "7f19ca19-c670-4382-a759-978c181f242c", url: "" ]) {
              sh 'docker push dkrinke/supermarketapi:latest'
            }
          }

          stage('Deploy') {
            sh 'scp deploy.sh daniel@192.168.1.123:~'
            sh 'ssh daniel@192.168.1.123 ./deploy.sh'
          }
        }
      }
    } catch (err) {
        currentBuild.result = "FAILURE"
        throw err
    } finally {
        // Success or failure, always send notifications
        notifyBuild(currentBuild.result)
    }
}


def notifyBuild(String buildStatus = 'STARTED') {
  // build status of null means successful
  buildStatus =  buildStatus ?: 'SUCCESSFUL'

  // Default values
  def colorName = 'RED'
  def colorCode = '#FF0000'
  def subject = "${buildStatus}: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'"
  def summary = "${subject} (${env.BUILD_URL})"

  // Override default values based on build status
  if (buildStatus == 'STARTED') {
    color = 'YELLOW'
    colorCode = '#FFFF00'
  } else if (buildStatus == 'SUCCESSFUL') {
    color = 'GREEN'
    colorCode = '#00FF00'
  } else {
    color = 'RED'
    colorCode = '#FF0000'
  }

  // Send notifications
  slackSend (color: colorCode, message: summary)
}
