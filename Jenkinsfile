pipeline {
   agent any
   stages {
       stage('Clone') {
          steps {                
                echo "1.Clone Stage"     
            }        
        }
        stage('Test') {
           steps {                
               echo "2.Test Stage"
            }        
        }
        stage('Build') {
           steps {                
                echo "3.Build Docker Image Stage"  
            }        
        }
        stage('Push') {
           steps {                
              echo "4.Push Docker Image Stage"
            }        
        }
        stage('YAML') {
           steps {                
               echo "5. Change YAML File Stage" 
            }        
        }
        stage('Deploy') {
           steps {                
              echo "6. Deploy Stage"  
            }        
        }
    }
    post {        
        always {            
            echo 'post always'            
            deleteDir() /* clean up our workspace */        
        }        
        success {            
            echo 'post success '        
        }        
        unstable {            
            echo 'post unstable'        
        }        
        failure {            
            echo 'post failure'         
        }        
        changed {            
            echo 'post changed'        
        }    
    }
}
