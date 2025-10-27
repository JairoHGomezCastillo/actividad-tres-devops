pipeline {
    agent any

    options {
        disableConcurrentBuilds()
      }

    environment {
        IMAGE_NAME = "razas-api"
        IMAGE_TAG = "v${BUILD_NUMBER}"
        CHART_PATH = "helm-chart/razas-api"
        NAMESPACE = "default"
        DOCKERHUB_REPO = "jahudev/razas-api"
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

        stage('Push Docker Image') {
          steps {
            withCredentials([usernamePassword(credentialsId: 'dockerhub-creds', usernameVariable: 'DH_USER', passwordVariable: 'DH_PASS')]) {
              sh """
                echo "\nTagging and pushing to Docker Hub..."
                docker tag ${IMAGE_NAME}:${IMAGE_TAG} ${DOCKERHUB_REPO}:${IMAGE_TAG}
                echo '${DH_PASS}' | docker login -u '${DH_USER}' --password-stdin
                docker push ${DOCKERHUB_REPO}:${IMAGE_TAG}
              """
            }
          }
        }

        // Este stage actualiza el Helm chart para que ArgoCD lo detecte y haga sync automático
        stage('Update Helm values and push (GitOps for ArgoCD)') {
          steps {
            withCredentials([usernamePassword(credentialsId: 'github-creds', usernameVariable: 'GIT_USER', passwordVariable: 'GIT_TOKEN')]) {
              sh """
                set -eux

                # Asegura estar en main con lo último
                git fetch origin main
                git checkout main
                git pull --rebase origin main

                echo "Updating Helm values.yaml tag -> ${IMAGE_TAG}"
                sed -i 's/^\\s*tag:\\s*.*/  tag: ${IMAGE_TAG}/' ${CHART_PATH}/values.yaml

                git config user.email "jenkins@local"
                git config user.name "Jenkins CI"

                git add ${CHART_PATH}/values.yaml
                git commit -m "ci: bump image tag to ${IMAGE_TAG}"

                # Cambia la URL de origin para incluir credenciales y poder hacer push
                git remote set-url origin https://${GIT_USER}:${GIT_TOKEN}@github.com/JairoHGomezCastillo/actividad-tres-devops.git
                git push origin HEAD:main
              """
            }
          }
        }
        /*
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
        }*/
    }

    post {
        always {
            sh 'docker image prune -f || true'
        }
    }
}
