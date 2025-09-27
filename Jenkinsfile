pipeline {
agent any


environment {
    DOCKER_REGISTRY = "docker.io"
    APP_NAME = "jenkins-go-name-app"
}

tools {
        go 'Go'
    }

stages {
    stage('Clean Workspace') {
      steps {
        deleteDir()
      }
   } 

    stage('Checkout') {
        steps {
            checkout scm
        }
    }

    

    stage('Dependencies') {
        steps {
            dir('fancy-adventure') {
                sh 'go mod tidy'
            }
        }
    }

    stage('Run Tests') {
        steps {
            dir('fancy-adventure') {
                sh 'go test ./... -v'
            }
        }
    }

    
    stage('Docker Build & Push') {
      steps {
        script {
            // Get the current commit SHA
            def commitSha = sh(script: 'git rev-parse --short HEAD', returnStdout: true).trim()

            withCredentials([usernamePassword(
                credentialsId: 'dockerhub-creds',
                usernameVariable: 'DOCKER_USERNAME',
                passwordVariable: 'DOCKER_PASSWORD'
            )]) {
                // Login to Docker
                sh 'echo $DOCKER_PASSWORD | docker login -u $DOCKER_USERNAME --password-stdin'

                // Build and push
                sh """
                  docker build -t $DOCKER_REGISTRY/$DOCKER_USERNAME/$APP_NAME:latest .
                  docker tag $DOCKER_REGISTRY/$DOCKER_USERNAME/$APP_NAME:latest $DOCKER_REGISTRY/$DOCKER_USERNAME/$APP_NAME:${commitSha}
                  docker push $DOCKER_REGISTRY/$DOCKER_USERNAME/$APP_NAME:latest
                  docker push $DOCKER_REGISTRY/$DOCKER_USERNAME/$APP_NAME:${commitSha}
                """
            }
        }
      }
    }

}


}
