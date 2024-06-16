# URL_Shortener

A simple URL shortener service that will accept a URL as an argument over a REST API 
and return a shortened URL as a result.

Also a metrics API that returns top 3 domain names that have been shortened the most
number of times.

Docker Image for containerization
https://hub.docker.com/repository/docker/avbhatta/myurl/general

API Testing 
1. POST Method
![image](https://github.com/AvijitBhattacharjee/URL_Shortener/assets/49098193/f9935a9e-8c28-4d2c-b06f-cd96778f6f76)

2. GET Method
![image](https://github.com/AvijitBhattacharjee/URL_Shortener/assets/49098193/396d684f-b0de-44fa-b316-2004db6dabbc)

3. GET Top 3 Domains
![image](https://github.com/AvijitBhattacharjee/URL_Shortener/assets/49098193/708343cf-8b81-46cd-9b39-05cdd684e44f)

CI/CD Pipeline - 
Using GitHub Actions - 
![image](https://github.com/AvijitBhattacharjee/URL_Shortener/assets/49098193/055fa8df-e413-47e1-b0a1-07ce611b66d7)
![image](https://github.com/AvijitBhattacharjee/URL_Shortener/assets/49098193/3e2d07fb-873c-4f3d-9e29-fadf270f75c8)

Deployment - 
Using Helm Chart 
bash deployment.sh


