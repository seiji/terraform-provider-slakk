provider slack {}

data slack_user seiji {
  name = "seiji"
}

resource slack_channel tf_public {
  name       = "tf-public"
  is_private = false
}

output channel {
  value = slack_channel.tf_public
}

output user {
  value = data.slack_user.seiji
}
