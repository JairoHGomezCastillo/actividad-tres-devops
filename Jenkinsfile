pipeline {
    agent any

    environment {
        IMAGE_NAME = "razas-api"
        IMAGE_TAG = "v${BUILD_NUMBER}"
        CHART_PATH = "helm-chart/razas-api"
        NAMESPACE = "default"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Run Unit Tests') {
            steps {
                sh '''
                  echo "Running Go unit tests..."
                  go test ./... -v
                '''
            }
        }

        stage('Build Docker Image') {
            steps {
                sh '''
                  echo "Building Docker image..."
                  docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .
                '''
            }
        }

        stage('Deploy to Kubernetes with Helm') {
            steps {
                sh '''
                  echo "Deploying ${IMAGE_NAME}:${IMAGE_TAG} to Kubernetes..."
                  kubectl config use-context docker-desktop
                  helm upgrade --install razas-api ${CHART_PATH} \
                    -n ${NAMESPACE} \
                    --set image.repository=${IMAGE_NAME} \
                    --set image.tag=${IMAGE_TAG} \
                    --wait --timeout 120s

                  kubectl get pods -n ${NAMESPACE}
                '''
            }
        }
    }

    post {
        always {
            sh 'docker image prune -f || true'
        }
    }
}
