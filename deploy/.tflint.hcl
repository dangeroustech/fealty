config {
  module = false
  force  = false
  disabled_by_default = false
}

rule "terraform_naming_convention" {
  format  = "mixed_snake_case"
  enabled = true
}

rule "terraform_unused_declarations" {
  enabled = true
}

rule "terraform_unused_required_providers" {
  enabled = true
}

rule "terraform_comment_syntax" {
  enabled = true
}
