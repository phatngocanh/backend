pipeline {
    agent any
    environment {
        PORT = credentials('BACKEND_PORT_DONHANG')
        

        DOCKER_TAG = 'latest'
        CONTAINER_NAME = 'donhang-be-container'
    }

    stages {
        stage('Remove Old Docker Image') {
            steps {
                script {
                    echo "Stopping and removing old Docker container..."
                    sh "docker stop ${env.CONTAINER_NAME} || true"
                    sh "docker rm ${env.CONTAINER_NAME} || true"
                    
                    echo "Removing old Docker image..."
                    sh "docker rmi donhang-be:${env.DOCKER_TAG} || true"
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    sh "docker build -t donhang-be:${env.DOCKER_TAG} ."
                }
            }
        }

        stage('Run Container') {
            steps {
                script {
                    sh """
                        docker run -d \\
                            --restart unless-stopped \\
                            --name ${env.CONTAINER_NAME} \\
                            -p ${env.PORT}:${env.PORT} \\
                            -e PORT="${env.PORT}" \\
                            donhang-be:${env.DOCKER_TAG}
                    """
                }
            }
        }
    }

    post {
        success {
            echo 'Pipeline completed successfully'
            cleanWs()
        }
        failure {
            echo 'Pipeline failed'
            script {
                sh "docker stop ${env.CONTAINER_NAME} || true"
                sh "docker rm ${env.CONTAINER_NAME} || true"
                cleanWs()
            }
        }
        always {
            echo 'Pipeline completed'
            cleanWs()
        }
    }
}
