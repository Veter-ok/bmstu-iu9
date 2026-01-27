import asyncio
import websockets
import json

class Wallet:
    def __init__(self):
        self.server_id = 6
        self.name = "Вова 2"
        self.INIT_NUMBER = 10
        self.amount = self.INIT_NUMBER
        self.block = False
        self.found_wallets = {}
        self.balances = {}
        self.queue = []
        self.IP = '172.20.10.2'
        self.IP2 = '172.20.10.'
        self.ip_controller = ''
    
    async def handle_client(self, websocket):
        try:
            async for message in websocket:
                if message == "wallet":
                    await websocket.send(str(self.server_id))
                    ip, p = websocket.remote_address
                    if ip != self.IP:
                        self.queue.append(ip)
                    if self.ip_controller == ip:
                        await self.update_controller()
                    print(f"Sent ID {self.server_id} to {websocket.remote_address}")
                elif message == "block":
                    self.block = True
                    print("block")
                elif message == "unblock":
                    self.block = False
                    print("unblock")
                elif message.startswith("send:"):
                    print(f"Received: {message} from {websocket.remote_address}")
                    _, from_id, to_id, amount = message.split(":")
                    from_id, to_id, amount = int(from_id), int(to_id), int(amount)
                    
                    if to_id == self.server_id:
                        self.amount += amount
                        self.balances[to_id] += amount
                        self.balances[from_id] -= amount
                        await websocket.send("ok")
                        print(f"Got {amount} from {from_id}. New_balance: {self.amount}")
                    else:
                        if from_id in self.balances:
                            self.balances[from_id] -= amount
                        if to_id in self.balances:
                            self.balances[to_id] += amount
                        await websocket.send("ok")
                elif message == "balance":
                    await websocket.send(self.create_json_balance())
                else:
                    await websocket.send("Unknown command")
        except Exception as e:
            print(f"Client error: {e}")
    
    async def start_server(self, host=None, port=3000):
        if host is None:
            host = self.IP
        try:
            server = await websockets.serve(self.handle_client, host, port)
            print(f"Wallet server {self.server_id} started on ws://{host}:{port}")
            return server
        except Exception as e:
            print(f"Failed to start server: {e}")
            return None

    async def scan_network(self, ip_range=None):
        if ip_range is None:
            ip_range = self.IP2
        print(f"Scanning network {ip_range}*")
        
        tasks = []
        for i in range(1, 255):
            ip = ip_range + str(i)
            task = self.check_wallet(ip, 3000) 
            tasks.append(task)
        
        results = await asyncio.gather(*tasks, return_exceptions=True)
        found_count = 0
        id = 0
        for i, result in enumerate(results):
            ip = ip_range + str(i + 1)
            if isinstance(result, int):
                if result > 0:
                    self.found_wallets[result] = ip
                    found_count += 1
                    id = result
                    print(f"+ Found wallet: ID {result} at {ip}")
                elif result == -2:
                    self.ip_controller = ip
                    print(f"+ Found controller on ip: ({ip})")
        if found_count == 1:
            self.balances[id] = self.INIT_NUMBER
        return found_count
    
    async def check_wallet(self, host, port, timeout=1):
        if host in ['0.0.0.0', 'localhost', '127.0.0.1']:
            return -1
        try:
            #print(f"Trying to connect to {host}...") 
            async with websockets.connect(f"ws://{host}:{port}") as ws:
                await ws.send("wallet")
                response = await asyncio.wait_for(ws.recv(), timeout=timeout)
                #response = await ws.recv()
                if int(response) > 0 or int(response) == -2:
                    #print(f"{host} Success")
                    return int(response)
                return -1
        except Exception:
            return -1
    
    def create_json_balance(self):
        d = {}
        for wallet_id, _ in self.found_wallets.items():
            d[str(wallet_id)] = {
                "balance": str(self.balances[wallet_id]),
                "name": self.name if wallet_id == self.server_id else f"Wallet_{wallet_id}"
            }
        return json.dumps(d)


    async def update_balance(self, timeout=0.5):
        try:
            for wallet_id, ip in self.found_wallets.items():
                if wallet_id != self.server_id:
                    money_data = await self.get_balance(ip)
                    for id_str, balance_str in money_data.items():
                        id_int = int(id_str)
                        balance_int = int(balance_str["balance"])
                        self.balances[id_int] = balance_int
                        if self.server_id == id_int:
                            self.amount = balance_int
            if not self.server_id in self.balances:
                self.balances[self.server_id] = self.INIT_NUMBER
                self.amount = self.INIT_NUMBER
        except Exception as e:
            print(f"error balance {e}")
        
    async def get_balance(self, ip, timeout=0.5):
        try:
            async with websockets.connect(f"ws://{ip}:3000") as ws:
                await ws.send("balance")
                response = await asyncio.wait_for(ws.recv(), timeout=timeout)
                balance_data = json.loads(response)
                return balance_data
        except Exception as e:
            print(f"Error getting balance from {ip}: {e}")
            print(balance_data)
            return {}
    
    async def send_money(self, id, send_money, timeout=0.5):
        if id not in self.found_wallets:
            print("Not such wallet!")
            return False
        if self.block:
            print("Someone is sending money!")
            return False
        if self.amount < send_money:
            print("Don't have enough money!")
            return False
        if send_money <= 0:
            print("Wrong sum!")
            return False
        if id == self.server_id:
            print("This is you!!!")
            return False
        #ip = self.found_wallets[id]
        try:
            for wallet_id, ips in self.found_wallets.items():
                if wallet_id != self.server_id:
                    try:
                        async with websockets.connect(f"ws://{ips}:3000") as ws:
                            await ws.send("block")
                    except Exception:
                        pass
            
            #send to needed address
            for wallet_id, ips in self.found_wallets.items():
                try:
                    async with websockets.connect(f"ws://{ips}:3000") as ws:
                        await ws.send(f"send:{self.server_id}:{id}:{send_money}")
                        response = await ws.recv()
                        if response == "ok" and wallet_id == id:
                            self.amount -= send_money
                            #self.balances[self.server_id] -= send_money
                            #self.balances[id] += send_money
                            print(f"Sent {send_money} to {id}. New_balance: {self.amount}")
                except Exception:
                    pass
            
            if self.ip_controller != '':
                try:
                    async with websockets.connect(f"ws://{self.ip_controller}:3000") as ws:
                        await ws.send(f"money_balance")
                        response = await ws.recv()
                        if response == "ok":
                            await ws.send(self.create_json_balance())
                except Exception:
                    print(f"Error notifying controller {self.controller_id}: {e}")


            for wallet_id, ips in self.found_wallets.items():
                if wallet_id != self.server_id:
                    try:
                        async with websockets.connect(f"ws://{ips}:3000") as ws:
                            await ws.send("unblock")
                    except Exception:
                        pass
            return True
        except Exception as e:
            print(f"Exception in sending: {e}")
            return False
    
    async def update_controller(self):
        try:
            async with websockets.connect(f"ws://{self.ip_controller}:3000") as ws:
                await ws.send(f"money_balance")
                response = await ws.recv()
                if response == "ok":
                    await ws.send(self.create_json_balance())
        except Exception as e:
            print(f"Error controller-update: {e}")
     # Запускаем обработку очереди в фоне
    async def interactive(self):
    # Запускаем обработку очереди в фоне
        async def process_queue():
            while True:
                if self.queue:
                    ip = self.queue.pop(0)
                    if ip not in self.found_wallets.values() and self.ip_controller != ip:
                        try:
                            async with websockets.connect(f"ws://{ip}:3000") as ws:
                                await ws.send("wallet")
                                response = await ws.recv()
                                response = int(response)
                                if response > 0:
                                        if response not in self.found_wallets:
                                            self.balances[response] = self.INIT_NUMBER
                                        self.found_wallets[response] = ip
                                        print(f"+ Added wallet {response} from queue")
                                elif response == -2:
                                        self.ip_controller, _ = ws.remote_address
                                        print(f"+ Added controller! on ip {self.ip_controller}")

                            if self.ip_controller != '' and ip == self.ip_controller:
                                await self.update_controller()

                        except Exception:
                            pass
                await asyncio.sleep(1)  
        
        queue_task = asyncio.create_task(process_queue())
        
        try:
            while True:
                print("\n"+40*"=")
                print("1. Send money")
                print("2. Balance")
                print("3. Update Balance")
                print("4. Exit")
                print(self.queue)
                try:
                    choice = await asyncio.get_event_loop().run_in_executor(
                        None, input, "Choose option: "
                    )
                    
                    if choice == "1":
                        self.print_status()
                        id = input("ID for sending: ")
                        money = input("Money: ")
                        await self.send_money(int(id), int(money))
                    elif choice == "2":
                        self.print_status()
                    elif choice == "3":
                        await self.update_balance()
                    else:
                        break
                        
                except KeyboardInterrupt:
                    break
                    
        finally:
            queue_task.cancel()  

    def print_status(self):
        print(f"\nOur wallet: {self.server_id} (balance: {self.amount})")
        #print(self.found_wallets)
        #print(self.balances)
        if self.found_wallets:
            print("Wallets:")
            for wallet_id, _ in self.found_wallets.items():
                balance = self.balances[wallet_id]
                print(f"\t{wallet_id}: {balance} coins")

async def main():
    wallet = Wallet()

    server = await wallet.start_server()
    if not server:
        return
    
    await asyncio.sleep(2)
    found = await wallet.scan_network()

    # Выводим результаты
    print(f"\n" + "="*50)
    print(f"SCAN COMPLETED!")
    print(f"Found {found} wallet(s) in network")
    print(f"Our server ID: {wallet.server_id}")
    
    if found > 0:
        print("\nDiscovered wallets:")
        for wallet_id, ip in wallet.found_wallets.items():
            print(f"  ID: {wallet_id} -> IP: {ip}")
    else:
        print("\nNo other wallets found in the network")
    
    await wallet.update_balance()
    if  wallet.ip_controller != '':
        await wallet.update_controller()
    
    await wallet.interactive()

if __name__ == "__main__":
    asyncio.run(main())
