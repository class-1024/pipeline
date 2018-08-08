
 pipeline {
    agent {
        label 'Production'
    }
    stages {
        stage('Build') {            
            steps {                
                echo 'Building'            
            }        
        }        
        stage('Test') {            
            steps {                
                echo 'Testing'            
            }        
        }
        stage('Deploy - Staging') {            
            steps {                 
                 echo './deploy staging ./run-smoke-tests'      
            }        
        }        
        stage('Sanity check') {            
            steps {       
               input "Does the staging environment look ok?"     
            }        
        }        
        stage('Deploy') {            
            steps {    
                echo './deploy production'                 
            }        
        }    
    }

    post {        
        always {            
            echo 'One way or another, I have finished'            
            deleteDir() /* clean up our workspace */        
        }        
        success {            
            echo 'I succeeeded!'        
        }        
        unstable {            
            echo 'I am unstable :/'        
        }        
        failure {            
            echo 'I failed :('        
        }        
        changed {            
            echo 'Things were different before...'        
        }    
    }
}
