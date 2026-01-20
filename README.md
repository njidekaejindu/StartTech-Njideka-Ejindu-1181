# StartTech Full Stack Application – CI/CD Deployment

## Overview

This project is a full-stack web application with a React frontend and a backend API, fully automated using GitHub Actions and deployed on AWS.

The infrastructure and deployment pipeline have been built using:

- React (Frontend)
- AWS S3 for static hosting
- AWS CloudFront for CDN
- AWS Application Load Balancer (ALB)
- Terraform for infrastructure as code
- GitHub Actions for CI/CD automation

---

## Repository Structure

starttech-application/
├── backend/
├── frontend/
├── scripts/
├── .github/workflows/
├── docs/
└── README.md

---

## CI/CD Pipeline Summary

Whenever code is pushed to the master branch:

1. GitHub Actions installs dependencies  
2. React app is built  
3. Build files are synced to S3  
4. CloudFront cache is invalidated  
5. Users get the updated site immediately  

---

## Required GitHub Secrets

The following secrets are configured in GitHub:

- AWS_ACCESS_KEY_ID  
- AWS_SECRET_ACCESS_KEY  
- AWS_REGION  
- S3_BUCKET_NAME  
- CLOUDFRONT_DISTRIBUTION_ID  

---

## Local Development

To run the frontend locally:

cd frontend  
npm install  
npm start  

To build locally:

cd frontend  
npm run build  

---

## Deployment

Deployment is fully automated using:

.github/workflows/frontend-ci-cd.yml

No manual deployment is required.

## Evidence

All implementation evidence can be found in:

- /evidence folder in this repository
