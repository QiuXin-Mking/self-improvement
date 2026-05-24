# q
What are the two main components of the AWS project assignment?
# a
1. Web Application Component: Hosted on EC2 instances in an Auto Scaling Group (ASG) with an Application Load Balancer (ALB); supports image upload (store in S3, metadata in RDS MySQL) and a page displaying images/thumbnails/captions.
2. Serverless Component: Two Lambda functions triggered by S3 `ObjectCreated` events – an Annotation Function (invokes Gemini API, stores results in RDS) and a Thumbnail Generator (generates thumbnail, stores in `thumbnails/` folder of same S3 bucket).

# q
What are the detailed requirements for the Auto Scaling Group and Load Balancer in the web application?
# a
- Compute: EC2 ASG with scaling policies (based on CPU/memory utilization), max capacity >1 instance.
- Networking: VPC/subnets, security groups, IAM roles, Bastion Host for admin access.
- Load Balancer: ALB setup with listener rules, target groups, and health checks for traffic distribution.

# q
How are the serverless Lambda functions triggered and what are their interactions?
# a
Triggered by S3 `ObjectCreated` events (no SNS/EventBridge unless specified). Interactions:
- Annotation Function: Retrieves image from S3 → invokes Gemini API → stores description in RDS.
- Thumbnail Generator: Retrieves image from S3 → generates thumbnail → stores in `thumbnails/` folder in the same S3 bucket.
- Both require IAM roles for S3, RDS, and (for Annotation) external API access.

# q
What evidence is required for the Auto Scaling test?
# a
- EC2 scaling actions (instance count changes)
- CPU/memory metrics trends
- ALB request distribution metrics
Tools: Apache Bench, Python concurrent request scripts, or similar.

# q
What are the graded components of the report submission and their point values?
# a
1. Introduction – brief overview.
2. Architecture Diagrams (5 points): Web app architecture, serverless architecture, and component interaction (S3 triggers, shared resources).
3. Web Application Deployment (4 points): ASG, RDS, S3 configuration details.
4. Serverless Component Deployment (3 points): Event-driven design, Lambda packaging/deployment, IAM/security.
5. Auto Scaling Test Observation (1 point): Screenshots/logs demonstrating scaling/load distribution.
6. Summary & Lessons Learned.
7. Report Professionalism (1 point): Structure, formatting, accurate naming.

