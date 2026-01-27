#!/usr/bin/python3
import asyncio
import websockets

async def hello(websocket):
    name = await websocket.recv()
    print("client say: "+str(name))

    while(True):
        greeting = input("enter answer to client: ")
        await websocket.send(greeting)
        print(str(greeting))

async def main():
    async with websockets.serve(hello, "185.102.139.161", 5050):
        await asyncio.Future()  # run forever

if __name__ == "__main__":
    asyncio.run(main())