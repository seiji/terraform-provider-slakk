provider slack {}

data slack_user seiji {
  name = "seiji"
}

resource slack_channel public {
  name       = "tf-public"
  is_private = false
  user_ids   = [data.slack_user.seiji.id]
}

output channel {
  value = slack_channel.public
}

output user {
  value = data.slack_user.seiji
}
