# TCPChat(NetCat)

TCPChat is a project that recreates the functionality of the NetCat command-line utility in a Server-Client architecture. It allows the server to run in server mode, listening for incoming connections on a specified port, and the client to run in client mode, attempting to connect to a specified port and transmitting information to the server.

NetCat (nc) is a versatile command-line utility used for reading and writing data across network connections using TCP or UDP. It supports various network operations such as opening TCP connections, sending UDP packets, listening on arbitrary TCP and UDP ports, and more. For detailed information about NetCat, refer to the manual by executing `man nc` in your terminal.

TCPChat goes beyond NetCat by creating a group chat environment. It includes the following features:

- TCP connection between the server and multiple clients in a one-to-many relationship.
- Name requirement for each client to provide clear identification during conversations.
- Control over the quantity of connections to manage server capacity effectively.
- Clients can send messages to the chat, fostering interactive communication.
- Empty messages from clients are not broadcasted to other participants.
- Messages sent are timestamped with the time of sending and the username of the sender in the format: `[2020-01-20 15:48:41][client.name]:[client.message]`.
- When a client joins the chat, they receive all previous messages sent in the chat.
- Server informs the rest of the clients when a new client joins the group.
- Server notifies the other clients when a client exits the chat.
- All clients receive messages sent by other clients, enabling real-time collaboration.
- If a client leaves the chat, the remaining clients stay connected without interruption.
- If no port is specified, the default port 8989 is used. Otherwise, the program responds with a usage message: `[USAGE]: ./TCPChat $port`.

## Instructions

To set up and run TCPChat, follow these instructions:

1. Ensure that you have Go installed on your machine.
2. Clone the TCPChat repository to your local environment.
3. Open a terminal and navigate to the project directory.
4. Start the TCP server by executing the following command:

   ```bash
   go run . [port]
   ```

   - Replace `[port]` with the desired port number. If no port is specified, the default port 8989 will be used.

5. Open separate terminals for each client and run the following command in each:

   ```bash
   nc localhost 8989 or nc localhost [port number]
   ```

6. Follow the prompts to enter a unique name for each client and start sending messages in the chat.

## Project Details

TCPChat is implemented using Go and incorporates various concurrent programming concepts to ensure efficient communication and synchronization between clients and the server. The key project details include:

- Go-routines: Go-routines are utilized to handle concurrent execution and enable smooth communication between multiple clients and the server.
- Channels and Mutexes: Channels and Mutexes are employed to manage shared resources and synchronize access to critical sections of the code, ensuring data consistency and preventing race conditions.
- Connection Limit: The server enforces a maximum limit of 10 client connections to maintain optimal performance.
- Error Handling: The code is designed to handle errors gracefully, both on the server and client sides, ensuring a reliable and robust chat experience.
- Unit Testing: The project includes test files to facilitate unit testing of the server connection and client functionality.
- Best Practices: The code adheres to industry best practices to ensure readability, maintainability, and adherence to Go coding conventions.

Feel free to explore and enhance TCPChat according to your needs and requirements. Enjoy collaborative and interactive chat experiences with TCPChat!
