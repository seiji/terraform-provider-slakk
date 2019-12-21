provider slack {
}

data slack_channel this {
  name = "general"
}

resource slack_channel this {
  name = "tf_test"
  # is_private = true
}

output channel {
  value = data.slack_channel.this
}
