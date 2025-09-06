# osc-utility

 <p align="center">
  <img width="128" height="128" src="misc/icon.png">
</p>

osc-utility is a simple CLI tool for testing the [Open Sound Control (OSC)](https://en.wikipedia.org/wiki/Open_Sound_Control) connections. The utility can send OSC command as well as spawn a OSC server to listen to messages from other OSC-enabled devices.


## Installation

Just head to the release section on the right and download the executable for your system. You can then execute it right away.


## Usage

You can get more-detailed help by calling `osc-utility --help`, `osc-utility message --help` and `osc-utility server --help`.

### Send a message

To send a message to `/channel/1/255` using the default host (localhost) and the port `9000`:

```shell script
osc-utility message --address /channel/1/255 --port 9000
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
osc-utility messsage --address /channel/1 --port 9000 --bool true
```

OSC Messages can contain multiple values of the same type. Osc-utility allows this by separate this values by comma (do not insert any whitespace between the elements):

```shell script
# Send the values "Hello World", "Foo", and "Bar"
osc-utility messsage --address /channel/1 --port 9000 --string "Hello World,Foo,Bar"
```

Naturally it's possible to send values of multiple types at the same time:

```shell script
osc-utility messsage --address /channel/1 --port 9000 --string "Foo,Bar" -int 23,5
```


### Receive messages (Server Mode)

To run a OSC server on the default host (localhost) on port `9000` run:

```shell script
osc-utility server --port 9000
```

You will now see all incoming messages.


### Debugging

To enable debug logging, run the utility with the `--debug` flag.

```shell script
osc-utility --debug
```
