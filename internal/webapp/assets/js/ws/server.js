
class ServerLogMessageObserver {
    listeners=[]
    addListener(listener){
        this.listeners.push(listener)
        console.log(this.listeners)
    }
    removeListener(listener){
        this.listeners.remove(listener)
    }
    notify(message){
        this.listeners.forEach((l)=>l(message))
    }
}

const ServerLogObserver = new ServerLogMessageObserver();

function connectWebsocket(wsUrl) {
    let ws = new WebSocket(wsUrl)
    let wsMessageReader = new WebsocketMessageReader()

    ws.onopen = function (evt) {
        console.log("opened channel")
    };
    ws.onmessage = function (evt) {
        const messages = wsMessageReader.readMessages(evt.data);
        for (const k in messages) {
            if (!messages.hasOwnProperty(k)) {
                continue
            }
            const message = messages[k]
            const data = JSON.parse(message)
            ServerLogObserver.notify(data)
        }
        ws.onclose = function (evt) {
            console.log("closed channel")
            console.log(evt)
        };
        ws.onerror = function (evt) {
            console.log("error")
            console.log(evt)
        };
    };
}
