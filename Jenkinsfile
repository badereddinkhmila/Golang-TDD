/* groovylint-disable-next-line CompileStatic */
pipeline {
    agent any

    tools {
        // Configure Go version in Jenkins (Manage Jenkins -> Global Tool Configuration)
        go '1.26'
    }

    environment {
        // Enable Go modules
        GO111MODULE = 'on'
        // Prevent Jenkins from re-interpolating variables
        GO_VERSION = '1.26'
    }

    stages {
        stage('Lint') {
            steps {
                script {
                    echo "Running Go lint on branch ${env.BRANCH_NAME}"
                    docker.image('golangci/golangci-lint:v1.55.2').inside {
                        sh 'golangci-lint run ./...'
                    }
                }
            }
        }

        stage('Test') {
            steps {
                echo "Running Go tests on branch ${env.BRANCH_NAME}"
                sh 'make test'
            }
        }

        stage('Build') {
            steps {
                echo "Building Go app on branch ${env.BRANCH_NAME}"
                sh 'go build -v ./...'
            }
        }

        stage('Deploy to Staging') {
            when {
                // This stage only runs for the staging branch
                branch 'staging'
            }
            steps {
                echo 'Deploying to Staging Environment'
            // Example: sh 'kubectl apply -f k8s/staging/'
            // Example: sh 'gcloud app deploy --project=my-project-staging'
            }
        }

        stage('Deploy to Production - Approval Required') {
            when {
                // This stage only runs for the production/main branch
                branch 'main'
            }
            input {
                message 'Approve deployment to Production?'
                ok 'Deploy'
                submitter 'admin' // Optional: Restrict who can approve
            }
            steps {
                echo 'Deploying to Production Environment'
            // Example: sh 'kubectl apply -f k8s/prod/'
            // Example: sh 'gcloud app deploy --project=my-project-prod'
            }
        }
    }

    post {
        always {
            // Clean up workspace to save disk space
            cleanWs()
        }
        failure {
            echo "Pipeline failed for branch ${env.BRANCH_NAME}"
        }
    }
}
