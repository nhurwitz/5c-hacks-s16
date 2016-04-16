from websocket import create_connection

# https://pypi.python.org/pypi/websocket-client/

ws = create_connection("ws://localhost:8000/ws")
while True:
  result = ws.recv()
  print "Received '%s'" % result
