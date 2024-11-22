locals {
  required_tags = {
    project     = var.project_name,
    environment = var.environment,
  }
}

resource "aws_vpc" "rds_vpc" {
  cidr_block = var.vpc_cidr
}

resource "aws_flow_log" "example" {
  traffic_type = "ALL"
  vpc_id       = aws_vpc.rds_vpc.id
}

resource "aws_subnet" "subnet_rds" {
  vpc_id     = aws_vpc.rds_vpc.id
  cidr_block = var.vpc_cidr
}

resource "aws_db_subnet_group" "main_rds" {
  name       = "rds_subnet"
  subnet_ids = [aws_subnet.subnet_rds.id]
}

# DB - RDS Instance
resource "aws_db_instance" "main" {
  engine                 = var.engine_name
  engine_version         = var.engine_version
  allocated_storage      = var.storage
  db_subnet_group_name   = aws_db_subnet_group.main_rds.name
  identifier             = var.identifier
  instance_class         = var.instance_class
  multi_az               = var.multi_az
  db_name                = var.database_name
  username               = var.database_username
  password               = var.database_password
  port                   = var.database_port
  publicly_accessible    = var.publicly_accessible
  vpc_security_group_ids = [var.db_security_group]
  skip_final_snapshot    = var.database_snapshot

  tags = local.required_tags

  # corrected tests
  storage_encrypted          = true
  auto_minor_version_upgrade = true
  # Monitoring is enabled
  monitoring_interval = 60
  monitoring_role_arn = var.monitoring_role

  backup_retention_period             = 7 # This enables automated backups with a 7-day retention period
  iam_database_authentication_enabled = true
  # Logging is enabled
  enabled_cloudwatch_logs_exports = ["audit", "error", "general", "slowquery"]

  copy_tags_to_snapshot = true
  deletion_protection   = true
}
