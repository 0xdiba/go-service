resource "aws_ecs_task_definition" "go-svc-tdef" {
  family                = "go-svc-tdef"
  container_definitions = templatefile("container_def.json", {image_version=var.image_version, task_role=var.task_role})
}

resource "aws_ecs_service" "go-svc" {
  name            = "go-svc-tf"
  cluster         = var.cluster_name
  task_definition = aws_ecs_task_definition.go-svc-tdef.arn
  desired_count   = 1
  iam_role        = var.iam_role

  load_balancer {
    target_group_arn = var.target_grp
    container_name   = "go-svc"
    container_port   = 8000
  }

  capacity_provider_strategy {
    capacity_provider = var.capacity_provider
    weight            = 1
    base              = 0
  }
}
