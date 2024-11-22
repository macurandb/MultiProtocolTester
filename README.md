MultiProtocolTester
===================

**MultiProtocolTester** is a command-line application designed to send and receive messages using different protocols (e.g., Kafka, RabbitMQ). This project provides flexible configuration handling, allowing you to test various protocols using customizable configuration files.

Table of Contents
-----------------

*   [Installation](#installation)

*   [Usage](#usage)

    *   [Commands](#commands)

    *   [Protocol Support](#protocol-support)

*   [Configuration](#configuration)

*   [Makefile Commands](#makefile-commands)

*   [Contributing](#contributing)

*   [License](#license)


Installation
------------

To get started with **MultiProtocolTester**, follow these steps:

1.  Clone the repository:
    ```bash 
      git clone https://github.com/yourusername/MultiProtocolTester.gitcd MultiProtocolTester
    ```

2. Install the necessary dependencies:
    ```bash 
    go mod tidy
    ````

3. Build the project:
    ```bash
    make build
    ```

This will generate the prototester binary in the project directory.

Usage
-----

### Commands

The prototester CLI supports the following commands for interacting with various protocols:

*   **Send a message**: prototester \[protocol\] send --message "Your message"

*   **Receive messages**: prototester \[protocol\] receive


### Protocol Support

Supported protocols:

*   **Kafka**

*   **RabbitMQ** (or other protocols, depending on implementation)


To interact with these protocols, the CLI provides flexible configuration options for connecting to services like Kafka.

### Example Usage

*   Sending a Kafka message:
    ```bash
      ./prototester kafka send --message "Hello Kafka!"
    ```

*   Receiving messages from Kafka:
    ```bash
      ./prototester kafka receive
    ```
*   Sending a RabbitMQ message (assuming RabbitMQ is implemented):

    ```bash
      ./prototester rabbitmq send --message "Hello RabbitMQ!"
    ```
* Receiving messages from RabbitMQ:
    ```bash
      ./prototester rabbitmq receive
    ```


Configuration
-------------

You can specify a configuration file for each protocol using the --config flag. If no configuration file is provided, the default path config/\[protocol\]\_config.yaml is used.

*   Example Kafka configuration (config/kafka_config.yaml):
    ```yaml 
        kafka:
            bootstrap_servers: "localhost:9092"
            topics:
              send_topic: "test-topic"
              receive_topic: "test-topic"
            group_id: "test-consumer-group"
            timeout: 5000
    ```


To specify a custom configuration file, use the --config flag:
```bash
    ./prototester kafka send --message "Hello Kafka!" --config custom_config.yaml
```


Makefile Commands
-----------------

The project provides a **Makefile** to streamline common tasks:

*   Build the project:
    ```bash
     make build
    ```

*   Send a message using a protocol (e.g., Kafka):
    ```bash
     make kafka/send
    ```

*   Receive messages using a protocol (e.g., Kafka):
    ```bash
    make kafka/receive
    ```


You can replace kafka with other protocols (e.g., rabbitmq) in the Makefile commands.

Contributing
------------

We welcome contributions to the **MultiProtocolTester** project. If you'd like to contribute, please follow these steps:

1.  Fork the repository.

2.  Create a feature branch: git checkout -b my-new-feature.

3.  Commit your changes: git commit -m 'Add some feature'.

4.  Push to the branch: git push origin my-new-feature.

5.  Open a pull request.


Please ensure that your code follows the coding standards and includes tests where appropriate.

License
-------

This project is licensed under the MIT License. See the LICENSE file for details.
