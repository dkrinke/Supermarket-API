node('master') {

    try {
      notifyBuild('STARTED')

      withEnv ([ "GOPATH=${env.WORKSPACE}" ])  {
        stage('Check Environment') {
              //For debugging purposes (Check go version and path to working directory)
              sh 'go version'
              sh 'pwd'
        }

        stage('Checkout') {
          //Check out the code
          checkout scm
        }

        stage('Build') {
          //Install gorilla/mux dependency
          sh 'go get github.com/gorilla/mux'
          //Build the supermarketAPI app
          sh 'go install supermarketAPI'
        }

        stage('Test') {
          //Run tests
          sh 'go test ./...'
        }

        //Only execute the below commands if in master branch
        if (env.BRANCH_NAME == 'master') {
          stage('Build Image') {
            //Build the docker image
            sh 'docker build -t dkrinke/supermarketapi:latest . -f ./src/supermarketAPI/Dockerfile'
          }

          stage('Integration Test') {
            sh 'docker run -d -p 127.0.0.1:9000:8080 --name supermarket-api dkrinke/supermarketapi:latest'
            sh 'go test ./src/supermarketAPI -integration'
            sh 'docker stop supermarket-api'
            sh 'docker rm supermarket-api'
          }

          stage('Publish Image') {
            //Publish the docker image to dockerhub
            withDockerRegistry([ credentialsId: "7f19ca19-c670-4382-a759-978c181f242c", url: "" ]) {
              sh 'docker push dkrinke/supermarketapi:latest'
            }
          }

          stage('Deploy') {
            //Copy over the deployment script to the api server
            sh 'scp deploy.sh daniel@192.168.1.123:~'
            //Run the deployment script
            sh 'ssh daniel@192.168.1.123 ./deploy.sh'
          }
        }
      }
    } catch (err) {
        //Set result to failure if any error has occurred
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
