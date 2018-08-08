
 pipeline {
    stage('Clone') {
      echo "1.Clone Stage"
    }
    stage('Test') {
      echo "2.Test Stage"
    }
    stage('Build') {
      echo "3.Build Docker Image Stage"
    }
    stage('Push') {
      echo "4.Push Docker Image Stage"
    }
    stage('YAML') {
      echo "5. Change YAML File Stage"
    }
    stage('Deploy') {
      echo "6. Deploy Stage"
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
