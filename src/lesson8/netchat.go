package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type client struct {
	channel  chan<- string // an outgoing message channel
	nickname string        // a nickname for user in this channel
}

type extMessage struct {
	dest  string // destination nickname
	text  string // text of the message
	exit  bool   // flag close connection or not
	autor string // autor of the message
}

type nicknames []string // storage of uniq nicknames. Has a problem with race state, when too many clients

// command - structure for simple commands in the chatroom
type command struct {
	name        string
	description string
	handler     func(*extMessage)
}

type commands []command

var (
	entering      = make(chan client)
	leaving       = make(chan client)
	messages      = make(chan extMessage) // all incoming client messages
	baseNickNames = make(nicknames, 0)    // all nicknames
	chatCommands  = make(commands, 0)     // all chat commands
)

func main() {
	//	init logging of app
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetLevel(log.DebugLevel)
	// fill commands for chat
	fillCMDS(&chatCommands)

	tcpport := flag.Int("p", 8000, "Server tcp port which is using for accept connections")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *tcpport))
	if err != nil {
		log.Fatal(err)
	}
	log.WithFields(log.Fields{"loglevel": log.GetLevel()}).Infof("TCP server successfully started on %d tcp port", *tcpport)

	// start communicator to background processing
	go communicator()

	// recieve connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

// help - returns a cmd name and a cmd description
func (cmds *commands) help() string {
	output := ""
	for _, cmd := range *cmds {
		output += "/" + cmd.name + " - using for " + cmd.description + "\n"
	}
	return output
}

// addNewCMD - add new cmd in the cmds slice
func (cmds *commands) addNewCMD(cmd command) {
	tmpcmds := make([]command, 0)
	for _, c := range *cmds {
		tmpcmds = append(tmpcmds, c)
	}
	tmpcmds = append(tmpcmds, cmd)
	*cmds = tmpcmds
}

// execCMD - execute cmd
func (cmds *commands) execCMD(cmd string, msg *extMessage) {
	for _, c := range *cmds {
		if c.name == cmd {
			c.handler(msg)
		}
	}
}

// nickIsExists - check exists or not nick <name> in the storage of nicknames
func (nicks *nicknames) nickIsExists(name string) bool {
	for _, nick := range *nicks {
		if nick == name {
			return true
		}
	}
	return false
}

// addNewNick - add new nickname <name> in the storage of nicknames
func (nicks *nicknames) addNewNick(name string) {
	tmpnames := make([]string, 0)
	for _, nick := range *nicks {
		tmpnames = append(tmpnames, nick)
	}
	tmpnames = append(tmpnames, name)
	*nicks = tmpnames
}

// delNick - deletes nickname <name> from the storage of nicknames
func (nicks *nicknames) delNick(name string) {
	tmpnames := make([]string, 0)
	for _, nick := range *nicks {
		if nick != name {
			tmpnames = append(tmpnames, nick)
		}
	}
	*nicks = tmpnames
}

// fillCMDS - fills the slice of commands
func fillCMDS(cmds *commands) {
	// add cmd HELP
	cmd := command{
		name:        "help",
		description: "prints list available commands and them description",
		handler: func(msg *extMessage) {
			msg.text = cmds.help()
		},
	}
	cmds.addNewCMD(cmd)

	// add cmd WHO
	cmd = command{
		name:        "who",
		description: "prints list of connected users",
		handler: func(msg *extMessage) {
			for _, nick := range baseNickNames {
				msg.text += nick + "\n"
			}
		},
	}
	cmds.addNewCMD(cmd)

	// add cmd EXIT
	cmd = command{
		name:        "exit",
		description: "close connection to server",
		handler: func(msg *extMessage) {
			msg.text = "Goodbuy!"
			msg.exit = true
		},
	}
	cmds.addNewCMD(cmd)
}

// communicator - channel processing for sending and receiving messages, user inputs and outputs
func communicator() {
	clients := make(map[string]client) // all connected clients
	for {
		select {
		case msg := <-messages:
			// processing incoming message
			switch msg.dest {
			case "all": //broadcast
				for _, cli := range clients {
					if cli.nickname != msg.autor {
						cli.channel <- msg.text + "\n" + cli.nickname + ">"
					}
				}
			default: // personal message
				for _, cli := range clients {
					if cli.nickname == msg.dest {
						cli.channel <- msg.text + "\n" + cli.nickname + ">"
					}
				}
			}
			log.Debugf("User %s write message with length=%dsimbols for %s", msg.autor, len(msg.text), msg.dest)

		case cli := <-entering:
			clients[cli.nickname] = cli
			baseNickNames.addNewNick(cli.nickname)
			log.Debugf("New user [%s] connected to the chat", cli.nickname)

		case cli := <-leaving:
			delete(clients, cli.nickname)
			close(cli.channel)
			baseNickNames.delNick(cli.nickname)
			log.Debugf("User [%s] disconnected from the chat", cli.nickname)

		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go toClientWriter(conn, ch)
	ip := conn.RemoteAddr().String()
	log.Infof("Received the new connection to server from%s\n", ip)

	// get nickname from client
	who := setNickName(conn, &baseNickNames, ch)
	log.Debugf("From %s user set name: %s", ip, who)
	welcome := who + ">"
	ch <- welcome
	messages <- extMessage{dest: "all", text: who + " has arrived"}
	entering <- client{channel: ch, nickname: who}
	input := bufio.NewScanner(conn)
	for input.Scan() {
		tmpMsg := input.Text()
		if len(tmpMsg) < 2 {
			log.Warnf("Wrong input <%s>", tmpMsg)
			ch <- welcome
			continue
		}
		// check dest/cmd
		prefix := tmpMsg[:1]
		log.Debugf("PEFIX: %s\n", prefix)
		msg := extMessage{
			dest:  "all",
			autor: who,
		}
		switch prefix {
		case "@":
			//
			ixEndName := strings.Index(tmpMsg, " ")
			tmpNick := strings.Trim(tmpMsg[1:ixEndName], ",")
			msg.dest = tmpNick
			msg.text = "@" + who + ": " + tmpMsg[1:]
		case "/":
			// cmd received
			rawCmd := strings.Trim(tmpMsg, " ")
			cmd := strings.ToLower(rawCmd[1:])
			log.Debugf("CMD: %s\n", cmd)
			msg.dest = who
			chatCommands.execCMD(cmd, &msg)
		default:
			msg.exit = false
			msg.text = "@" + who + ": " + tmpMsg
		}
		ch <- welcome
		messages <- msg

		if msg.exit {
			time.Sleep(500 * time.Millisecond)
			break
		}
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- client{channel: ch, nickname: who}
	messages <- extMessage{dest: "all", text: who + " has left chatroom"}
	conn.Close()
}

// setNickName - returns the username that he entered with the uniqueness check
func setNickName(conn net.Conn, baseNicks *nicknames, ch chan<- string) string {
	who := conn.RemoteAddr().String()
	input := bufio.NewScanner(conn)
	msg := "Type /help for available commands help\nWhat is your name? (or use " + who + " like anonimus):"
	for {
		ch <- msg
		input.Scan()
		tmpwho := input.Text()
		if len(tmpwho) > 0 {
			who = tmpwho
		}
		if baseNickNames.nickIsExists(who) {
			msg = "This nick is occupied, type new:"
			continue
		}
		break
	}
	return who
}

// toClientWriter - consumes <ch> channel and writes messages to the client
func toClientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		_, err := fmt.Fprintf(conn, "%v", msg) // NOTE: ignoring network errors
		if err != nil {
			log.Errorf("Message [%s] don't delivered, %v", msg, err)
			// continue
		}
	}
}
