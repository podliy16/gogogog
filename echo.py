import socket, os, time, sys

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

port = int(sys.argv[1])
s.bind(("localhost", port))
print(port)
s.listen(1)
while 1:
	conn, addr = s.accept()
	print(conn, addr)
	while 1:
		try:
			data = conn.recv(1024)
			time.sleep(5)
			conn.send(data)
		except:
			print('PANIC')
			break

conn.close()