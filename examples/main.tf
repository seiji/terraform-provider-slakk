data slack_conversations current {
  exclude_archived = true
  types            = ["public_channel", "private_channel"]
}

data slack_emojis current {}

output conversations {
  value = data.slack_conversations.current
}

output emojis {
  value = data.slack_emojis.current
}
