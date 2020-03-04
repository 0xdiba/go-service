variable "cluster_name" {
  type        = string
  description = "The name of the cluster to deploy the service on"
}

variable "iam_role" {
  type        = string
  description = "The IAM role used to deploy the service"
}

variable "target_grp" {
  type        = string
  description = "The loadbalancer target group, used to load balances between different container instances."
}

variable "capacity_provider" {
  type        = string
  description = "The capacity provider used to manage service resources"
}

variable "image_version" {
  type        = string
  description = "The docker image version to deploy"
}

variable "task_role" {
  type        = string
  description = "The task role to be used to deploy service tasks"
}
