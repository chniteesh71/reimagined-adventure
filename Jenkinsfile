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
    stage('Checkout') {
        steps {
            checkout scm
        }
    }

    stage('Clean Workspace') {
      steps {
        deleteDir()
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

    stage('Lint') {
      steps {
        sh '''
          if ! command -v golangci-lint >/dev/null 2>&1; then
            echo "Installing golangci-lint..."
            curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
              | sh -s -- -b $(go env GOPATH)/bin v1.61.0
          fi
          golangci-lint run ./fancy-adventure
        '''
      }
   }


    stage('Docker Build & Push') {
        when {
            expression { return env.BRANCH_NAME == 'main' }
        }
        steps {
            script {
                withCredentials([usernamePassword(
                    credentialsId: 'dockerhub-creds',
                    usernameVariable: 'DOCKER_USERNAME',
                    passwordVariable: 'DOCKER_PASSWORD'
                )]) {
                    sh 'echo $DOCKER_PASSWORD | docker login -u $DOCKER_USERNAME --password-stdin'
                    sh """
                      docker build -t $DOCKER_REGISTRY/$DOCKER_USERNAME/$APP_NAME:latest .
                      docker tag $DOCKER_REGISTRY/$DOCKER_USERNAME/$APP_NAME:latest $DOCKER_REGISTRY/$DOCKER_USERNAME/$APP_NAME:${GIT_COMMIT}
                      docker push $DOCKER_REGISTRY/$DOCKER_USERNAME/$APP_NAME:latest
                      docker push $DOCKER_REGISTRY/$DOCKER_USERNAME/$APP_NAME:${GIT_COMMIT}
                    """
                }
            }
        }
    }
}


}
