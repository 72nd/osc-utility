# osc-utility

 <p align="center">
  <img width="128" height="128" src="misc/icon.png">
</p>

osc-utility is a simple CLI tool for testing the [Open Sound Control (OSC)](https://en.wikipedia.org/wiki/Open_Sound_Control) connections. The utility can send OSC command as well as spawn a OSC server to listen to messages from other OSC-enabled devices.


## Breaking CLI Changes with v0.3.0

In version 0.3.0, the CLI was redesigned with the intention of aligning it more closely with other CLI applications and preventing the implicit setting of flags. The principal change is that multiple values of the same type when sending messages are no longer separated by commas (e.g., `--int 3,4`), but rather the flag can be set multiple times (`--int 3 --int 4`). Users who rely on the legacy CLI may use version [0.2.2](https://github.com/72nd/osc-utility/releases/tag/v0.2.2a).


## Installation

Just head to the release section on the right and download the executable for your system. You can then execute it right away.


## Usage

You can get more-detailed help by calling `osc-utility --help`, `osc-utility message --help` and `osc-utility server --help`.

### Send a message

To send a message to `/channel/1/255` using the default host (localhost) and the port `9000`:

```shell script
osc-utility message --address /channel/1/255 --port 9000
```

To prevent any potential confusion, using localhost as the host will now output a info message. The host can be defined using the host flag:

```shell script
osc-utility message --address /channel/1/255 --port 9000 --host 192.168.1.100
```

OSC allows to sending a payload which can be either a string (text), int, float or bool. Remember to put strings containing spaces into quotes: 

```shell script
# String
osc-utility messsage --address /channel/1 --port 9000 --string "Hello World"

# Int
osc-utility messsage --address /channel/1 --port 9000 --int 23

# Float
osc-utility messsage --address /channel/1 --port 9000 --float 23.5

# Bool
# Use one of [true, t, 1] for true
# Use one of [false, f, 0] for false
osc-utility messsage --address /channel/1 --port 9000 --bool true
```

OSC Messages can contain multiple values of the same type. Osc-utility allows this by providing the same flag multiple times:

```shell script
# Send the values "Hello World", "Foo", and "Bar"
osc-utility messsage --address /channel/1 --port 9000 --string "Hello World" --string "Foo" --string "Bar"
```

Naturally it's possible to send values of multiple types at the same time:

```shell script
osc-utility messsage --address /channel/1 --port 9000 --string "Foo" --string "Bar" --int 23 --int 5
```


### Receive messages (Server Mode)

To run a OSC server on the default host (localhost) on port `9000` run:

```shell script
osc-utility server --port 9000
```

You will now see all incoming messages.


## Logging and Output

### JSON Output

When using the `--json-log` flag, the output will be in JSON format. This can be useful for parsing the output with tools like [jq](https://stedolan.github.io/jq/) or incorporating osc-utility into other tools. For example `osc-utility --json-log server --port 9000` will output something like this:

```json
{
    "time": "2025-09-06T14:18:30.207864+01:00",
    "level": "INFO",
    "msg": "new message",
    "address": "/channel/1",
    "booleans": [
        false,
        true
    ],
    "strings": [
        "Hello",
        "World"
    ],
    "integers": [
        2,
        8
    ],
    "floats": [
        23.42
    ]
}
```

Please note that with the server command in JSON mode, after starting the server, no output is provided indicating which host and port the server is running on.


### Debugging

To enable debug logging, run the utility with the `--debug` flag.

```shell script
osc-utility --debug message --address /channel/1 --port 9000
```
