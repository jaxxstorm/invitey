# Invitey

Invite a user to lots of channels with the same prefix

# Usage

```
go run main.go --token="<token>" --email="<email>" --channel_prefix="team-"
team-engineering
team-hr
team-marketing
Would you like to add the user to the following found channels?: y
INFO[0004] Checking channel                              channel_name=team-engineering user=jaxxstorm432
INFO[0004] Inviting user to channel                      channel_name=team-engineering user=jaxxstorm432
INFO[0001] User is already in channel                    channel_name=team-engineering user=jaxxstorm432
INFO[0001] Checking channel                              channel_name=team-hr user=jaxxstorm432
INFO[0001] Inviting user to channel                      channel_name=team-hr user=jaxxstorm432
INFO[0002] User successfully invited to channel          channel_name=team-hr user=jaxxstorm432
INFO[0004] Checking channel                              channel_name=team-marketing user=jaxxstorm432
INFO[0004] Inviting user to channel                      channel_name=team-marketing user=jaxxstorm432
WARN[0004] slack app is not present in channel, cannot invite user  channel_name=team-marketing user=jaxxstorm432
```

