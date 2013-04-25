import json, socket



if __name__ == '__main__':
    s = socket.create_connection(("localhost", 1234))
    s.sendall(
        json.dumps({"id": 1, "method": "RPC.Echo", "params": ["Gimmie stuff"]}).encode())
    
    data = s.recv(4096).decode()
    print data
    # data = json.loads(data)
    # for k, v in data.iteritems():
    	# print k, v
