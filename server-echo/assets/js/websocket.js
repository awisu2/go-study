class App extends React.Component {
  constructor(props) {
    super(props);
    // stateのセットは初期化のみ(constructor内でのsetStateはNG)
    this.state = { conn: null, message: "", logs: [], };
  }

  componentDidMount() {
    // setStateのある初期処理はここで行う
    this.connectWebSocket()
  }

  addLog(log) {
    this.setState({logs: [log].concat(this.state.logs)})
  }

  setMessage(message) {
    this.setState({message: message})
  }

  setConn(conn) {
    this.setState({conn: conn})
  }

  connectWebSocket() {
    var conn = this.createConnection()

    if (conn == null) {
      this.addLog(<b>Your browser does not support WebSockets.</b>);
      return 
    }
    this.setConn(conn)
    this.addLog("connnected")
  }

  // websocket connection作成
  createConnection() {
    if (!window["WebSocket"]) {
      return null
    }

    var conn = new WebSocket("ws://" + document.location.host + "/ws");
    conn.onclose = (evt) => {
      this.addLog(<b>Connection closed.</b>);
      this.setConn(null)
    };

    conn.onmessage = (evt) => {
      var messages = evt.data.split("\n");
      for (var i = 0; i < messages.length; i++) {
        this.addLog(messages[i]);
      }
    };

    return conn
  }

  sendMessage(e) {
    e.preventDefault()

    if (!this.state.conn) {
      this.addLog("connection nothing")
      return
    }

    if (!this.state.message) {
      this.addLog("no message")
      return
    }

    this.state.conn.send(JSON.stringify({message: this.state.message}))
  }

  render() {
    var logsTag = this.state.logs.map(log => {
      return <div>{log}</div>
    })

    return (
      <div>
        <form onSubmit={this.sendMessage.bind(this)}>
          <input type="submit" value="Send" />
          <input type="text" value={this.state.message} onChange={(e) => this.setMessage(e.target.value)} autofocus />
        </form>
        <div>{logsTag}</div>
      </div>
    );
  }
}

// rendering
ReactDOM.render(React.createElement(App), document.querySelector("#app"));
