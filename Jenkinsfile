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
            sshagent (credentials: ['314012a8-120d-43d5-b0ef-0a299301fae8']) {
              sh 'ssh -o StrictHostKeyChecking=no -l daniel 192.168.1.123 uname -a'
              sh 'sudo docker stop supermarket-api'
              sh 'sudo docker run -d -it --rm -p 0.0.0.0:8080:8080 --name supermarket-api dkrinke/supermarketapi:latest'
            }
          }
        }

        // githubNotify status: "SUCCESS", credentialsId: "680e5762-840b-46ea-883c-c7bb0310a357	", account: "dkrinke", repo: "Supermarket-API"

      }
    } catch (err) {
        currentBuild.result = "FAILURE"
        // githubNotify status: "FAILURE", credentialsId: "680e5762-840b-46ea-883c-c7bb0310a357	", account: "dkrinke", repo: "Supermarket-API"
        throw err
    }
}
