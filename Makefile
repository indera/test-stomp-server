

server:
	go run server.go


setup:
	go mod init stomp_server_example
	go mod tidy


# CLI usage

cli:
	# https://jasonrbriggs.github.io/stomp.py/quickstart.html
	# pip install stomp.py
	stomp -H localhost -P 61613


# > subscribe /queue/test
# Subscribing to '/queue/test' with acknowledge set to 'auto', id set to '1'
# > send /queue/test hello world
# >
# message-id: 1
# subscription: 1
#
# hello world


