# added to fix the rule
resource "aws_iam_role" "rds_monitoring_role" {
  name               = "rds-monitoring-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "monitoring.rds.amazonaws.com"
        }
      },
      {
        Effect = "Allow",
        Principal = {
          Service = "support.amazonaws.com"
        },
        Action = "sts:AssumeRole"
      }
    ]
  })
}

# added to fix the rule
resource "aws_iam_role_policy_attachment" "rds_monitoring_policy" {
  role       = aws_iam_role.rds_monitoring_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonRDSEnhancedMonitoringRole"
}

# added to fix the rule
resource "aws_accessanalyzer_analyzer" "example_us_east_1" {
  analyzer_name     = "example-analyzer-us-east-1"
  type              = "ACCOUNT"
}

# added to fix the rule
resource "aws_iam_role_policy_attachment" "support_role_policy_attachment" {
  role       = aws_iam_role.rds_monitoring_role.name
  policy_arn = "arn:aws:iam::aws:policy/AWSSupportAccess"
}
