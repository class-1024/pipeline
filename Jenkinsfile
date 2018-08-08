pipeline {
   agent any
   stages {
       stage('Clone') {
          steps {                
                echo "1.Clone Stage" 
                script {
                   build_tag = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()
                   echo "build_tag:${build_tag}"
               }
            }        
        }
        stage('Test') {
           steps {                
                echo "2.Test Stage"
                sh " helm install --name my-release stable/wordpress"
            }        
        }
        stage('Build') {
           steps {                
                echo "3.Build Docker Image Stage"  
                sh "docker build -t registry.cn-shenzhen.aliyuncs.com/kinot-k8s/kinot-web-api:${build_tag} ."
            }        
        }
        stage('Push') {
           steps {                
              echo "4.Push Docker Image Stage"
                sh "docker login -u=广州科玛技术有限公司 -p kinot-2018 registry.cn-shenzhen.aliyuncs.com"
                sh "docker push registry.cn-shenzhen.aliyuncs.com/kinot-k8s/kinot-web-api:${build_tag}"
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
